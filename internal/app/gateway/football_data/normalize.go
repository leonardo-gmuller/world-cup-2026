package football_data

import (
	"strings"

	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/constants"
)

func normalizeStage(stage string) string {
	switch strings.ToUpper(stage) {

	case "GROUP_STAGE":
		return constants.StageGroupStage

	case "LAST_64":
		return constants.StageGroupStage

	case "LAST_32":
		return constants.StageRoundOf32

	case "LAST_16":
		return constants.StageRoundOf16

	case "QUARTER_FINALS":
		return constants.StageQuarterFinal

	case "SEMI_FINALS":
		return constants.StageSemiFinal

	case "THIRD_PLACE":
		return constants.StageThirdPlace

	case "FINAL":
		return constants.StageFinal

	default:
		return constants.StageGroupStage
	}
}

func normalizeStatus(status string) string {
	switch status {
	case "TIMED", "SCHEDULED", "POSTPONED":
		return "scheduled"

	case "IN_PLAY", "2H", "HT", "LIVE":
		return "live"

	case "FINISHED":
		return "finished"

	default:
		return "scheduled"
	}
}
