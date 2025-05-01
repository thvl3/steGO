package stego

import "fmt"

// Error types for steganography operations
type ErrorType int

const (
	ErrInvalidInput ErrorType = iota
	ErrFileOperation
	ErrImageProcessing
	ErrMessageTooLarge
	ErrInvalidImageFormat
	ErrInvalidMessage
	ErrEncoding
	ErrDecoding
	ErrSecurity
)

// StegoError represents a steganography-specific error
type StegoError struct {
	Type    ErrorType
	Message string
	Err     error
}

// Error implements the error interface
func (e *StegoError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// NewError creates a new StegoError
func NewError(t ErrorType, msg string, err error) *StegoError {
	return &StegoError{
		Type:    t,
		Message: msg,
		Err:     err,
	}
}

// Error messages
const (
	ErrMsgInvalidInput       = "invalid input parameters"
	ErrMsgFileOperation      = "file operation failed"
	ErrMsgImageProcessing    = "image processing failed"
	ErrMsgMessageTooLarge    = "message is too large for the image"
	ErrMsgInvalidImageFormat = "invalid image format"
	ErrMsgInvalidMessage     = "invalid message format"
	ErrMsgEncoding           = "encoding failed"
	ErrMsgDecoding           = "decoding failed"
	ErrMsgSecurity           = "security operation failed"
)

// Is checks if the error is of a specific type
func Is(err error, t ErrorType) bool {
	if e, ok := err.(*StegoError); ok {
		return e.Type == t
	}
	return false
}
