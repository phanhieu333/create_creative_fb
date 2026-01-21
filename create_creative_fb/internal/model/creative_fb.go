package model

type LinkData struct {
	Link        string `json:"link"`
	Picture     string `json:"picture,omitempty"`
	ImageHash   string `json:"image_hash,omitempty"`
	Name        string `json:"name,omitempty"`
	Message     string `json:"message,omitempty"`
	Description string `json:"description,omitempty"`
}

type ImageDetail struct {
	PageID   string   `json:"page_id"`
	LinkData LinkData `json:"link_data"`
}

type CallToAction struct {
	Type  string                 `json:"type"`
	Value map[string]interface{} `json:"value"`
}

type VideoData struct {
	CallToAction *CallToAction `json:"call_to_action,omitempty"`
	ImageURL     string        `json:"image_url"`
	VideoID      string        `json:"video_id"`
}

type VideoDetail struct {
	PageID    string    `json:"page_id"`
	VideoData VideoData `json:"video_data"`
}

type ChildAttachmentsData struct {
	Link        string `json:"link"`
	VideoID     string `json:"video_id"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ImageHash   string `json:"image_hash"`
}

type LinkDataCarousel struct {
	Link             string                 `json:"link"`
	ImageHash        string                 `json:"image_hash,omitempty"`
	Picture          string                 `json:"picture,omitempty"`
	ChildAttachments []ChildAttachmentsData `json:"child_attachments"`
}

type CarouselDetail struct {
	PageID   string           `json:"page_id"`
	LinkData LinkDataCarousel `json:"link_data"`
}
