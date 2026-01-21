package facebook

import (
	"log"
)

func ValidateInput(in CreativeInput) {
	if in.Type == "" {
		log.Fatal("Missing -type")
	}

	if in.AccountID == "" || in.PageID == "" {
		log.Fatal("Missing -account or -page")
	}

	if in.Type == "single_video" {
		if in.VideoID == "" {
			log.Fatal("Missing -video Id")
		}
		if in.Thumbnail == "" {
			log.Fatal("Missing -thumbnail")
		}
	}

	if in.Type == "carousel" {

		// 1. Carousel phải có ít nhất 2 card
		if len(in.ChildAttachmentsInput) < 2 {
			log.Fatal("carousel requires at least 2 child_attachments")
		}

		// 2. Check media của từng child
		allChildHasMedia := true
		for i, c := range in.ChildAttachmentsInput {
			if c.ImageHash == "" {
				allChildHasMedia = false
				log.Printf("Child attachment %d missing image_hash", i+1)
			}
		}

		// 3. Check fallback image
		hasImage := in.ImageHash != ""
		hasPicture := in.Picture != ""

		if hasImage && hasPicture {
			log.Fatal("Provide only one of -image or -picture for carousel fallback")
		}

		// 4. Rule Facebook: nếu có child thiếu media → cần fallback
		if !allChildHasMedia && !hasImage && !hasPicture {
			log.Fatal("Carousel with missing media in child attachments requires a fallback image (-image or -picture)")
		}
	}

}
