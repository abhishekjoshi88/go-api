package domain

// ErrorDetails is a struct used for storing response of error details
type ErrorDetails struct {
	Code        string `json:"errorCode"`
	Description string `json:"errorDescription"`
}

func (er ErrorDetails) Error() string {
	return er.Code
}

//UnexpectedError details
var UnexpectedError = ErrorDetails{Code: "unexpectedError", Description: "An unexpected error occurred. Please try again later."}

//ErrDB detais
var ErrDB = ErrorDetails{Code: "noContentFound", Description: "No content found"}

//ErrInvalidID Corner case error
var ErrInvalidID = ErrorDetails{Code: "invalidID", Description: "Cannot use id=0 in param"}
