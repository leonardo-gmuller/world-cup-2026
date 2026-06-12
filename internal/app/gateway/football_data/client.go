package football_data

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
		baseURL: cfg.APIFootball.BaseURL,
		apiKey:  cfg.APIFootball.APIKey,
		client: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}
