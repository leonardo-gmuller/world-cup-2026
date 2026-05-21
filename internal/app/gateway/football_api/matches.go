package football_api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	match_usecase "github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/usecase/match"
)

func (c *Client) FetchWorldCupMatches(
	ctx context.Context,
) ([]match_usecase.ExternalMatchOutput, error) {

	url := fmt.Sprintf(
		"%s/competitions/WC/matches",
		c.baseURL,
	)

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Auth-Token", c.apiKey)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("football api returned status %d, error: %s", resp.StatusCode, string(body))
	}

	var payload MatchesResponse
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return nil, err
	}

	return mapMatches(payload.Matches)
}
