package api_football

type FixturesResponse struct {
	Response []FixtureDTO `json:"response"`
}

type FixtureDTO struct {
	Fixture FixtureInfoDTO `json:"fixture"`
	League  LeagueDTO      `json:"league"`
	Teams   TeamsDTO       `json:"teams"`
	Goals   GoalsDTO       `json:"goals"`
}

type FixtureInfoDTO struct {
	ID     int64     `json:"id"`
	Date   string    `json:"date"`
	Status StatusDTO `json:"status"`
}

type StatusDTO struct {
	Short string `json:"short"`
	Long  string `json:"long"`
}

type LeagueDTO struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
	Season  int    `json:"season"`
	Round   string `json:"round"`
}

type TeamsDTO struct {
	Home TeamDTO `json:"home"`
	Away TeamDTO `json:"away"`
}

type TeamDTO struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
}

type GoalsDTO struct {
	Home *int `json:"home"`
	Away *int `json:"away"`
}
