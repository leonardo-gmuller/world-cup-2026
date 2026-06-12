package types

type Job string

const (
	ImportMatches             Job = "import-matches"
	CalculatePredictionPoints Job = "calculate-prediction-points"
	SyncLiveResults           Job = "sync-live-results"
)
