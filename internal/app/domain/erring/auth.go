package erring

var (
	ErrInvalidCredentials = NewAppError("auth:invalid-credentials", "invalid credentials")
	ErrEmailAlreadyUsed   = NewAppError("auth:email-already-used", "email already used")
)
