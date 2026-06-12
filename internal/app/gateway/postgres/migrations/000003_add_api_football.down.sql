ALTER TABLE matches
DROP COLUMN api_football_id,
DROP COLUMN result_source,
DROP COLUMN last_live_sync_at;

DROP INDEX IF EXISTS idx_matches_api_football_id;
