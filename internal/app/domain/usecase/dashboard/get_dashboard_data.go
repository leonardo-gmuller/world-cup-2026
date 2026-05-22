package dashboard_usecase

import (
	"context"
)

type HomeDashboardOutput struct {
	NextMatch NextMatchOutput `json:"next_match"`

	PredictionsCount int64 `json:"predictions_count"`
	GroupsCount      int64 `json:"groups_count"`
	MatchesCount     int64 `json:"matches_count"`
	UserPoints       int64 `json:"user_points"`
	TotalPlayers     int64 `json:"total_players"`
	TotalPoints      int64 `json:"total_points"`
	BestRanking      int64 `json:"best_ranking"`
}

type NextMatchOutput struct {
	ID           string  `json:"id"`
	HomeTeamName *string `json:"home_team_name"`
	AwayTeamName *string `json:"away_team_name"`
	Stage        string  `json:"stage"`
	GroupName    *string `json:"group_name"`
	StartAt      string  `json:"start_at"`
	HomeTeamFlag *string `json:"home_team_flag"`
	AwayTeamFlag *string `json:"away_team_flag"`
}

func (u *DashboardUseCase) GetDashboardData(
	ctx context.Context,
	userID int64,
) (*HomeDashboardOutput, error) {
	nextMatch, err := u.matchRepo.GetNextMatch(ctx)
	if err != nil {
		return nil, err
	}

	predictionsCount, err := u.predictionRepo.CountPredictionsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	groupsCount, err := u.groupRepo.CountGroupsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	matchesCount, err := u.matchRepo.CountMatches(ctx)
	if err != nil {
		return nil, err
	}

	bestRanking, err := u.predictionRepo.GetBestRankingByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &HomeDashboardOutput{
		NextMatch: NextMatchOutput{
			ID:           nextMatch.UUID.String(),
			HomeTeamName: nextMatch.HomeTeamName,
			AwayTeamName: nextMatch.AwayTeamName,
			Stage:        nextMatch.Stage,
			GroupName:    nextMatch.GroupName,
			StartAt:      nextMatch.StartsAt.Format("2006-01-02T15:04:05Z07:00"),
			HomeTeamFlag: nextMatch.HomeTeamFlagURL,
			AwayTeamFlag: nextMatch.AwayTeamFlagURL,
		},
		PredictionsCount: predictionsCount,
		GroupsCount:      groupsCount,
		MatchesCount:     matchesCount,
		UserPoints:       bestRanking.TotalPoints,
		TotalPlayers:     bestRanking.TotalPlayers,
		TotalPoints:      bestRanking.TotalPoints,
		BestRanking:      int64(bestRanking.Position),
	}, nil
}
