package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/kodra-pay/reconciliation-service/internal/dto"
	"github.com/kodra-pay/reconciliation-service/internal/repositories"
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
	run, err := h.svc.CreateRun(c.Context(), req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(run)
}

func (h *ReconciliationHandler) GetRun(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id") // Use c.ParamsInt
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid reconciliation run ID")
	}
	run, err := h.svc.GetRun(c.Context(), id)
	if err != nil {
		if err == repositories.ErrRunNotFound {
			return fiber.NewError(fiber.StatusNotFound, "reconciliation run not found")
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(run)
}

func (h *ReconciliationHandler) ListRuns(c *fiber.Ctx) error {
	runs, err := h.svc.ListRuns(c.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(runs)
}
