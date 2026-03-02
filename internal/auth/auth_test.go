package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// Create new mock request
	req := httptest.NewRequest(http.MethodGet, "/test", nil)

	// Test for token format
	req.Header.Set("Authorization", "_")

	_, err := GetAPIKey(req.Header)
	if err == nil {
		t.Fatal("must have failed with an malformed Authorization token")
	}
	// Mock empty auth header
	req.Header.Set("Authorization", "")

	// Test empty token, must fail
	_, err = GetAPIKey(req.Header)
	if err != ErrNoAuthHeaderIncluded {
		t.Fatal("must have failed with empty Authorization token")
	}

}
