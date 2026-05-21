package match_repository

import (
	"time"

	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/postgres/sqlc"
)

func mapTeam(row sqlc.Team) *entity.Team {
	return &entity.Team{
		ID:   row.ID,
		UUID: row.Uuid,
		ExternalID: func() *string {
			if row.ExternalID.Valid {
				return &row.ExternalID.String
			}
			return nil
		}(),
		Name: row.Name,
		ShortName: func() *string {
			if row.ShortName.Valid {
				return &row.ShortName.String
			}
			return nil
		}(),
		Code: func() *string {
			if row.Code.Valid {
				return &row.Code.String
			}
			return nil
		}(),
		FlagURL: func() *string {
			if row.FlagUrl.Valid {
				return &row.FlagUrl.String
			}
			return nil
		}(),
		CreatedAt: row.CreatedAt.Time,
		UpdatedAt: row.UpdatedAt.Time,
		DeletedAt: func() *time.Time {
			if row.DeletedAt.Valid {
				return &row.DeletedAt.Time
			}
			return nil
		}(),
	}
}

func mapMatch(row sqlc.Match) *entity.Match {
	return &entity.Match{
		ID:   row.ID,
		UUID: row.Uuid,
		ExternalID: func() *string {
			if row.ExternalID.Valid {
				return &row.ExternalID.String
			}
			return nil
		}(),
		Stage: row.Stage,
		GroupName: func() *string {
			if row.GroupName.Valid {
				return &row.GroupName.String
			}
			return nil
		}(),
		HomeTeamID: func() *int64 {
			if row.HomeTeamID.Valid {
				return &row.HomeTeamID.Int64
			}
			return nil
		}(),
		AwayTeamID: func() *int64 {
			if row.AwayTeamID.Valid {
				return &row.AwayTeamID.Int64
			}
			return nil
		}(),
		HomeTeamName: func() *string {
			if row.HomeTeamName.Valid {
				return &row.HomeTeamName.String
			}
			return nil
		}(),
		AwayTeamName: func() *string {
			if row.AwayTeamName.Valid {
				return &row.AwayTeamName.String
			}
			return nil
		}(),
		StartsAt: row.StartsAt.Time,
		HomeScore: func() *int {
			if row.HomeScore.Valid {
				score := int(row.HomeScore.Int32)
				return &score
			}
			return nil
		}(),
		AwayScore: func() *int {
			if row.AwayScore.Valid {
				score := int(row.AwayScore.Int32)
				return &score
			}
			return nil
		}(),
		Status: row.Status,
		WinnerTeamID: func() *int64 {
			if row.WinnerTeamID.Valid {
				return &row.WinnerTeamID.Int64
			}
			return nil
		}(),
		ImportedAt: func() *time.Time {
			if row.ImportedAt.Valid {
				return &row.ImportedAt.Time
			}
			return nil
		}(),
		CreatedAt: row.CreatedAt.Time,
		UpdatedAt: row.UpdatedAt.Time,
		DeletedAt: func() *time.Time {
			if row.DeletedAt.Valid {
				return &row.DeletedAt.Time
			}
			return nil
		}(),
	}
}

func mapListMatchesRow(row sqlc.ListMatchesRow) *entity.Match {
	match := mapMatch(sqlc.Match{
		ID:           row.ID,
		Uuid:         row.Uuid,
		ExternalID:   row.ExternalID,
		Stage:        row.Stage,
		GroupName:    row.GroupName,
		HomeTeamID:   row.HomeTeamID,
		AwayTeamID:   row.AwayTeamID,
		HomeTeamName: row.HomeTeamName,
		AwayTeamName: row.AwayTeamName,
		StartsAt:     row.StartsAt,
		HomeScore:    row.HomeScore,
		AwayScore:    row.AwayScore,
		Status:       row.Status,
		WinnerTeamID: row.WinnerTeamID,
		ImportedAt:   row.ImportedAt,
		CreatedAt:    row.CreatedAt,
		UpdatedAt:    row.UpdatedAt,
		DeletedAt:    row.DeletedAt,
	})

	if row.HomeTeamFlagUrl.Valid {
		match.HomeTeamFlagURL = &row.HomeTeamFlagUrl.String
	}

	if row.AwayTeamFlagUrl.Valid {
		match.AwayTeamFlagURL = &row.AwayTeamFlagUrl.String
	}

	return match
}
