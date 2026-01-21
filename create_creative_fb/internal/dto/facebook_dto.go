package dto

import (
	"creative_fb/internal/model"
)

type CreateCreativeSingleImageRequest struct {
	ObjectStorySpecImage model.ImageDetail `json:"object_story_spec"`
}

type CreateCreativeSingleVideoRequest struct {
	ObjectStorySpecVideo model.VideoDetail `json:"object_story_spec"`
}

type CreateCreativeCarouselRequest struct {
	ObjectStorySpecCarousel model.CarouselDetail `json:"object_story_spec"`
}

type CreateCreativeResponse struct {
	ID string `json:"id"`
}
