package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/hoshina-dev/papi/internal/infra/rabbitmq"
	storage "github.com/hoshina-dev/papi/internal/infra/s3"
	"github.com/hoshina-dev/papi/internal/model"
	"github.com/hoshina-dev/papi/internal/repository"
)

type OptimizationService struct {
	storage            storage.StorageService
	publisher          rabbitmq.Publisher
	part3DModelRepo    repository.Part3DModelRepository
	webhookURL         string
	rabbitmqExchange   string
	rabbitmqRoutingKey string
}

func NewOptimizationService(
	storage storage.StorageService,
	publisher rabbitmq.Publisher,
	part3DModelRepo repository.Part3DModelRepository,
	webhookURL string,
	rabbitmqExchange string,
	rabbitmqRoutingKey string,
) *OptimizationService {
	return &OptimizationService{
		storage:            storage,
		publisher:          publisher,
		part3DModelRepo:    part3DModelRepo,
		webhookURL:         webhookURL,
		rabbitmqExchange:   rabbitmqExchange,
		rabbitmqRoutingKey: rabbitmqRoutingKey,
	}
}

type Optimize3DParams struct {
	SourceURL                 string
	DracoCompressionLevel     *int32
	DracoPositionQuantization *int32
	DracoTexcoordQuantization *int32
	DracoNormalQuantization   *int32
	DracoGenericQuantization  *int32
}

func (s *OptimizationService) Optimize3D(ctx context.Context, params Optimize3DParams) (jobID uuid.UUID, status string, err error) {
	jobID = uuid.New()

	destKey := s.generateDestinationKey(jobID)

	destURL, err := s.storage.GeneratePresignedUploadURL(ctx, destKey, "model/gltf-binary")
	if err != nil {
		return uuid.Nil, "", fmt.Errorf("failed to generate destination URL: %w", err)
	}

	sourcePresignedURL, err := s.extractOrGenerateSourceURL(ctx, params.SourceURL)
	if err != nil {
		return uuid.Nil, "", fmt.Errorf("failed to process source URL: %w", err)
	}

	part3DModel := &model.Part3DModel{
		ID:           jobID,
		RawURL:       params.SourceURL,
		ProcessedKey: &destKey,
		Status:       model.Part3DModelStatusProcessing,
	}

	err = s.part3DModelRepo.Create(ctx, part3DModel)
	if err != nil {
		return uuid.Nil, "", fmt.Errorf("failed to create 3D model record: %w", err)
	}

	job := model.OptimizationJob{
		UUID:         jobID.String(),
		SourceGLMURL: sourcePresignedURL,
		DestGLMURL:   destURL,
		WebhookURL:   s.webhookURL,
	}

	if params.DracoCompressionLevel != nil {
		v := int(*params.DracoCompressionLevel)
		job.DracoCompressionLevel = &v
	}
	if params.DracoPositionQuantization != nil {
		v := int(*params.DracoPositionQuantization)
		job.DracoPositionQuantization = &v
	}
	if params.DracoTexcoordQuantization != nil {
		v := int(*params.DracoTexcoordQuantization)
		job.DracoTexcoordQuantization = &v
	}
	if params.DracoNormalQuantization != nil {
		v := int(*params.DracoNormalQuantization)
		job.DracoNormalQuantization = &v
	}
	if params.DracoGenericQuantization != nil {
		v := int(*params.DracoGenericQuantization)
		job.DracoGenericQuantization = &v
	}

	err = s.publisher.Publish(ctx, s.rabbitmqExchange, s.rabbitmqRoutingKey, job)
	if err != nil {
		return uuid.Nil, "", fmt.Errorf("failed to publish optimization job: %w", err)
	}

	return jobID, "processing", nil
}

func (s *OptimizationService) generateDestinationKey(jobID uuid.UUID) string {
	timestamp := time.Now().Unix()
	return fmt.Sprintf("optimized/%d-%s.glb", timestamp, jobID.String())
}

func (s *OptimizationService) extractOrGenerateSourceURL(ctx context.Context, sourceURL string) (string, error) {
	if isPresignedURL(sourceURL) {
		return sourceURL, nil
	}

	return s.storage.GeneratePresignedDownloadURL(ctx, sourceURL)
}

func isPresignedURL(url string) bool {
	if url == "" {
		return false
	}

	if !strings.HasPrefix(url, "https://") && !strings.HasPrefix(url, "http://") {
		return false
	}

	return strings.Contains(url, "X-Amz-Algorithm") || strings.Contains(url, "Signature")
}
