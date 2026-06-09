package auth_repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/postgres/sqlc"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/pkg/uow"
)

type PasswordResetTokenRepository struct {
	*sqlc.Queries
}

func NewPasswordResetTokenRepository(db uow.DBTX) *PasswordResetTokenRepository {
	return &PasswordResetTokenRepository{
		Queries: sqlc.New(db),
	}
}

func (r *PasswordResetTokenRepository) CreatePasswordResetToken(
	ctx context.Context,
	userID int64,
	token string,
	expiresAt time.Time,
) (*entity.PasswordResetToken, error) {
	row, err := r.Queries.CreatePasswordResetToken(ctx, sqlc.CreatePasswordResetTokenParams{
		UserID:    userID,
		Token:     token,
		ExpiresAt: pgtype.Timestamptz{Time: expiresAt, Valid: true},
	})
	if err != nil {
		return nil, err
	}

	return &entity.PasswordResetToken{
		ID:        row.ID,
		UserID:    row.UserID,
		Token:     row.Token,
		Used:      row.Used,
		ExpiresAt: row.ExpiresAt.Time,
		CreatedAt: row.CreatedAt.Time,
		UsedAt: func() *time.Time {
			if row.UsedAt.Valid {
				return &row.UsedAt.Time
			} else {
				return nil
			}
		}(),
	}, nil
}

func (r *PasswordResetTokenRepository) GetPasswordResetTokenByToken(
	ctx context.Context,
	token string,
) (*entity.PasswordResetToken, error) {
	row, err := r.Queries.GetPasswordResetTokenByToken(ctx, token)
	if err != nil {
		return nil, err
	}

	return &entity.PasswordResetToken{
		ID:        row.ID,
		UserID:    row.UserID,
		Token:     row.Token,
		Used:      row.Used,
		ExpiresAt: row.ExpiresAt.Time,
		CreatedAt: row.CreatedAt.Time,
		UsedAt: func() *time.Time {
			if row.UsedAt.Valid {
				return &row.UsedAt.Time
			} else {
				return nil
			}
		}(),
	}, nil
}

func (r *PasswordResetTokenRepository) InvalidatePasswordResetTokensByUserID(
	ctx context.Context,
	userID int64,
) error {
	return r.Queries.InvalidatePasswordResetTokensByUserID(ctx, userID)
}

func (r *PasswordResetTokenRepository) UsePasswordResetToken(
	ctx context.Context,
	id int64,
) error {
	return r.Queries.UsePasswordResetToken(ctx, id)
}

func (r *PasswordResetTokenRepository) DeleteExpiredPasswordResetTokens(
	ctx context.Context,
) error {
	return r.Queries.DeleteExpiredPasswordResetTokens(ctx)
}
