package facebook

import (
	"bytes"
	"creative_fb/internal/dto"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	BaseURL    = "https://graph.facebook.com/v24.0"
	APIVersion = "v24.0"
)

// Client handles Facebook API requests
type Client struct {
	accessToken string
	httpClient  *http.Client
}

// NewClient creates a new Facebook API client
func NewClient(accessToken string) *Client {
	return &Client{
		accessToken: accessToken,
		httpClient:  &http.Client{},
	}
}

func (c *Client) doJSONRequest(
	method string,
	endpoint string,
	payload []byte,
) ([]byte, int, error) {

	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(payload))
	if err != nil {
		return nil, 0, fmt.Errorf("failed to create request: %w", err)
	}

	q := req.URL.Query()
	q.Add("access_token", c.accessToken)
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf("failed to read response body: %w", err)
	}

	fmt.Printf("Response Body: %s\n", body)

	return body, resp.StatusCode, nil
}

// CreateCreative creates a new creative ad
func (c *Client) CreateCreativeSingleImage(accountID string, creative dto.CreateCreativeSingleImageRequest) (*dto.CreateCreativeResponse, error) {
	endpoint := fmt.Sprintf("%s/act_%s/adcreatives", BaseURL, accountID)

	payload, err := json.Marshal(creative)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	body, status, err := c.doJSONRequest(
		http.MethodPost,
		endpoint,
		payload,
	)
	if err != nil {
		return nil, err
	}

	if status >= 400 {
		var errResp map[string]interface{}
		json.Unmarshal(body, &errResp)
		return nil, fmt.Errorf("API error (status %d): %v", status, errResp)
	}

	var result dto.CreateCreativeResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

// CreateCreativeSingleVideo creates a new creative ad with single video
func (c *Client) CreateCreativeSingleVideo(accountID string, creative dto.CreateCreativeSingleVideoRequest) (*dto.CreateCreativeResponse, error) {
	endpoint := fmt.Sprintf("%s/act_%s/adcreatives", BaseURL, accountID)

	payload, err := json.Marshal(creative)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	body, status, err := c.doJSONRequest(
		http.MethodPost,
		endpoint,
		payload,
	)
	if err != nil {
		return nil, err
	}

	if status >= 400 {
		var errResp map[string]interface{}
		json.Unmarshal(body, &errResp)
		return nil, fmt.Errorf("API error (status %d): %v", status, errResp)
	}

	var result dto.CreateCreativeResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

func (c *Client) CreateCreativeCarousel(accountID string, creative dto.CreateCreativeCarouselRequest) (*dto.CreateCreativeResponse, error) {
	endpoint := fmt.Sprintf("%s/act_%s/adcreatives", BaseURL, accountID)
	payload, err := json.Marshal(creative)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}
	body, status, err := c.doJSONRequest(
		http.MethodPost,
		endpoint,
		payload,
	)
	if err != nil {
		return nil, err
	}

	if status >= 400 {
		var errResp map[string]interface{}
		json.Unmarshal(body, &errResp)
		return nil, fmt.Errorf("API error (status %d): %v", status, errResp)
	}

	var result dto.CreateCreativeResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

func (c *Client) CreateCreativeFlexible(accountID string, creative dto.CreateCreativeFlexibleRequest) (*dto.CreateCreativeResponse, error) {
	endpoint := fmt.Sprintf("%s/act_%s/adcreatives", BaseURL, accountID)
	payload, err := json.Marshal(creative)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}
	body, status, err := c.doJSONRequest(
		http.MethodPost,
		endpoint,
		payload,
	)
	if err != nil {
		return nil, err
	}

	if status >= 400 {
		var errResp map[string]interface{}
		json.Unmarshal(body, &errResp)
		return nil, fmt.Errorf("API error (status %d): %v", status, errResp)
	}

	var result dto.CreateCreativeResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}
