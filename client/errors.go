package client

import (
	"errors"
	"fmt"
)

// Sentinel errors exported by the client package.
// Use errors.Is to check which class an error belongs to.
var (
	// ErrRateLimited is returned when the server responds with HTTP 402 or 429.
	ErrRateLimited = errors.New("rate limited (402/429)")

	// ErrServerError is returned when the server responds with any 5xx status.
	ErrServerError = errors.New("server error (5xx)")

	// ErrNotFound is returned when the server responds with HTTP 404.
	ErrNotFound = errors.New("not found (404)")

	// ErrUnexpectedStatus is returned for any non-2xx status that does not map
	// to a more specific sentinel above.  It also matches every *APIError via Is.
	ErrUnexpectedStatus = errors.New("unexpected HTTP status")
)

// APIError is the concrete error type returned by all non-2xx responses.
// It always carries the HTTP status code and at least one non-empty detail
// string (either a parsed JSON message or a raw body snippet).
type APIError struct {
	// StatusCode is the HTTP status returned by the server.
	StatusCode int
	// Message is extracted from the server's JSON error body {"message":"..."}.
	// Empty when the body is not JSON or has no "message" field.
	Message string
	// Body is the raw response body, capped at 512 bytes.
	// Populated when Message is empty.
	Body string
}

// Error implements the error interface.
// The returned string always contains the HTTP status code and a non-empty
// detail: the JSON message, the raw body snippet, or "<empty body>".
func (e *APIError) Error() string {
	detail := e.Message
	if detail == "" {
		detail = e.Body
	}
	if detail == "" {
		detail = "<empty body>"
	}
	return fmt.Sprintf("maestro API error: status %d: %s", e.StatusCode, detail)
}

// Is makes errors.Is(err, ErrRateLimited) etc. work for *APIError values.
// Every *APIError also matches ErrUnexpectedStatus.
func (e *APIError) Is(target error) bool {
	switch target {
	case ErrRateLimited:
		return e.StatusCode == 402 || e.StatusCode == 429
	case ErrNotFound:
		return e.StatusCode == 404
	case ErrServerError:
		return e.StatusCode >= 500 && e.StatusCode < 600
	case ErrUnexpectedStatus:
		return true
	}
	return false
}

// newAPIError constructs an *APIError from a status code, optionally parsed
// JSON message, and the raw body bytes.  The body is trimmed to 512 bytes.
func newAPIError(statusCode int, message string, rawBody []byte) *APIError {
	e := &APIError{StatusCode: statusCode, Message: message}
	if message == "" && len(rawBody) > 0 {
		body := string(rawBody)
		if len(body) > 512 {
			body = body[:512] + "..."
		}
		e.Body = body
	}
	return e
}
