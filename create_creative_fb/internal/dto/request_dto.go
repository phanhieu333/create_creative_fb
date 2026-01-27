package dto

import (
	"creative_fb/internal/model"
)

// CreateCreativeRequest is the API request body for creating a creative
type CreateCreativeRequest struct {
	Type                      string                       `json:"type"`
	AccountID                 string                       `json:"account_id"`
	PageID                    string                       `json:"page_id"`
	Name                      string                       `json:"name"`
	ImageHash                 string                       `json:"image_hash"`
	Link                      string                       `json:"link"`
	Picture                   string                       `json:"picture"`
	Message                   string                       `json:"message"`
	VideoID                   string                       `json:"video_id"`
	Thumbnail                 string                       `json:"thumbnail"`
	AdvantageOptimizeCreative string                       `json:"advantage_optimize"`
	Features                  map[string]string            `json:"features"`
	ChildAttachments          []model.ChildAttachmentsData `json:"child_attachments"`
}

type CreativeResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	ID      string `json:"id"`
	Error   string `json:"error,omitempty"`
}
