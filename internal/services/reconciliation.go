package services

import (
	"context"

	"github.com/google/uuid"

	"github.com/kodra-pay/reconciliation-service/internal/dto"
)

type ReconciliationService struct{}

func NewReconciliationService() *ReconciliationService { return &ReconciliationService{} }

func (s *ReconciliationService) CreateRun(_ context.Context, req dto.ReconRunRequest) dto.ReconRunResponse {
	return dto.ReconRunResponse{
		ID:     "recon_" + uuid.NewString(),
		Source: req.Source,
		Status: "pending",
	}
}

func (s *ReconciliationService) GetRun(_ context.Context, id string) dto.ReconRunResponse {
	return dto.ReconRunResponse{
		ID:     id,
		Source: "unknown",
		Status: "pending",
	}
}
