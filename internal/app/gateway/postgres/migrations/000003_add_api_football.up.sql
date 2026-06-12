ALTER TABLE matches
ADD COLUMN api_football_id BIGINT,
ADD COLUMN result_source VARCHAR(30) NOT NULL DEFAULT 'football_data',
ADD COLUMN last_live_sync_at TIMESTAMPTZ;

CREATE UNIQUE INDEX idx_matches_api_football_id
ON matches(api_football_id)
WHERE api_football_id IS NOT NULL;