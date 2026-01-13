package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/kodra-pay/reconciliation-service/internal/handlers"
	"github.com/kodra-pay/reconciliation-service/internal/repositories"
	"github.com/kodra-pay/reconciliation-service/internal/services"
)

func Register(app *fiber.App, serviceName string) {
	health := handlers.NewHealthHandler(serviceName)
	health.Register(app)

	// Core reconciliation endpoints
	repo := repositories.NewInMemoryRunRepository()
	svc := services.NewReconciliationService(repo)
	handler := handlers.NewReconciliationHandler(svc)

	app.Get("/reconciliation", handler.ListRuns)
	app.Get("/reconciliation/:id", handler.GetRun)
	app.Post("/reconciliation/process", handler.CreateRun)
}
