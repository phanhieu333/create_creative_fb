package handlers

import (
	"flag"
	"fmt"
	"strings"

	"creative_fb/internal/facebook"
	"creative_fb/internal/model"
)

type CreativeHandler struct{}

func NewCreativeHandler() *CreativeHandler {
	return &CreativeHandler{}
}

// arrayFlags implements flag.Value for parsing multiple flag values
type arrayFlags []string

func (i *arrayFlags) String() string {
	return strings.Join(*i, ",")
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

// ParseFlags parses command line flags into CreativeInput
func (h *CreativeHandler) ParseFlags() facebook.CreativeInput {
	input := facebook.CreativeInput{
		Features: make(map[string]string),
	}

	var childAttachments arrayFlags
	var features arrayFlags

	// Define all flags
	flag.StringVar(&input.Type, "type", "", "Creative type: single_image | single_video | carousel | flexible")
	flag.StringVar(&input.AccountID, "account", "", "Facebook Ad Account ID")
	flag.StringVar(&input.PageID, "page", "", "Facebook Page ID")
	flag.StringVar(&input.Name, "name", "", "Creative name")
	flag.StringVar(&input.ImageHash, "image", "", "Image hash")
	flag.StringVar(&input.Link, "link", "", "Landing page URL")
	flag.StringVar(&input.Picture, "picture", "", "Picture URL")
	flag.StringVar(&input.Message, "message", "", "Creative message")
	flag.StringVar(&input.VideoID, "video", "", "Video ID")
	flag.StringVar(&input.Thumbnail, "thumbnail", "", "Thumbnail URL")
	flag.StringVar(&input.AdvantageOptimizeCreative, "advantage_optimize", "off", "Enable advantage optimize creative: on | off")

	flag.Var(
		&childAttachments,
		"child_attachment",
		"Carousel child attachment: name=...,image_hash=...,desc=...,link=...,video_id=...",
	)

	flag.Var(
		&features,
		"feature",
		"Enable feature: feature_name=on/off (e.g., -feature advantage_plus_creative=on)",
	)

	flag.Parse()

	// Parse child attachments
	parsedChildren, err := h.parseChildAttachments(childAttachments)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse child attachments: %v", err))
	}
	input.ChildAttachmentsInput = parsedChildren

	// Parse features
	for _, feature := range features {
		parts := strings.SplitN(feature, "=", 2)
		if len(parts) == 2 {
			input.Features[parts[0]] = parts[1]
		}
	}

	return input
}

// parseChildAttachments parses carousel child attachment strings
func (h *CreativeHandler) parseChildAttachments(inputs []string) ([]model.ChildAttachmentsData, error) {
	var result []model.ChildAttachmentsData

	for _, raw := range inputs {
		item := model.ChildAttachmentsData{}

		parts := strings.Split(raw, ",")
		for _, part := range parts {
			kv := strings.SplitN(part, "=", 2)
			if len(kv) != 2 {
				return nil, fmt.Errorf("invalid child_attachment format: %s (expected key=value pairs)", raw)
			}

			key := kv[0]
			val := kv[1]

			switch key {
			case "name":
				item.Name = val
			case "image_hash":
				item.ImageHash = val
			case "link":
				item.Link = val
			case "desc":
				item.Description = val
			case "video_id":
				item.VideoID = val
			default:
				return nil, fmt.Errorf("unknown child_attachment key: %s", key)
			}
		}

		result = append(result, item)
	}

	return result, nil
}
