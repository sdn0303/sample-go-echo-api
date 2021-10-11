package utils

import "log"

// ErrorHandlerIFace Interface for error handler
type ErrorHandlerIFace interface {
	Fatalf(err error, msg string)
}

type errorHandler struct{}

// NewErrorHandler constructs an error handler
func NewErrorHandler() ErrorHandlerIFace {
	return &errorHandler{}
}

// Fatalf print error message
func (eh *errorHandler) Fatalf(err error, msg string) {
	if err != nil {
		log.Fatalf(msg+": %v\n", err)
	}
}
