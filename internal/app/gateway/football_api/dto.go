package football_api

type MatchesResponse struct {
	Response []MatchDTO `json:"response"`
}

type MatchDTO struct {
	Fixture FixtureDTO `json:"fixture"`
	League  LeagueDTO  `json:"league"`
	Teams   TeamsDTO   `json:"teams"`
	Goals   GoalsDTO   `json:"goals"`
}

type FixtureDTO struct {
	ID     int64     `json:"id"`
	Date   string    `json:"date"`
	Status StatusDTO `json:"status"`
}

type StatusDTO struct {
	Short string `json:"short"`
}

type LeagueDTO struct {
	Round string `json:"round"`
}

type TeamsDTO struct {
	Home TeamDTO `json:"home"`
	Away TeamDTO `json:"away"`
}

type TeamDTO struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
	Logo string `json:"logo"`
}

type GoalsDTO struct {
	Home *int `json:"home"`
	Away *int `json:"away"`
}
