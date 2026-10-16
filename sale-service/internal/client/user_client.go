package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type UserResponse struct {
	ID       string `json:"id"`
	FullName string `json:"full_name"`
	UserName string `json:"user_name"`
}

type GetUserResponse struct {
	Message string       `json:"message"`
	User    UserResponse `json:"user"`
}

type UserClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewUserClient(baseURL string) *UserClient {
	return &UserClient{
		BaseURL: strings.TrimRight(baseURL, "/"),
		HTTPClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (c *UserClient) GetUserByID(
	ctx context.Context,
	id string,
) (*UserResponse, error) {
	endpoint := c.BaseURL + "/api/v1/users/" + url.PathEscape(id)

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		endpoint,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("create user request: %w", err)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("call user service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"user service returned status %d",
			resp.StatusCode,
		)
	}

	var result GetUserResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("decode user response: %w", err)
	}

	return &result.User, nil
}
