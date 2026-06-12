package api_football

import (
	"net/http"
	"time"

	"github.com/leonardo-gmuller/world-cup-2026/internal/app/config"
)

type Client struct {
	baseURL string
	apiKey  string
	league  int64
	season  int
	client  *http.Client
}

func New(cfg config.Config) *Client {
	return &Client{
		baseURL: cfg.APIFootball.BaseURL,
		apiKey:  cfg.APIFootball.APIKey,
		league:  int64(cfg.APIFootball.League),
		season:  cfg.APIFootball.Season,
		client: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}
