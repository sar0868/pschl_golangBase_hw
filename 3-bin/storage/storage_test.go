package storage

import (
	"binjson/bins"
	"binjson/files"
	"os"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
)

var expectedJson string = "{\"Bins\":[{\"id\":\"1\",\"private\":true,\"createdAt\":\"2025-12-01T12:00:00Z\",\"name\":\"bin1\"},{\"id\":\"2\",\"private\":false,\"createdAt\":\"2025-12-01T12:00:00Z\",\"name\":\"bin2\"}]}"

func TestSaveBinToJson(t *testing.T) {
	name := "save data in file"

	bins := bins.BinList{
		Bins: []bins.Bin{
			bins.Bin{
				Id:        "1",
				Private:   true,
				CreatedAt: time.Date(2025, time.December, 1, 12, 0, 0, 0, time.UTC),
				Name:      "bin1",
			},
			bins.Bin{
				Id:        "2",
				Private:   false,
				CreatedAt: time.Date(2025, time.December, 1, 12, 0, 0, 0, time.UTC),
				Name:      "bin2",
			},
		},
	}
	path := "test.json"
	want := []byte(expectedJson)
	wantErr := false

	t.Run(name, func(t *testing.T) {
		gotErr := SaveBinToJson(bins, path)
		assert.Equal(t, wantErr, gotErr != nil)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			t.Errorf("file %v don't create", path)
		}
		file, err := files.Read(path)
		if err != nil {
			t.Errorf("error read file: %v", path)
		}
		assert.Equal(t, string(want), string(file))
		err = os.Remove(path)
		if err != nil {
			t.Error(err.Error())
		}
	})
}

func TestGetBinListToJson(t *testing.T) {
	name := "positive test"
	path := "expected.json"
	want := bins.BinList{
		Bins: []bins.Bin{
			bins.Bin{
				Id:        "1",
				Private:   true,
				CreatedAt: time.Date(2025, time.December, 1, 12, 0, 0, 0, time.UTC),
				Name:      "bin1",
			},
			bins.Bin{
				Id:        "2",
				Private:   false,
				CreatedAt: time.Date(2025, time.December, 1, 12, 0, 0, 0, time.UTC),
				Name:      "bin2",
			},
		},
	}
	wantLen := 2
	wantErr := false
	err := writeExpectedFile(path, []byte(expectedJson))
	if err != nil {
		t.Errorf("error write expected file: %v", err)
	}

	t.Run(name, func(t *testing.T) {
		got, gotErr := GetBinListToJson(path)
		if gotErr != nil {
			if !wantErr {
				t.Errorf("GetBinListToJson() failed: %v", gotErr)
			}
			return
		}
		assert.Equal(t, want, got)
		assert.Equal(t, wantLen, len(got.Bins))
	})
	err = os.Remove(path)
	if err != nil {
		t.Error(err.Error())
	}
}

func writeExpectedFile(path string, data []byte) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err = file.Write(data); err != nil {
		return err
	}
	return nil
}
