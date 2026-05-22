package entity

type UserRanking struct {
	GroupID      int64
	UserID       int64
	TotalPlayers int64
	TotalPoints  int64
	Position     int
}
