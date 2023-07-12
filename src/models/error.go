package model

import (
	"github.com/gofiber/fiber/v2"
)

// ErrorResponse represents the structure for error responses
type ErrorResponse struct {
	Message string `json:"message"`
}

// Error represents a custom error object that includes an HTTP status code
type Error struct {
	Status  int
	Message string
}

// NewError creates a new Error object with the specified status code and message
func NewError(status int, message string) *Error {
	return &Error{
		Status:  status,
		Message: message,
	}
}

// Error returns the error message
func (e *Error) Error() string {
	return e.Message
}

// ToResponse converts the error into a JSON response
func (e *Error) ToResponse(c *fiber.Ctx) error {
	response := ErrorResponse{
		Message: e.Message,
	}
	return c.Status(e.Status).JSON(response)
}
