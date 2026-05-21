package util

import (
	"database/sql"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func StrPtr(n sql.NullString) *string {
	if n.Valid {
		s := n.String
		return &s
	}
	return nil
}

func TimePtr(n sql.NullTime) *time.Time {
	if n.Valid {
		t := n.Time
		return &t
	}
	return nil
}

func StrOrEmpty(n sql.NullString) string {
	if n.Valid {
		return n.String
	}
	return ""
}

func DerefOrEmpty(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func TextPtr(s *string) pgtype.Text {
	if s == nil {
		return pgtype.Text{Valid: false}
	}
	return pgtype.Text{String: *s, Valid: true}
}

func Int8Ptr(n *int64) pgtype.Int8 {
	if n == nil {
		return pgtype.Int8{Valid: false}
	}
	return pgtype.Int8{Int64: int64(*n), Valid: true}
}

func Int4Ptr(n *int) pgtype.Int4 {
	if n == nil {
		return pgtype.Int4{Valid: false}
	}
	return pgtype.Int4{Int32: int32(*n), Valid: true}
}

func PtrTime(t *time.Time) pgtype.Timestamp {
	if t == nil {
		return pgtype.Timestamp{Valid: false}
	}
	return pgtype.Timestamp{Time: *t, Valid: true}
}
