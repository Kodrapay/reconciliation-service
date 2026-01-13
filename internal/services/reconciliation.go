package services

import (
	"context"
	"errors"
	"time"

	"github.com/kodra-pay/reconciliation-service/internal/dto"
	"github.com/kodra-pay/reconciliation-service/internal/models"
	"github.com/kodra-pay/reconciliation-service/internal/repositories"
)

type ReconciliationService struct {
	repo repositories.RunRepository
}

func NewReconciliationService(repo repositories.RunRepository) *ReconciliationService {
	return &ReconciliationService{repo: repo}
}

// CreateRun kicks off a reconciliation execution and stores the run metadata.
func (s *ReconciliationService) CreateRun(_ context.Context, req dto.ReconRunRequest) (dto.ReconRunResponse, error) {
	if req.Source == "" {
		return dto.ReconRunResponse{}, errors.New("source is required")
	}

	run := &models.ReconciliationRun{
		Source:    req.Source,
		Date:      req.Date,
		Status:    "completed", // placeholder; future: async processing pipeline
		Matched:   0,
		Unmatched: 0,
		CreatedAt: time.Now(),
	}

	created, err := s.repo.Create(run)
	if err != nil {
		return dto.ReconRunResponse{}, err
	}

	return toResponse(created), nil
}

func (s *ReconciliationService) GetRun(_ context.Context, id int) (dto.ReconRunResponse, error) {
	run, err := s.repo.Get(id)
	if err != nil {
		return dto.ReconRunResponse{}, err
	}
	return toResponse(run), nil
}

func (s *ReconciliationService) ListRuns(_ context.Context) ([]dto.ReconRunResponse, error) {
	runs, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	resp := make([]dto.ReconRunResponse, 0, len(runs))
	for _, run := range runs {
		resp = append(resp, toResponse(run))
	}
	return resp, nil
}

func toResponse(run *models.ReconciliationRun) dto.ReconRunResponse {
	return dto.ReconRunResponse{
		ID:        run.ID,
		Source:    run.Source,
		Date:      run.Date,
		Status:    run.Status,
		Matched:   run.Matched,
		Unmatched: run.Unmatched,
		CreatedAt: run.CreatedAt.Format(time.RFC3339),
	}
}
