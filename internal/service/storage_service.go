package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	storage "github.com/hoshina-dev/papi/internal/infra/s3"
)

type StorageService struct {
	storage storage.StorageService
}

func NewStorageService(storage storage.StorageService) *StorageService {
	return &StorageService{storage: storage}
}

func (s *StorageService) GenerateUploadURL(ctx context.Context, contentType string) (uploadURL, fileKey string, err error) {
	fileKey = s.generateFileKey()

	uploadURL, err = s.storage.GeneratePresignedUploadURL(ctx, fileKey, contentType)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate upload URL: %w", err)
	}

	return uploadURL, fileKey, nil
}

func (s *StorageService) generateFileKey() string {
	return fmt.Sprintf("uploads/%d-%s", time.Now().Unix(), uuid.New().String())
}

func (s *StorageService) DeleteFile(ctx context.Context, fileKey string) error {
	return s.storage.Delete(ctx, fileKey)
}
