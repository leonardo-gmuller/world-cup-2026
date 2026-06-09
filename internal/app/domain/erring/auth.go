package erring

var (
	ErrInvalidCredentials = NewAppError("auth:invalid-credentials", "invalid credentials")
	ErrEmailAlreadyUsed   = NewAppError("auth:email-already-used", "email already used")
	ErrInvalidResetToken  = NewAppError("auth:invalid-reset-token", "invalid reset token")
	ErrExpiredResetToken  = NewAppError("auth:expired-reset-token", "expired reset token")
	ErrUsedResetToken     = NewAppError("auth:used-reset-token", "used reset token")
	ErrPasswordsDontMatch = NewAppError("auth:passwords-dont-match", "passwords don't match")
)
