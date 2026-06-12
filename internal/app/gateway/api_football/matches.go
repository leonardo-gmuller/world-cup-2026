package api_football

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	match_usecase "github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/usecase/match"
)

func (c *Client) FetchTodayMatches(
	ctx context.Context,
	date time.Time,
) ([]match_usecase.ExternalMatchOutput, error) {
	url := fmt.Sprintf(
		"%s/fixtures?date=%s",
		c.baseURL,
		date.Format("2006-01-02"),
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("x-apisports-key", c.apiKey)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("api-football returned status %d: %s", resp.StatusCode, string(body))
	}

	var payload FixturesResponse
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return nil, err
	}

	filtered := make([]FixtureDTO, 0)

	for _, fixture := range payload.Response {
		if fixture.League.ID != c.league {
			continue
		}

		if fixture.League.Season != c.season {
			continue
		}

		filtered = append(filtered, fixture)
	}

	return mapFixtures(filtered)
}
