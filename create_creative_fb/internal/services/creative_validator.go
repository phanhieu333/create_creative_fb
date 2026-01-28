package services

import (
	"creative_fb/internal/dto"
	"fmt"
)

// CreativeValidator validates creative input
type CreativeValidator struct{}

func NewCreativeValidator() *CreativeValidator {
	return &CreativeValidator{}
}

func (v *CreativeValidator) ValidateInput(req dto.CreateCreativeRequest) error {
	if req.ObjectStory == nil {
		return fmt.Errorf("missing object_story_spec")
	}

	if req.AccountID == "" {
		return fmt.Errorf("missing -account or -page flag")
	}

	return nil

}

// func (v *CreativeValidator) validateCarousel(input facebook.CreativeInput) error {
// 	// 1. Carousel must have at least 2 cards
// 	if len(input.ChildAttachmentsInput) < 2 {
// 		return fmt.Errorf("carousel requires at least 2 child_attachments")
// 	}

// 	// 2. Check media for each child
// 	allChildHasMedia := true
// 	for i, child := range input.ChildAttachmentsInput {
// 		if child.ImageHash == "" {
// 			allChildHasMedia = false
// 			return fmt.Errorf("child attachment %d missing image_hash", i+1)
// 		}
// 	}

// 	// 3. Check fallback image
// 	hasImage := input.ImageHash != ""
// 	hasPicture := input.Picture != ""

// 	if hasImage && hasPicture {
// 		return fmt.Errorf("provide only one of -image or -picture for carousel fallback")
// 	}

// 	// 4. Rule: if child is missing media, need fallback image
// 	if !allChildHasMedia && !hasImage && !hasPicture {
// 		return fmt.Errorf("carousel with missing media requires fallback image (-image or -picture)")
// 	}

// 	return nil
// }
