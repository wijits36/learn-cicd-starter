package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name      string
		header    http.Header
		wantKey   string
		wantErr   error
		expectErr bool
	}{
		{
			name:    "valid ApiKey header",
			header:  http.Header{"Authorization": []string{"ApiKey my-secret-key"}},
			wantKey: "my-secret-key",
		},
		{
			name:      "missing Authorization header",
			header:    http.Header{},
			wantErr:   ErrNoAuthHeaderIncluded,
			expectErr: true,
		},
		{
			name:      "wrong scheme",
			header:    http.Header{"Authorization": []string{"Bearer my-secret-key"}},
			expectErr: true,
		},
		{
			name:      "missing key value",
			header:    http.Header{"Authorization": []string{"ApiKey"}},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.header)

			if tt.expectErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				if tt.wantErr != nil && !errors.Is(err, tt.wantErr) {
					t.Fatalf("expected error %v, got %v", tt.wantErr, err)
				}
				if got != "" {
					t.Fatalf("expected empty key on error, got %q", got)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.wantKey {
				t.Fatalf("expected key %q, got %q", tt.wantKey, got)
			}
		})
	}
}
