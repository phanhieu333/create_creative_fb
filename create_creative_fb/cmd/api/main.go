package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"creative_fb/internal/dto"
	"creative_fb/internal/facebook"
	"creative_fb/internal/model"

	"github.com/joho/godotenv"
)

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
	ChildAttachmentsInput []string
}

func main() {
	_ = godotenv.Load()

	input := parseFlags()
	validateInput(input)

	accessToken := os.Getenv("FACEBOOK_ACCESS_TOKEN")
	client := facebook.NewClient(accessToken)

	if err := dispatchCreate(client, input); err != nil {
		log.Fatalf("Failed to create creative: %v", err)
	}
}

type arrayFlags []string

func (i *arrayFlags) String() string {
	return strings.Join(*i, ",")
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}
func parseFlags() CreativeInput {
	input := CreativeInput{}

	var childAttachments arrayFlags

	flag.StringVar(&input.Type, "type", "", "single_image | single_video")
	flag.StringVar(&input.AccountID, "account", "", "Ad Account ID")
	flag.StringVar(&input.PageID, "page", "", "Facebook Page ID")
	flag.StringVar(&input.Name, "name", "", "Creative name")
	flag.StringVar(&input.ImageHash, "image", "", "Image hash")
	flag.StringVar(&input.Link, "link", "", "Landing page URL")
	flag.StringVar(&input.Picture, "picture", "", "Picture URL")
	flag.StringVar(&input.Message, "message", "", "Creative message")
	flag.StringVar(&input.VideoID, "video", "", "Video ID")
	flag.StringVar(&input.Thumbnail, "thumbnail", "", "Thumbnail URL")
	flag.StringVar(&input.Picture, "picture", "", "Picture URL")
	flag.Var(
		&childAttachments,
		"child_attachment",
		"Carousel child attachment: name=...,image_hash=...,desc=...,link=...,video_id=...",
	)
	flag.Parse()
	input.ChildAttachmentsInput = childAttachments
	return input
}
func validateInput(in CreativeInput) {
	if in.Type == "" {
		log.Fatal("Missing -type")
	}

	if in.AccountID == "" || in.PageID == "" {
		log.Fatal("Missing -account or -page")
	}

	if in.Type == "single_video" {
		if in.VideoID == "" {
			log.Fatal("Missing -video Id")
		}
		if in.Thumbnail == "" {
			log.Fatal("Missing -thumbnail")
		}
	}

	if in.Type == "carousel" {
		if len(in.ChildAttachmentsInput) < 2 {
			log.Fatal("At least two -child_attachment are required for carousel")
		}
		if in.ImageHash == "" && in.Picture == "" {
			log.Fatal("Missing -image or -picture for carousel")
		}
	}

}
func dispatchCreate(client *facebook.Client, in CreativeInput) error {
	switch in.Type {
	case "single_image":
		return createSingleImage(client, in)
	case "single_video":
		return createSingleVideo(client, in)
	case "carousel":
		return createCreativeCarousel(client, in)
	default:
		return fmt.Errorf("unsupported creative type: %s", in.Type)
	}
}
func createSingleImage(client *facebook.Client, in CreativeInput) error {
	creative := dto.CreateCreativeSingleImageRequest{
		ObjectStorySpecImage: model.ImageDetail{
			PageID: in.PageID,
			LinkData: model.LinkData{
				Link: in.Link,
			},
		},
	}

	res, err := client.CreateCreativeSingleImage(in.AccountID, creative)
	if err != nil {
		return err
	}

	fmt.Printf("✓ Creative created successfully!\n")
	fmt.Printf("  Creative ID: %s\n", res.ID)
	return nil
}

func createSingleVideo(client *facebook.Client, in CreativeInput) error {
	creative := dto.CreateCreativeSingleVideoRequest{
		ObjectStorySpecVideo: model.VideoDetail{
			PageID: in.PageID,
			VideoData: model.VideoData{
				VideoID:  in.VideoID,
				ImageURL: in.Thumbnail,
			},
		},
	}

	res, err := client.CreateCreativeSingleVideo(in.AccountID, creative)
	if err != nil {
		return err
	}

	fmt.Printf("✓ Creative created successfully!\n")
	fmt.Printf("  Creative ID: %s\n", res.ID)
	return nil
}

func createCreativeCarousel(client *facebook.Client, in CreativeInput) error {
	childAttachments, err := parseChildAttachmentsInput(in.ChildAttachmentsInput)
	creative := dto.CreateCreativeCarouselRequest{
		ObjectStorySpecCarousel: model.CarouselDetail{
			PageID: in.PageID,
			LinkData: model.LinkDataCarousel{
				Link:             in.Link,
				ChildAttachments: childAttachments,
			},
		},
	}

	fmt.Printf("creative %+v \n", creative)

	res, err := client.CreateCreativeCarousel(in.AccountID, creative)
	if err != nil {
		return err
	}
	fmt.Printf("✓ Creative created successfully!\n")
	fmt.Printf("  Creative ID: %s\n", res.ID)
	return nil
}
func parseChildAttachmentsInput(
	input []string,
) ([]model.ChildAttachmentsData, error) {

	var attachments []model.ChildAttachmentsData

	for _, raw := range input {
		fields := map[string]string{}

		for _, part := range strings.Split(raw, ",") {
			kv := strings.SplitN(part, "=", 2)
			if len(kv) != 2 {
				return nil, fmt.Errorf("invalid child_attachment format: %s", raw)
			}
			fields[kv[0]] = kv[1]
		}

		attachments = append(attachments, model.ChildAttachmentsData{
			Name:        fields["name"],
			Description: fields["desc"],
			ImageHash:   fields["image_hash"],
			Link:        fields["link"],
			VideoID:     fields["video_id"],
		})
	}

	return attachments, nil
}
