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

func main() {
	_ = godotenv.Load()

	input := parseFlags()

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
func parseFlags() facebook.CreativeInput {
	input := facebook.CreativeInput{}

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
	flag.Var(
		&childAttachments,
		"child_attachment",
		"Carousel child attachment: name=...,image_hash=...,desc=...,link=...,video_id=...",
	)
	flag.Parse()
	parsedChildren, err := parseChildAttachments(childAttachments)
	if err != nil {
		log.Fatal(err)
	}
	input.ChildAttachmentsInput = parsedChildren
	return input
}

func dispatchCreate(client *facebook.Client, in facebook.CreativeInput) error {
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
func createSingleImage(client *facebook.Client, in facebook.CreativeInput) error {
	facebook.ValidateInput(in)

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

func createSingleVideo(client *facebook.Client, in facebook.CreativeInput) error {
	facebook.ValidateInput(in)

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

func createCreativeCarousel(client *facebook.Client, in facebook.CreativeInput) error {
	facebook.ValidateInput(in)

	creative := dto.CreateCreativeCarouselRequest{
		ObjectStorySpecCarousel: model.CarouselDetail{
			PageID: in.PageID,
			LinkData: model.LinkDataCarousel{
				Link:             in.Link,
				ChildAttachments: in.ChildAttachmentsInput,
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

func parseChildAttachments(inputs []string) ([]model.ChildAttachmentsData, error) {
	var result []model.ChildAttachmentsData

	for _, raw := range inputs {
		item := model.ChildAttachmentsData{}

		parts := strings.Split(raw, ",")
		for _, part := range parts {
			kv := strings.SplitN(part, "=", 2)
			if len(kv) != 2 {
				return nil, fmt.Errorf("invalid child_attachment format: %s", raw)
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
			}
		}

		result = append(result, item)
	}

	return result, nil
}
