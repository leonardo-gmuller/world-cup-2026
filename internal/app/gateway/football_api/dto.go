package football_api

type MatchesResponse struct {
	Filters     FiltersDTO     `json:"filters"`
	ResultSet   ResultSetDTO   `json:"resultSet"`
	Competition CompetitionDTO `json:"competition"`
	Matches     []MatchDTO     `json:"matches"`
}

type FiltersDTO struct {
	Season string `json:"season"`
}

type ResultSetDTO struct {
	Count  int    `json:"count"`
	First  string `json:"first"`
	Last   string `json:"last"`
	Played int    `json:"played"`
}

type CompetitionDTO struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	Type   string `json:"type"`
	Emblem string `json:"emblem"`
}

type MatchDTO struct {
	ID       int64    `json:"id"`
	UTCDate  string   `json:"utcDate"`
	Status   string   `json:"status"`
	Matchday *int     `json:"matchday"`
	Stage    string   `json:"stage"`
	Group    *string  `json:"group"`
	HomeTeam TeamDTO  `json:"homeTeam"`
	AwayTeam TeamDTO  `json:"awayTeam"`
	Score    ScoreDTO `json:"score"`
}

type TeamDTO struct {
	ID        *int64  `json:"id"`
	Name      *string `json:"name"`
	ShortName *string `json:"shortName"`
	TLA       *string `json:"tla"`
	Crest     *string `json:"crest"`
}

type ScoreDTO struct {
	Winner   *string      `json:"winner"`
	Duration string       `json:"duration"`
	FullTime ScoreTimeDTO `json:"fullTime"`
	HalfTime ScoreTimeDTO `json:"halfTime"`
}

type ScoreTimeDTO struct {
	Home *int `json:"home"`
	Away *int `json:"away"`
}
