package services

import (
	"fmt"

	"creative_fb/internal/facebook"
)

// CreativeValidator validates creative input
type CreativeValidator struct{}

func NewCreativeValidator() *CreativeValidator {
	return &CreativeValidator{}
}

func (v *CreativeValidator) ValidateInput(input facebook.CreativeInput) error {
	if input.Type == "" {
		return fmt.Errorf("missing -type flag")
	}

	if input.AccountID == "" || input.PageID == "" {
		return fmt.Errorf("missing -account or -page flag")
	}

	switch input.Type {
	case "single_video":
		return v.validateSingleVideo(input)
	case "carousel":
		return v.validateCarousel(input)
	case "single_image", "flexible":
		// single_image and flexible have minimal validation
		return nil
	default:
		return fmt.Errorf("unsupported creative type: %s", input.Type)
	}
}

func (v *CreativeValidator) validateSingleVideo(input facebook.CreativeInput) error {
	if input.VideoID == "" {
		return fmt.Errorf("missing -video flag for single_video type")
	}
	if input.Thumbnail == "" {
		return fmt.Errorf("missing -thumbnail flag for single_video type")
	}
	return nil
}

func (v *CreativeValidator) validateCarousel(input facebook.CreativeInput) error {
	// 1. Carousel must have at least 2 cards
	if len(input.ChildAttachmentsInput) < 2 {
		return fmt.Errorf("carousel requires at least 2 child_attachments")
	}

	// 2. Check media for each child
	allChildHasMedia := true
	for i, child := range input.ChildAttachmentsInput {
		if child.ImageHash == "" {
			allChildHasMedia = false
			return fmt.Errorf("child attachment %d missing image_hash", i+1)
		}
	}

	// 3. Check fallback image
	hasImage := input.ImageHash != ""
	hasPicture := input.Picture != ""

	if hasImage && hasPicture {
		return fmt.Errorf("provide only one of -image or -picture for carousel fallback")
	}

	// 4. Rule: if child is missing media, need fallback image
	if !allChildHasMedia && !hasImage && !hasPicture {
		return fmt.Errorf("carousel with missing media requires fallback image (-image or -picture)")
	}

	return nil
}
