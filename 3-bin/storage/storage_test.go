package storage

import (
	"binjson/bins"
	"binjson/files"
	"os"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
)

func TestSaveBinToJson(t *testing.T) {
	tests := []struct {
		name    string
		bin     bins.Bin
		path    string
		want    []byte
		wantErr bool
	}{
		{
			name: "save data in file",
			bin: bins.Bin{
				Id:        "1",
				Private:   true,
				CreatedAt: time.Date(2025, time.December, 1, 12, 0, 0, 0, time.UTC),
			},
			path:    "test.json",
			want:    []byte("{\"id\":\"1\",\"private\":true,\"createdAt\":\"2025-12-01T12:00:00Z\",\"name\":\"\"}"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := SaveBinToJson(tt.bin, tt.path)
			assert.Equal(t, tt.wantErr, gotErr != nil)
			if _, err := os.Stat(tt.path); os.IsNotExist(err) {
				t.Errorf("file %v don't create", tt.path)
			}
			file, err := files.Read(tt.path)
			if err != nil {
				t.Errorf("error read file: %v", tt.path)
			}
			assert.Equal(t, string(tt.want), string(file))
			err = os.Remove(tt.path)
			if err != nil {
				t.Error(err.Error())
			}
		})
	}
}

func TestGetBinListToJson(t *testing.T) {
	name := "positive test"
	path := "test,json"
	want := &bins.BinList{
		Bins: []bins.Bin{
			*bins.NewBin("1", true, "first"),
			*bins.NewBin("2", true, "second"),
		},
	}
	data :=

		t.Run(name, func(t *testing.T) {
			got, gotErr := GetBinListToJson(path)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetBinListToJson() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GetBinListToJson() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("GetBinListToJson() = %v, want %v", got, tt.want)
			}
		})
}
