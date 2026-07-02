package client

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// helper: build a test Client pointing at a given server URL.
func newTestClient(serverURL string) *Client {
	return &Client{
		apiKey:     "test-key",
		network:    "preprod",
		version:    "v1",
		HTTPClient: http.DefaultClient,
		BaseUrl:    serverURL,
	}
}

// --- sendRequest path tests ---

// TestSendRequest402_RateLimited verifies that an HTTP 402 with a JSON body
// returns an error that satisfies errors.Is(err, ErrRateLimited) and that
// Error() contains "402".
func TestSendRequest402_RateLimited(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusPaymentRequired) // 402
		body, _ := json.Marshal(errorResponse{Code: 402, Message: "quota exceeded"})
		w.Write(body) //nolint:errcheck
	}))
	defer srv.Close()

	c := newTestClient(srv.URL)
	req, _ := http.NewRequest("GET", srv.URL+"/some/path", nil)
	var responseBody string
	err := c.sendRequest(req, &responseBody)

	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, ErrRateLimited) {
		t.Errorf("expected ErrRateLimited, got: %v", err)
	}
	if !strings.Contains(err.Error(), "402") {
		t.Errorf("Error() should contain '402', got: %q", err.Error())
	}
	if err.Error() == "" {
		t.Error("Error() must not be empty")
	}
}

// TestSendRequest429_RateLimited verifies that HTTP 429 also maps to ErrRateLimited.
func TestSendRequest429_RateLimited(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(429)
		w.Write([]byte(`{"code":429,"message":"too many requests"}`)) //nolint:errcheck
	}))
	defer srv.Close()

	c := newTestClient(srv.URL)
	req, _ := http.NewRequest("GET", srv.URL+"/some/path", nil)
	var body string
	err := c.sendRequest(req, &body)

	if !errors.Is(err, ErrRateLimited) {
		t.Errorf("expected ErrRateLimited for 429, got: %v", err)
	}
}

// TestSendRequest404_NotFound verifies that HTTP 404 maps to ErrNotFound.
func TestSendRequest404_NotFound(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound) // 404
		w.Write([]byte(`{"code":404,"message":"resource not found"}`)) //nolint:errcheck
	}))
	defer srv.Close()

	c := newTestClient(srv.URL)
	req, _ := http.NewRequest("GET", srv.URL+"/some/path", nil)
	var body string
	err := c.sendRequest(req, &body)

	if !errors.Is(err, ErrNotFound) {
		t.Errorf("expected ErrNotFound, got: %v", err)
	}
}

// TestSendRequestEmptyBody_NonEmpty verifies that a non-200 response with an
// empty body still produces a non-empty Error() string. This is the regression
// for the old '%d-on-body' bug.
func TestSendRequestEmptyBody_NonEmpty(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(503)
		// empty body
	}))
	defer srv.Close()

	c := newTestClient(srv.URL)
	req, _ := http.NewRequest("GET", srv.URL+"/some/path", nil)
	var body string
	err := c.sendRequest(req, &body)

	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if err.Error() == "" {
		t.Error("Error() must not be empty for empty-body non-200")
	}
	if !strings.Contains(err.Error(), "503") {
		t.Errorf("Error() should contain '503', got: %q", err.Error())
	}
}

// --- unexpectedError path tests (via TransactionOutputFromReference / EvaluateTx) ---

// TestUnexpectedError502_HTMLBody verifies that a 502 with an HTML body returns
// ErrServerError and a non-empty Error() containing a body snippet.
func TestUnexpectedError502_HTMLBody(t *testing.T) {
	htmlBody := `<html><body><h1>502 Bad Gateway</h1><p>The server encountered an error.</p></body></html>`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(502)
		w.Write([]byte(htmlBody)) //nolint:errcheck
	}))
	defer srv.Close()

	c := newTestClient(srv.URL)

	// TransactionOutputFromReference routes through unexpectedError.
	_, err := c.TransactionOutputFromReference("abc123", 0, nil)

	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, ErrServerError) {
		t.Errorf("expected ErrServerError, got: %v (type %T)", err, err)
	}
	if err.Error() == "" {
		t.Error("Error() must not be empty for HTML body response")
	}
	// Should contain a snippet of the body
	if !strings.Contains(err.Error(), "502") {
		t.Errorf("Error() should contain '502', got: %q", err.Error())
	}
	if !strings.Contains(err.Error(), "Bad Gateway") && !strings.Contains(err.Error(), "html") {
		t.Errorf("Error() should contain body snippet, got: %q", err.Error())
	}
}

// TestUnexpectedError404_NotFound verifies 404 via a method using unexpectedError.
func TestUnexpectedError404_NotFound(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"code":404,"message":"not found"}`)) //nolint:errcheck
	}))
	defer srv.Close()

	c := newTestClient(srv.URL)
	_, err := c.TransactionOutputFromReference("deadbeef", 0, nil)

	if !errors.Is(err, ErrNotFound) {
		t.Errorf("expected ErrNotFound, got: %v", err)
	}
}

// TestEvaluateTx502_ServerError verifies the EvaluateTx path (post→unexpectedError).
func TestEvaluateTx502_ServerError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(502)
		w.Write([]byte("upstream connect error")) //nolint:errcheck
	}))
	defer srv.Close()

	c := newTestClient(srv.URL)
	_, err := c.EvaluateTx("deadbeef")

	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, ErrServerError) {
		t.Errorf("expected ErrServerError, got: %v", err)
	}
}

// TestAPIError_Is verifies the Is method directly for all sentinels.
func TestAPIError_Is(t *testing.T) {
	tests := []struct {
		name   string
		status int
		target error
		want   bool
	}{
		{"402→RateLimited", 402, ErrRateLimited, true},
		{"429→RateLimited", 429, ErrRateLimited, true},
		{"404→NotFound", 404, ErrNotFound, true},
		{"500→ServerError", 500, ErrServerError, true},
		{"503→ServerError", 503, ErrServerError, true},
		{"400→NotRateLimited", 400, ErrRateLimited, false},
		{"400→NotServerError", 400, ErrServerError, false},
		{"any→UnexpectedStatus", 400, ErrUnexpectedStatus, true},
		{"402→UnexpectedStatus", 402, ErrUnexpectedStatus, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			e := &APIError{StatusCode: tc.status, Message: "test"}
			got := errors.Is(e, tc.target)
			if got != tc.want {
				t.Errorf("errors.Is(&APIError{%d}, %v) = %v, want %v", tc.status, tc.target, got, tc.want)
			}
		})
	}
}

// TestAPIError_ErrorAlwaysNonEmpty verifies that Error() is never empty
// regardless of Message and Body being empty.
func TestAPIError_ErrorAlwaysNonEmpty(t *testing.T) {
	cases := []struct {
		name   string
		apiErr *APIError
	}{
		{"both empty", &APIError{StatusCode: 502}},
		{"message set", &APIError{StatusCode: 402, Message: "quota exceeded"}},
		{"body set", &APIError{StatusCode: 503, Body: "service unavailable"}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := tc.apiErr.Error()
			if s == "" {
				t.Error("Error() must not return empty string")
			}
			if !strings.Contains(s, "50") && !strings.Contains(s, "40") {
				t.Errorf("Error() should contain status code substring, got: %q", s)
			}
		})
	}
}
