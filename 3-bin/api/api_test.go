package api_test

import (
	"binjson/api"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestGetBin(t *testing.T) {
	tests := []struct {
		name   string
		path   string
		status int
	}{
		{
			name:   "test",
			path:   "https://api.jsonbin.io/v3/b/693ff04443b1c97be9efa06a",
			status: 200,
		},
	}
	for _, tt := range tests {
		api.GetKey()
		t.Run(tt.name, func(t *testing.T) {
			result := api.GetBin(tt.path)
			assert.Equal(t, tt.status, result)
		})
	}
}
