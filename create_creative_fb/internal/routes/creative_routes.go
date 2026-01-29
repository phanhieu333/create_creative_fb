package routes

import (
	"creative_fb/internal/dto"
	"creative_fb/internal/facebook"
	"creative_fb/internal/services"
	"encoding/json"
	"fmt"

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
	var in facebook.CreateCreativeInput

	if err := c.BodyParser(&in); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreativeResponse{
			Success: false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}
	creative := facebook.MapInputToDTO(in)
	b, _ := json.MarshalIndent(creative, "", "  ")
	fmt.Println(string(b))
	result, err := cr.service.CreateCreative(creative)
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
