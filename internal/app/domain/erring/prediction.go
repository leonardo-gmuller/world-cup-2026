package erring

var (
	ErrPredictionClosed  = NewAppError("prediction:closed", "prediction is closed for this match")
	ErrMatchWithoutScore = NewAppError("match:without-score", "match does not have score")
	ErrMatchNotFinished  = NewAppError("match:not-finished", "match is not finished")
)
