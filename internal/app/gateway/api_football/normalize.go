package api_football

import (
	"strings"

	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/constants"
)

func normalizeStatus(status string) string {
	switch strings.ToUpper(status) {
	case "TBD", "NS":
		return "scheduled"

	case "1H", "2H", "HT", "ET", "BT", "P", "SUSP", "INT", "LIVE":
		return "live"

	case "FT", "AET", "PEN":
		return "finished"

	case "PST", "CANC", "ABD", "AWD", "WO":
		return "cancelled"

	default:
		return "scheduled"
	}
}

func normalizeStage(round string) string {
	value := strings.ToLower(round)

	switch {
	case strings.Contains(value, "group"):
		return constants.StageGroupStage
	case strings.Contains(value, "round of 32"):
		return constants.StageRoundOf32
	case strings.Contains(value, "round of 16"):
		return constants.StageRoundOf16
	case strings.Contains(value, "quarter"):
		return constants.StageQuarterFinal
	case strings.Contains(value, "semi"):
		return constants.StageSemiFinal
	case strings.Contains(value, "third"):
		return constants.StageThirdPlace
	case strings.Contains(value, "final"):
		return constants.StageFinal
	default:
		return constants.StageGroupStage
	}
}
