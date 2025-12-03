package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/kodra-pay/reconciliation-service/internal/dto"
	"github.com/kodra-pay/reconciliation-service/internal/services"
)

type ReconciliationHandler struct {
	svc *services.ReconciliationService
}

func NewReconciliationHandler(svc *services.ReconciliationService) *ReconciliationHandler {
	return &ReconciliationHandler{svc: svc}
}

func (h *ReconciliationHandler) CreateRun(c *fiber.Ctx) error {
	var req dto.ReconRunRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}
	return c.JSON(h.svc.CreateRun(c.Context(), req))
}

func (h *ReconciliationHandler) GetRun(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(h.svc.GetRun(c.Context(), id))
}
