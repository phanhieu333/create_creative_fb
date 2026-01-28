package services

import (
	"creative_fb/internal/dto"
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

func (s *CreativeService) CreateCreative(req dto.CreateCreativeRequest) (*dto.CreateCreativeResponse, error) {
	// fmt.Printf("Input: %+v\n", *req.ObjectStory)
	// if err := s.validator.ValidateInput(req); err != nil {
	// 	return nil, err
	// }
	creative := dto.CreateCreativeRequest{
		ObjectStory:          req.ObjectStory,
		DegreesOfFreedomSpec: req.DegreesOfFreedomSpec,
	}
	return s.repo.CreateCreative(req.AccountID, creative)

}

// func (s *CreativeService) createSingleImage(req dto.CreateCreativeRequest) (*dto.CreateCreativeResponse, error) {
// 	creative := dto.CreateCreativeRequest{
// 		ObjectStory:          req.ObjectStory,
// 		DegreesOfFreedomSpec: req.DegreesOfFreedomSpec,
// 	}

// 	// creative.DegreesOfFreedomSpec = s.buildDegreesOfFreedomSpec(req)

// 	return s.repo.CreateSingleImage(req.AccountID, creative)
// }

// func (s *CreativeService) createSingleVideo(req dto.CreateCreativeRequest) (*dto.CreateCreativeResponse, error) {
// 	creative := dto.CreateCreativeRequest{
// 		ObjectStory:          req.ObjectStory,
// 		DegreesOfFreedomSpec: req.DegreesOfFreedomSpec,
// 	}

// 	return s.repo.CreateSingleVideo(req.AccountID, creative)
// }

// func (s *CreativeService) createCarousel(req dto.CreativeBase) (*dto.CreateCreativeResponse, error) {
// 	creative := dto.CreateCreativeCarouselRequest{
// 		ObjectStorySpecCarousel: model.CarouselDetail{
// 			PageID: input.PageID,
// 			LinkData: model.LinkDataCarousel{
// 				Link:             input.Link,
// 				ChildAttachments: input.ChildAttachmentsInput,
// 			},
// 		},
// 	}

// 	creative.DegreesOfFreedomSpec = s.buildDegreesOfFreedomSpec(input)

// 	return s.repo.CreateCarousel(input.AccountID, creative)
// }

// func (s *CreativeService) createFlexible(input facebook.CreativeInput) (*dto.CreateCreativeResponse, error) {
// 	creative := dto.CreateCreativeFlexibleRequest{
// 		ObjectStorySpecFlexible: model.FlexibleDetail{
// 			PageID: input.PageID,
// 		},
// 		AssetFeedSpec: model.AssetFeedSpec{
// 			AdFormats: []string{"SINGLE_IMAGE", "SINGLE_VIDEO"},
// 		},
// 	}

// 	creative.DegreesOfFreedomSpec = s.buildDegreesOfFreedomSpec(input)

// 	return s.repo.CreateFlexible(input.AccountID, creative)
// }
