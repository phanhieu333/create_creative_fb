package facebook

import "creative_fb/internal/dto"

type CreateCreativeInput struct {
	AccountID        string            `json:"account_id"`
	PageID           string            `json:"page_id,omitempty"`
	Type             string            `json:"type,omitempty"`
	Link             string            `json:"link,omitempty"`
	Message          string            `json:"message,omitempty"`
	Title            string            `json:"title,omitempty"`
	Description      string            `json:"description,omitempty"`
	ImageHash        string            `json:"image_hash,omitempty"`
	ImageURL         string            `json:"image_url,omitempty"`
	VideoID          string            `json:"video_id,omitempty"`
	CallToAction     *dto.CallToAction `json:"call_to_action,omitempty"`
	ChildAttachments []struct {
		Link         string            `json:"link,omitempty"`
		VideoID      string            `json:"video_id,omitempty"`
		Name         string            `json:"name,omitempty"`
		Description  string            `json:"description,omitempty"`
		ImageHash    string            `json:"image_hash,omitempty"`
		CallToAction *dto.CallToAction `json:"call_to_action,omitempty"`
	} `json:"child_attachments,omitempty"`
}

// mapping từ input -> dto.CreateCreativeRequest
func MapInputToDTO(in CreateCreativeInput) dto.CreateCreativeRequest {
	cr := dto.CreateCreativeRequest{
		AccountID: in.AccountID,
		ObjectStory: &dto.ObjectStorySpec{
			PageID: in.PageID,
		},
	}

	switch in.Type {
	case "single_video":
		cr.ObjectStory.VideoData = &dto.VideoData{
			VideoID:      in.VideoID,
			ImageURL:     in.ImageURL,
			Title:        in.Title,
			Message:      in.Message,
			Description:  in.Description,
			CallToAction: in.CallToAction,
		}
	default:
		ld := &dto.LinkData{
			Link:         in.Link,
			Message:      in.Message,
			Description:  in.Description,
			ImageHash:    in.ImageHash,
			CallToAction: in.CallToAction,
		}
		for _, ca := range in.ChildAttachments {
			ld.ChildAttachments = append(ld.ChildAttachments, dto.ChildAttachmentsData{
				Link:         ca.Link,
				VideoID:      ca.VideoID,
				Name:         ca.Name,
				Description:  ca.Description,
				ImageHash:    ca.ImageHash,
				CallToAction: ca.CallToAction,
			})
		}
		cr.ObjectStory.LinkData = ld
	}

	return cr
}
