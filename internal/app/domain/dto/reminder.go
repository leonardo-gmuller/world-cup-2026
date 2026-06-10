package dto

import "time"

type PredictionReminderOutput struct {
	MatchID         int64
	GroupID         int64
	GroupName       string
	HomeTeamName    string
	AwayTeamName    string
	HomeTeamFlagURL string
	AwayTeamFlagURL string
	StartsAt        time.Time
}
