package services

import (
	"context"

	"github.com/kodra-pay/reconciliation-service/internal/dto"
)

type ReconciliationService struct{}

func NewReconciliationService() *ReconciliationService { return &ReconciliationService{} }

func (s *ReconciliationService) CreateRun(_ context.Context, req dto.ReconRunRequest) dto.ReconRunResponse {
	// In a real scenario, this would generate a unique int ID.
	// For this mock implementation, we return a placeholder.
	return dto.ReconRunResponse{
		ID:     1, // Placeholder for an auto-generated int ID
		Source: req.Source,
		Status: "pending",
	}
}

func (s *ReconciliationService) GetRun(_ context.Context, id int) dto.ReconRunResponse {
	return dto.ReconRunResponse{
		ID:     id,
		Source: "unknown",
		Status: "pending",
	}
}
