package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	// ClientPresignTTL is used for presigned URLs returned directly to clients (short-lived).
	ClientPresignTTL = 15 * time.Minute
	// JobPresignTTL is used for presigned URLs embedded in RabbitMQ job messages (long-lived).
	JobPresignTTL = 7 * 24 * time.Hour
)

type StorageService interface {
	GeneratePresignedUploadURL(ctx context.Context, key, contentType string, expiry time.Duration) (string, error)
	GeneratePresignedDownloadURL(ctx context.Context, key string, expiry time.Duration) (string, error)
	Delete(ctx context.Context, key string) error
}

type s3StorageService struct {
	client        *s3.Client
	presignClient *s3.PresignClient
	bucket        string
	baseURL       string
}

func NewS3StorageService(client *s3.Client, bucket, baseURL string) *s3StorageService {
	return &s3StorageService{
		client:        client,
		presignClient: s3.NewPresignClient(client),
		bucket:        bucket,
		baseURL:       baseURL,
	}
}

func (s *s3StorageService) GeneratePresignedUploadURL(ctx context.Context, key, contentType string, expiry time.Duration) (string, error) {
	result, err := s.presignClient.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(key),
		ContentType: aws.String(contentType),
	}, s3.WithPresignExpires(expiry))
	if err != nil {
		return "", fmt.Errorf("failed to generate presigned upload URL: %w", err)
	}
	return result.URL, nil
}

func (s *s3StorageService) GeneratePresignedDownloadURL(ctx context.Context, key string, expiry time.Duration) (string, error) {
	result, err := s.presignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	}, s3.WithPresignExpires(expiry))
	if err != nil {
		return "", fmt.Errorf("failed to generate presigned download URL: %w", err)
	}
	return result.URL, nil
}

func (s *s3StorageService) Delete(ctx context.Context, key string) error {
	_, err := s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	return err
}
