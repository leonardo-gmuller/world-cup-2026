package football_api

import (
	"net/http"
	"time"

	"github.com/leonardo-gmuller/world-cup-2026/internal/app/config"
)

type Client struct {
	baseURL string
	apiKey  string
	client  *http.Client
}

func New(cfg config.Config) *Client {
	return &Client{
		baseURL: cfg.FootballAPI.BaseURL,
		apiKey:  cfg.FootballAPI.APIKey,
		client: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}
