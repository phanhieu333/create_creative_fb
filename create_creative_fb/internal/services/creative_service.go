package services

import (
	"fmt"

	"creative_fb/internal/dto"
	"creative_fb/internal/facebook"
	"creative_fb/internal/model"
	"creative_fb/internal/repositories"
)

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

func (s *CreativeService) CreateCreative(input facebook.CreativeInput) (*dto.CreateCreativeResponse, error) {
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

	creative.DegreesOfFreedomSpec = s.buildDegreesOfFreedomSpec(input)

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

	creative.DegreesOfFreedomSpec = s.buildDegreesOfFreedomSpec(input)

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

	creative.DegreesOfFreedomSpec = s.buildDegreesOfFreedomSpec(input)

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

	creative.DegreesOfFreedomSpec = s.buildDegreesOfFreedomSpec(input)

	return s.repo.CreateFlexible(input.AccountID, creative)
}

func (s *CreativeService) buildDegreesOfFreedomSpec(input facebook.CreativeInput) *model.DegreesOfFreedomSpec {
	spec := &model.DegreesOfFreedomSpec{
		CreativeFeaturesSpec: model.CreativeFeaturesSpec{},
	}

	switch input.Type {
	case "single_image", "single_video":
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
			if status, exists := input.Features[feature]; exists {
				spec.CreativeFeaturesSpec[feature] = &model.CreativeFeatureEnrollment{
					EnrollStatus: status,
				}
			}
		}

	case "flexible":
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

	fmt.Printf("spec %v \n", spec)

	return spec
}
