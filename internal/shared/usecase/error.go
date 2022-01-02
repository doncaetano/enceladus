package usecase

type UseCaseErrorType int

const (
	SERVER_ERROR UseCaseErrorType = iota
	BAD_REQUEST
	UNAUTHORIZED
)

type UseCaseError struct {
	Type    UseCaseErrorType
	Message string
}

var availableErrors = []string{"SERVER_ERROR", "BAD_REQUEST", "UNAUTHORIZED"}

func (err UseCaseErrorType) String() string {
	return availableErrors[err]
}

func ServerError() *UseCaseError {
	return &UseCaseError{
		Type:    SERVER_ERROR,
		Message: "the server encountered an unexpected condition that prevented it from fulfilling the request",
	}
}

func BadRequestError(message string) *UseCaseError {
	return &UseCaseError{
		Type:    BAD_REQUEST,
		Message: message,
	}
}

func UnauthorizedError(message string) *UseCaseError {
	return &UseCaseError{
		Type:    UNAUTHORIZED,
		Message: message,
	}
}
