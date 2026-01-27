package services

import (
	"fmt"

	"creative_fb/internal/dto"
	"creative_fb/internal/facebook"
	"creative_fb/internal/model"
	"creative_fb/internal/repositories"
)

// CreativeService handles business logic for creative creation
type CreativeService struct {
	repo      *repositories.CreativeRepository
	validator *CreativeValidator
}

func NewCreativeService(repo *repositories.CreativeRepository) *CreativeService {
	return &CreativeService{
		repo:      repo,
		validator: NewCreativeValidator(),
	}
}

// CreateCreative dispatches creative creation based on type
func (s *CreativeService) CreateCreative(input facebook.CreativeInput) (*dto.CreateCreativeResponse, error) {
	// Validate input
	if err := s.validator.ValidateInput(input); err != nil {
		return nil, err
	}

	switch input.Type {
	case "single_image":
		return s.createSingleImage(input)
	case "single_video":
		return s.createSingleVideo(input)
	case "carousel":
		return s.createCarousel(input)
	case "flexible":
		return s.createFlexible(input)
	default:
		return nil, fmt.Errorf("unsupported creative type: %s", input.Type)
	}
}

func (s *CreativeService) createSingleImage(input facebook.CreativeInput) (*dto.CreateCreativeResponse, error) {
	creative := dto.CreateCreativeSingleImageRequest{
		ObjectStorySpecImage: model.ImageDetail{
			PageID: input.PageID,
			LinkData: model.LinkData{
				Link: input.Link,
			},
		},
	}

	if input.AdvantageOptimizeCreative == "on" {
		creative.DegreesOfFreedomSpec = s.buildDegreesOfFreedomSpec(input)
	}

	return s.repo.CreateSingleImage(input.AccountID, creative)
}

func (s *CreativeService) createSingleVideo(input facebook.CreativeInput) (*dto.CreateCreativeResponse, error) {
	creative := dto.CreateCreativeSingleVideoRequest{
		ObjectStorySpecVideo: model.VideoDetail{
			PageID: input.PageID,
			VideoData: model.VideoData{
				VideoID:  input.VideoID,
				ImageURL: input.Thumbnail,
			},
		},
	}

	if input.AdvantageOptimizeCreative == "on" {
		creative.DegreesOfFreedomSpec = s.buildDegreesOfFreedomSpec(input)
	}

	return s.repo.CreateSingleVideo(input.AccountID, creative)
}

func (s *CreativeService) createCarousel(input facebook.CreativeInput) (*dto.CreateCreativeResponse, error) {
	creative := dto.CreateCreativeCarouselRequest{
		ObjectStorySpecCarousel: model.CarouselDetail{
			PageID: input.PageID,
			LinkData: model.LinkDataCarousel{
				Link:             input.Link,
				ChildAttachments: input.ChildAttachmentsInput,
			},
		},
	}

	if input.AdvantageOptimizeCreative == "on" {
		creative.DegreesOfFreedomSpec = s.buildDegreesOfFreedomSpec(input)
	}

	return s.repo.CreateCarousel(input.AccountID, creative)
}

func (s *CreativeService) createFlexible(input facebook.CreativeInput) (*dto.CreateCreativeResponse, error) {
	creative := dto.CreateCreativeFlexibleRequest{
		ObjectStorySpecFlexible: model.FlexibleDetail{
			PageID: input.PageID,
		},
		AssetFeedSpec: model.AssetFeedSpec{
			AdFormats: []string{"SINGLE_IMAGE", "SINGLE_VIDEO"},
		},
	}

	if input.AdvantageOptimizeCreative == "on" {
		creative.DegreesOfFreedomSpec = s.buildDegreesOfFreedomSpec(input)
	}

	return s.repo.CreateFlexible(input.AccountID, creative)
}

func (s *CreativeService) buildDegreesOfFreedomSpec(input facebook.CreativeInput) *model.DegreesOfFreedomSpec {
	spec := &model.DegreesOfFreedomSpec{
		CreativeFeaturesSpec: model.CreativeFeaturesSpec{},
	}

	enrollStatus := "OPT_IN"
	if val, exists := input.Features["enroll_status"]; exists {
		enrollStatus = val
	}

	switch input.Type {
	case "single_image", "single_video":
		// For single asset: toggle features individually
		singleAssetFeatures := []string{
			"advantage_plus_creative",
			"cv_transformation",
			"enhance_cta",
			"image_animation",
			"image_brightness_and_contrast",
			"image_templates",
			"image_touchups",
			"inline_comment",
			"pac_relaxation",
			"product_extensions",
			"reveal_details_over_time",
			"show_summary",
			"site_extensions",
			"text_optimizations",
			"text_translation",
		}

		for _, feature := range singleAssetFeatures {
			if val, exists := input.Features[feature]; exists && (val == "on" || val == "true") {
				spec.CreativeFeaturesSpec[feature] = &model.CreativeFeatureEnrollment{
					EnrollStatus: enrollStatus,
				}
			}
		}

	case "carousel", "flexible":
		// For carousel/flexible: default OPT_IN for all features
		allFeatures := []string{
			"advantage_plus_creative",
			"cv_transformation",
			"enhance_cta",
			"image_animation",
			"image_brightness_and_contrast",
			"image_templates",
			"image_touchups",
			"inline_comment",
			"pac_relaxation",
			"product_extensions",
			"reveal_details_over_time",
			"show_summary",
			"site_extensions",
			"text_optimizations",
			"text_translation",
		}

		for _, feature := range allFeatures {
			spec.CreativeFeaturesSpec[feature] = &model.CreativeFeatureEnrollment{
				EnrollStatus: "OPT_IN",
			}
		}
	}

	return spec
}
