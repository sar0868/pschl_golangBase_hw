package files

import (
	"errors"
	"os"
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test_checkTypeFile(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     bool
	}{
		{
			name:     "Positive test (file.json)",
			filename: "file.json",
			want:     true,
		},
		{
			name:     "Positive test(file.dd.json)",
			filename: "file.dd.json",
			want:     true,
		},
		{
			name:     "Negative test (file)",
			filename: "file",
			want:     false,
		},
		{
			name:     "Negative test (file.json.txt)",
			filename: "file.json.txt",
			want:     false,
		},
		{
			name:     "Negative test ",
			filename: "",
			want:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkTypeFile(tt.filename)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestRead(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		want    []byte
		wantErr bool
	}{
		{
			name:    "positive test",
			path:    "file.json",
			want:    []byte("hello"),
			wantErr: false,
		},
		{
			name:    "negative test",
			path:    "file",
			want:    []byte("hello"),
			wantErr: true,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			err := createTestData(tt.path)
			if err != nil {
				t.Error(err.Error())
			}
			defer os.Remove(tt.path)
			got, err := Read(tt.path)
			if err != nil {
				assert.Equal(t, tt.wantErr, err != nil)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func createTestData(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := file.Write([]byte("hello")); err != nil {
		return errors.New("error write test file")
	}
	return nil
}
