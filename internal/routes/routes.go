package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kodra-pay/reconciliation-service/internal/handlers"
	"github.com/kodra-pay/reconciliation-service/internal/services"
)

func Register(app *fiber.App, service string) {
	health := handlers.NewHealthHandler(service)
	health.Register(app)

	svc := services.NewReconciliationService()
	h := handlers.NewReconciliationHandler(svc)
	api := app.Group("/recon")
	api.Post("/runs", h.CreateRun)
	api.Get("/runs/:id", h.GetRun)
}
