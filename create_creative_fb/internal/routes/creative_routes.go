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

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreativeResponse{
			Success: false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	fmt.Printf("Request body: %+v\n", req)

	// Map incoming request to internal facebook.CreativeInput
	input := facebook.CreativeInput{
		Type:                  req.Type,
		AccountID:             req.AccountID,
		PageID:                req.PageID,
		Features:              req.Features,
		ChildAttachmentsInput: req.ChildAttachments,
	}

	if req.ObjectStory != nil {
		if req.ObjectStory.PageID != "" {
			input.PageID = req.ObjectStory.PageID
		}
		if req.ObjectStory.LinkData != nil {
			input.Link = req.ObjectStory.LinkData.Link
			input.Message = req.ObjectStory.LinkData.Message
			input.ImageHash = req.ObjectStory.LinkData.ImageHash
			// CTA type is available at req.ObjectStory.LinkData.CallToAction.Type if needed
		}
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
