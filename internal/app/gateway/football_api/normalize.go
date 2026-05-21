package football_api

import "strings"

func normalizeStage(round string) string {
	round = strings.ToLower(round)

	switch {
	case strings.Contains(round, "group"):
		return "group_stage"

	case strings.Contains(round, "round of 16"):
		return "round_of_16"

	case strings.Contains(round, "quarter"):
		return "quarter_final"

	case strings.Contains(round, "semi"):
		return "semi_final"

	case strings.Contains(round, "final"):
		return "final"

	default:
		return "group_stage"
	}
}

func normalizeStatus(status string) string {
	switch status {
	case "NS":
		return "scheduled"

	case "1H", "2H", "HT", "LIVE":
		return "live"

	case "FT":
		return "finished"

	default:
		return "scheduled"
	}
}
