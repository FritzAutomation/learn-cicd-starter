package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// Case 1: No Authorization header
	t.Run("No Authorization Header", func(t *testing.T) {
		headers := http.Header{}
		_, err := GetAPIKey(headers)
		if err != ErrNoAuthHeaderIncluded {
			t.Errorf("expected error %v, got %v", ErrNoAuthHeaderIncluded, err)
		}
	})

	// Case 2: Malformed Authorization header
	t.Run("Malformed Authorization Header", func(t *testing.T) {
		headers := http.Header{}
		headers.Set("Authorization", "Bearer token")
		_, err := GetAPIKey(headers)
		if err == nil || err.Error() != "malformed authorization header" {
			t.Errorf("expected error 'malformed authorization header', got %v", err)
		}
	})

	// Case 3: Correct Authorization header
	t.Run("Correct Authorization Header", func(t *testing.T) {
		headers := http.Header{}
		headers.Set("Authorization", "ApiKey valid_api_key")
		apiKey, err := GetAPIKey(headers)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if apiKey != "valid_api_key" {
			t.Errorf("expected apiKey 'valid_api_key', got %v", apiKey)
		}
	})
}
