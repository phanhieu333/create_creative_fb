package routes

import (
	"creative_fb/internal/dto"
	"creative_fb/internal/facebook"
	"creative_fb/internal/services"
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
	var req dto.CreateCreativeRequest

	fmt.Printf("Request body: %+v\n", req)

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreativeResponse{
			Success: false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	input := facebook.CreativeInput{
		Type:                      req.Type,
		AccountID:                 req.AccountID,
		PageID:                    req.PageID,
		Name:                      req.Name,
		ImageHash:                 req.ImageHash,
		Link:                      req.Link,
		Picture:                   req.Picture,
		Message:                   req.Message,
		VideoID:                   req.VideoID,
		Thumbnail:                 req.Thumbnail,
		AdvantageOptimizeCreative: req.AdvantageOptimizeCreative,
		Features:                  req.Features,
		ChildAttachmentsInput:     req.ChildAttachments,
	}

	result, err := cr.service.CreateCreative(input)
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
