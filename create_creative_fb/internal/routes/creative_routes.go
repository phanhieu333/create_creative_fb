package routes

import (
	"creative_fb/internal/dto"
	"creative_fb/internal/services"

	"github.com/gofiber/fiber/v2"
)

type CreativeRoutes struct {
	service *services.CreativeService
}

func NewCreativeRoutes(service *services.CreativeService) *CreativeRoutes {
	return &CreativeRoutes{
		service: service,
	}
}

func (cr *CreativeRoutes) RegisterRoutes(app *fiber.App) {
	creativeGroup := app.Group("/api/v1/creatives")
	{
		creativeGroup.Post("", cr.CreateCreative)
	}
}

func (cr *CreativeRoutes) CreateCreative(c *fiber.Ctx) error {
	var req dto.CreateCreativeRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreativeResponse{
			Success: false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	result, err := cr.service.CreateCreative(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreativeResponse{
			Success: false,
			Message: "Failed to create creative",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.CreativeResponse{
		Success: true,
		Message: "Creative created successfully",
		ID:      result.ID,
	})
}
