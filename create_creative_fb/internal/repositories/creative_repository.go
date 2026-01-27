package repositories

import (
	"creative_fb/internal/dto"
	"creative_fb/internal/facebook"
)

// CreativeRepository handles all creative-related Facebook API operations
type CreativeRepository struct {
	client *facebook.Client
}

func NewCreativeRepository(client *facebook.Client) *CreativeRepository {
	return &CreativeRepository{
		client: client,
	}
}

func (r *CreativeRepository) CreateSingleImage(accountID string, req dto.CreateCreativeSingleImageRequest) (*dto.CreateCreativeResponse, error) {
	return r.client.CreateCreativeSingleImage(accountID, req)
}

func (r *CreativeRepository) CreateSingleVideo(accountID string, req dto.CreateCreativeSingleVideoRequest) (*dto.CreateCreativeResponse, error) {
	return r.client.CreateCreativeSingleVideo(accountID, req)
}

func (r *CreativeRepository) CreateCarousel(accountID string, req dto.CreateCreativeCarouselRequest) (*dto.CreateCreativeResponse, error) {
	return r.client.CreateCreativeCarousel(accountID, req)
}

func (r *CreativeRepository) CreateFlexible(accountID string, req dto.CreateCreativeFlexibleRequest) (*dto.CreateCreativeResponse, error) {
	return r.client.CreateCreativeFlexible(accountID, req)
}
