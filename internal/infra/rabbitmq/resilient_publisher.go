package rabbitmq

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// ErrNotAvailable is returned by Publish when the RabbitMQ connection is not
// yet established or is currently recovering from a disconnect.
var ErrNotAvailable = errors.New("rabbitmq not available")

// ResilientPublisher maintains a persistent AMQP connection and reconnects
// automatically when the connection drops. It implements the Publisher interface.
//
// Call Run(ctx) in a dedicated goroutine. Publish is safe to call concurrently
// and will return ErrNotAvailable while a reconnect is in progress.
type ResilientPublisher struct {
	url          string
	exchangeName string

	mu   sync.RWMutex
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewResilientPublisher(url, exchangeName string) *ResilientPublisher {
	return &ResilientPublisher{
		url:          url,
		exchangeName: exchangeName,
	}
}

// Run keeps the AMQP connection alive until ctx is cancelled. It retries on
// failure with a fixed 5-second backoff. It never returns a non-nil error so
// it is safe to run inside an errgroup without causing a group-wide shutdown.
func (p *ResilientPublisher) Run(ctx context.Context) {
	for {
		if err := p.connect(); err != nil {
			log.Printf("[rabbitmq] connection failed: %v — retrying in 5s", err)
			select {
			case <-ctx.Done():
				return
			case <-time.After(5 * time.Second):
				continue
			}
		}

		log.Printf("[rabbitmq] connected to exchange %q", p.exchangeName)

		notifyClose := p.conn.NotifyClose(make(chan *amqp.Error, 1))
		select {
		case <-ctx.Done():
			p.shutdown()
			return
		case amqpErr := <-notifyClose:
			log.Printf("[rabbitmq] connection lost: %v — reconnecting...", amqpErr)
			p.clearChannel()
		}
	}
}

// Publish serialises payload as JSON and publishes it as a persistent message.
// Returns ErrNotAvailable if the connection is currently down.
func (p *ResilientPublisher) Publish(ctx context.Context, exchange, routingKey string, payload any) error {
	p.mu.RLock()
	ch := p.ch
	p.mu.RUnlock()

	if ch == nil {
		return ErrNotAvailable
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	return ch.PublishWithContext(ctx, exchange, routingKey, false, false, amqp.Publishing{
		ContentType:  "application/json",
		DeliveryMode: amqp.Persistent,
		Body:         body,
	})
}

func (p *ResilientPublisher) connect() error {
	conn, ch, err := Connect(p.url)
	if err != nil {
		return err
	}

	if err := DeclareExchange(ch, p.exchangeName); err != nil {
		conn.Close()
		return err
	}

	p.mu.Lock()
	p.conn = conn
	p.ch = ch
	p.mu.Unlock()

	return nil
}

func (p *ResilientPublisher) clearChannel() {
	p.mu.Lock()
	p.ch = nil
	p.mu.Unlock()
}

func (p *ResilientPublisher) shutdown() {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.conn != nil && !p.conn.IsClosed() {
		p.conn.Close()
	}
	p.conn = nil
	p.ch = nil
}