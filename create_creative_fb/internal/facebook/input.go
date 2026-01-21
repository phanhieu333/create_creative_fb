package facebook

import "creative_fb/internal/model"

type CreativeInput struct {
	Type                  string
	AccountID             string
	PageID                string
	Name                  string
	ImageHash             string
	Link                  string
	Picture               string
	Message               string
	VideoID               string
	Thumbnail             string
	PictureURL            string
	ChildAttachmentsInput []model.ChildAttachmentsData
}
