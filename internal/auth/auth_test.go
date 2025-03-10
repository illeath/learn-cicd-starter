package auth

import (
	"net/http"
	"testing"
)

func TestGetAPI(t *testing.T) {
	tests := []struct {
		name           string
		authHeader     string
		expectedAPIKey string
		expectError    bool
	}{
		{
			name:           "valid auth header",
			authHeader:     "ApiKey abc123",
			expectedAPIKey: "abc123",
			expectError:    false,
		},
		{
			name:           "no auth header",
			authHeader:     "",
			expectedAPIKey: "",
			expectError:    true,
		},
		{
			name:           "malformed auth header",
			authHeader:     "ApiKey",
			expectedAPIKey: "",
			expectError:    true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			headers := make(http.Header)
			if tc.authHeader != "" {
				headers.Add("Authorization", tc.authHeader)
			}
			apiKey, err := GetAPIKey(headers)

			if tc.expectError && err == nil {
				t.Errorf("expected error but got none")
				return
			}

			if !tc.expectError && err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if apiKey != tc.expectedAPIKey {
				t.Errorf("expected api key %s but got %s", tc.expectedAPIKey, apiKey)
			}
		})
	}
}
