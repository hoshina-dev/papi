package service

import (
	"context"
	"fmt"
	"log"
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
	model3DRepo        repository.Model3DRepository
	webhookURL         string
	rabbitmqExchange   string
	rabbitmqRoutingKey string
}

func NewOptimizationService(
	storage storage.StorageService,
	publisher rabbitmq.Publisher,
	model3DRepo repository.Model3DRepository,
	webhookURL string,
	rabbitmqExchange string,
	rabbitmqRoutingKey string,
) *OptimizationService {
	return &OptimizationService{
		storage:            storage,
		publisher:          publisher,
		model3DRepo:        model3DRepo,
		webhookURL:         webhookURL,
		rabbitmqExchange:   rabbitmqExchange,
		rabbitmqRoutingKey: rabbitmqRoutingKey,
	}
}

func (s *OptimizationService) Optimize3D(ctx context.Context, params model.Optimize3DInput) (jobID uuid.UUID, status string, err error) {
	// 	Optional (Optimization)
	// | 			Variable	   	 | Default | Range | 			Description				 |
	// |-----------------------------|---------|-------|-------------------------------------|
	// | DRACO_COMPRESSION_LEVEL 	 | 	  10   |  0-10 | Compression level (higher = smaller)|
	// | DRACO_POSITION_QUANTIZATION |    14   |  8-30 | Vertex position precision 			 |
	// | DRACO_TEXCOORD_QUANTIZATION | 	  12   |  8-30 | UV coordinate precision 			 |
	// | DRACO_NORMAL_QUANTIZATION   | 	  10   |  8-30 | Surface normal precision 			 |
	// | DRACO_GENERIC_QUANTIZATION  | 	  8    |  8-30 | Other attributes precision 		 |
	dracoCompressionLevel, err := validateDracoParam(params.DracoCompressionLevel, 0, 10, 10)
	if err != nil {
		return uuid.Nil, "", err
	}
	dracoPositionQuantization, err := validateDracoParam(params.DracoPositionQuantization, 8, 30, 14)
	if err != nil {
		return uuid.Nil, "", err
	}
	dracoTexcoordQuantization, err := validateDracoParam(params.DracoTexcoordQuantization, 8, 30, 12)
	if err != nil {
		return uuid.Nil, "", err
	}
	dracoNormalQuantization, err := validateDracoParam(params.DracoNormalQuantization, 8, 30, 10)
	if err != nil {
		return uuid.Nil, "", err
	}
	dracoGenericQuantization, err := validateDracoParam(params.DracoGenericQuantization, 8, 30, 8)
	if err != nil {
		return uuid.Nil, "", err
	}

	jobID = uuid.New()

	destKey := s.generateDestinationKey(jobID)

	destURL, err := s.storage.GeneratePresignedUploadURL(ctx, destKey, "model/gltf-binary", storage.JobPresignTTL)
	if err != nil {
		return uuid.Nil, "", fmt.Errorf("failed to generate destination URL: %w", err)
	}

	sourcePresignedURL, err := s.extractOrGenerateSourceURL(ctx, params.SourceURL)
	if err != nil {
		return uuid.Nil, "", fmt.Errorf("failed to process source URL: %w", err)
	}

	// Exactly one of PartID or ProductID is set
	if (params.PartID == nil) == (params.ProductID == nil) {
		return uuid.Nil, "", fmt.Errorf("exactly one of PartID or ProductID must be provided")
	}

	model3D := &model.Model3D{
		ID:           jobID,
		PartID:       params.PartID,
		ProductID:    params.ProductID,
		RawKey:       params.SourceURL,
		ProcessedKey: &destKey,
		Status:       model.Model3DStatusProcessing,
	}

	err = s.model3DRepo.Create(ctx, model3D)
	if err != nil {
		return uuid.Nil, "", fmt.Errorf("failed to create 3D model record: %w", err)
	}

	job := model.OptimizationJob{
		UUID:                      jobID.String(),
		SourceGLMURL:              sourcePresignedURL,
		DestGLMURL:                destURL,
		WebhookURL:                s.webhookURL,
		DracoCompressionLevel:     dracoCompressionLevel,
		DracoPositionQuantization: dracoPositionQuantization,
		DracoTexcoordQuantization: dracoTexcoordQuantization,
		DracoNormalQuantization:   dracoNormalQuantization,
		DracoGenericQuantization:  dracoGenericQuantization,
	}

	err = s.publisher.Publish(ctx, s.rabbitmqExchange, s.rabbitmqRoutingKey, job)
	if err != nil {
		if delErr := s.model3DRepo.Delete(ctx, jobID); delErr != nil {
			log.Printf("failed to delete orphaned 3D model record %s after publish failure: %v", jobID, delErr)
		}
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

	return s.storage.GeneratePresignedDownloadURL(ctx, sourceURL, storage.JobPresignTTL)
}

func validateDracoParam(v *int32, min int, max int, defaultValue int) (*int, error) {
	i := defaultValue
	if v != nil {
		i = int(*v)
	}
	if i < min || i > max {
		return nil, fmt.Errorf("value must be between %d and %d", min, max)
	}
	return &i, nil
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

func (s *OptimizationService) GetModel3DResult(ctx context.Context, jobID uuid.UUID) (*model.Model3DResult, error) {
	m, err := s.model3DRepo.GetByID(ctx, jobID)
	if err != nil {
		return nil, fmt.Errorf("failed to get 3D model: %w", err)
	}
	return s.toModel3DResult(ctx, *m)
}

func (s *OptimizationService) GetModel3DByPartID(ctx context.Context, partID uuid.UUID) ([]*model.Model3DResult, error) {
	models, err := s.model3DRepo.GetByPartID(ctx, partID)
	if err != nil {
		return nil, fmt.Errorf("failed to get 3D model for part: %w", err)
	}
	return s.toModel3DResults(ctx, models)
}

func (s *OptimizationService) GetModel3DByProductID(ctx context.Context, productID uuid.UUID) ([]*model.Model3DResult, error) {
	models, err := s.model3DRepo.GetByProductID(ctx, productID)
	if err != nil {
		return nil, fmt.Errorf("failed to get 3D model for product: %w", err)
	}
	return s.toModel3DResults(ctx, models)
}

func (s *OptimizationService) toModel3DResults(ctx context.Context, models []model.Model3D) ([]*model.Model3DResult, error) {
	results := make([]*model.Model3DResult, len(models))
	for i, m := range models {
		result, err := s.toModel3DResult(ctx, m)
		if err != nil {
			return nil, err
		}
		results[i] = result
	}
	return results, nil
}

func (s *OptimizationService) toModel3DResult(ctx context.Context, m model.Model3D) (*model.Model3DResult, error) {
	result := &model.Model3DResult{
		JobID:     m.ID,
		PartID:    m.PartID,
		ProductID: m.ProductID,
		Status:    string(m.Status),
	}
	if m.Status == model.Model3DStatusReady && m.ProcessedKey != nil {
		url, err := s.storage.GeneratePresignedDownloadURL(ctx, *m.ProcessedKey, storage.ClientPresignTTL)
		if err != nil {
			return nil, fmt.Errorf("failed to generate download URL for job %s: %w", m.ID, err)
		}
		result.DownloadURL = &url
	}
	return result, nil
}
