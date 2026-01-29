package dto

type CreateCreativeRequest struct {
	AccountID            string                `json:"account_id"`
	ObjectStory          *ObjectStorySpec      `json:"object_story_spec,omitempty"`
	DegreesOfFreedomSpec *DegreesOfFreedomSpec `json:"degrees_of_freedom_spec,omitempty"`
}

type ObjectStorySpec struct {
	PageID    string     `json:"page_id,omitempty"`
	LinkData  *LinkData  `json:"link_data,omitempty"`
	VideoData *VideoData `json:"video_data,omitempty"`
}

type DegreesOfFreedomSpec struct {
	CreativeFeaturesSpec map[string]*CreativeFeatureEnrollment `json:"creative_features_spec"`
}

type CreativeFeatureEnrollment struct {
	EnrollStatus string `json:"enroll_status"`
}

type LinkData struct {
	Link             string                 `json:"link,omitempty"`
	Message          string                 `json:"message,omitempty"`
	Description      string                 `json:"description,omitempty"`
	ImageHash        string                 `json:"image_hash,omitempty"`
	CallToAction     *CallToAction          `json:"call_to_action,omitempty"`
	ChildAttachments []ChildAttachmentsData `json:"child_attachments,omitempty"`
}

type ChildAttachmentsData struct {
	Link         string        `json:"link"`
	VideoID      string        `json:"video_id,omitempty"`
	Name         string        `json:"name,omitempty"`
	Description  string        `json:"description,omitempty"`
	ImageHash    string        `json:"image_hash,omitempty"`
	CallToAction *CallToAction `json:"call_to_action,omitempty"`
}

type VideoData struct {
	VideoID      string        `json:"video_id"`
	ImageURL     string        `json:"image_url,omitempty"`
	Title        string        `json:"title,omitempty"`
	Message      string        `json:"message,omitempty"`
	Description  string        `json:"description,omitempty"`
	CallToAction *CallToAction `json:"call_to_action,omitempty"`
}

type CallToAction struct {
	Type  string             `json:"type,omitempty"`
	Value *CallToActionValue `json:"value,omitempty"`
}

type CallToActionValue struct {
	Link string `json:"link,omitempty"`
}

type CreativeResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	ID      string `json:"id"`
	Error   string `json:"error,omitempty"`
}
