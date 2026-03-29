package service

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/hoshina-dev/papi/internal/model"
)

type MockOptimizer struct {
	webhookURL string
	client     *http.Client
}

func NewMockOptimizer(webhookURL string) *MockOptimizer {
	return &MockOptimizer{
		webhookURL: webhookURL,
		client:     http.DefaultClient,
	}
}

func (m *MockOptimizer) Publish(job model.OptimizationJob) error {
	// Upload dummy optimized file
	dummyFileSize := int64(1024 * 1024)
	optimizedDummyFileSize := int64(512 * 1024)

	if err := m.uploadDummyFile(job.DestGLMURL, optimizedDummyFileSize); err != nil {
		return err
	}

	// Call webhook
	webhook := map[string]any{
		"uuid":                job.UUID,
		"status":              "success",
		"exit_code":           0,
		"logs":                "Mock optimization completed successfully",
		"timestamp":           time.Now(),
		"source_url":          job.SourceGLMURL,
		"dest_url":            job.DestGLMURL,
		"source_file_size":    dummyFileSize,
		"processed_file_size": optimizedDummyFileSize,
		"started_at":          time.Now(),
		"completed_at":        time.Now(),
		"duration_seconds":    2,
	}

	if job.DracoCompressionLevel != nil {
		webhook["draco_compression_level"] = job.DracoCompressionLevel
	}
	if job.DracoPositionQuantization != nil {
		webhook["draco_position_quantization"] = job.DracoPositionQuantization
	}
	if job.DracoTexcoordQuantization != nil {
		webhook["draco_texcoord_quantization"] = job.DracoTexcoordQuantization
	}
	if job.DracoNormalQuantization != nil {
		webhook["draco_normal_quantization"] = job.DracoNormalQuantization
	}
	if job.DracoGenericQuantization != nil {
		webhook["draco_generic_quantization"] = job.DracoGenericQuantization
	}

	if err := m.callWebhook(webhook); err != nil {
		return err
	}

	log.Printf("Mock optimization completed for job %s", job.UUID)
	return nil
}

func (m *MockOptimizer) uploadDummyFile(destURL string, fileSize int64) error {
	dummyFile := io.LimitReader(rand.Reader, fileSize)

	req, err := http.NewRequest(http.MethodPut, destURL, dummyFile)
	if err != nil {
		return fmt.Errorf("failed to create upload request: %v", err)
	}
	req.Header.Set("Content-Type", "model/gltf-binary")
	req.ContentLength = fileSize

	resp, err := m.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to upload file: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("upload failed with status %d: %s", resp.StatusCode, string(body))
	}

	return nil
}

func (m *MockOptimizer) callWebhook(payload any) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal webhook payload: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, m.webhookURL, bytes.NewReader(jsonPayload))
	if err != nil {
		return fmt.Errorf("failed to create webhook request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := m.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send webhook: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		return fmt.Errorf("webhook returned status %d: %s", res.StatusCode, string(body))
	}

	return nil
}
