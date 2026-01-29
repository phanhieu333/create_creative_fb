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
