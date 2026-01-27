package dto

import (
	"creative_fb/internal/model"
)

// CreateCreativeRequest represents the incoming HTTP request body.
// It supports the nested `object_story_spec` structure like the example
// provided by the user.
type CreateCreativeRequest struct {
	Type        string            `json:"type"`
	AccountID   string            `json:"account_id"`
	PageID      string            `json:"page_id"`
	ObjectStory *ObjectStorySpec  `json:"object_story_spec,omitempty"`
	Features    map[string]string `json:"features,omitempty"`
	// Keep child attachments if present for carousel
	ChildAttachments []model.ChildAttachmentsData `json:"child_attachments,omitempty"`
}

type ObjectStorySpec struct {
	PageID   string    `json:"page_id,omitempty"`
	LinkData *LinkData `json:"link_data,omitempty"`
	// video_data and other variants can be added here later
}

type LinkData struct {
	Link         string        `json:"link,omitempty"`
	Message      string        `json:"message,omitempty"`
	Description  string        `json:"description,omitempty"`
	ImageHash    string        `json:"image_hash,omitempty"`
	CallToAction *CallToAction `json:"call_to_action,omitempty"`
}

type CallToAction struct {
	Type string `json:"type,omitempty"`
}

type CreativeResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	ID      string `json:"id"`
	Error   string `json:"error,omitempty"`
}
