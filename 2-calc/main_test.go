package main

import (
	"math"
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test_sum(t *testing.T) {
	tests := []struct {
		name string
		arr  []float64
		want float64
	}{
		{
			name: "3+5+2=10",
			arr:  []float64{3, 5, 2},
			want: 10,
		},
		{
			name: "-3+5+2=4",
			arr:  []float64{-3, 5, 2},
			want: 4,
		},
		{
			name: "0+1+1=2",
			arr:  []float64{0.0, 1, 1},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sum(tt.arr)
			assert.Equal(t, tt.want, got)
		})
	}
}
func Test_avg(t *testing.T) {
	tests := []struct {
		name    string
		arr     []float64
		want    float64
		epsilon float64
	}{
		{
			name:    "avg [3, 5, 2] = 3.33",
			arr:     []float64{3, 5, 2},
			want:    3.33,
			epsilon: 1e-2,
		},
		{
			name:    "avg [-3, 5, 2] = 1.33",
			arr:     []float64{-3, 5, 2},
			want:    1.33,
			epsilon: 1e-2,
		},
		{
			name:    "avg [0, 1, 1] = 0,667",
			arr:     []float64{0.0, 1, 1},
			want:    0.66,
			epsilon: 1e-2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := avg(tt.arr)
			if math.Abs(tt.want-got) > tt.epsilon {
				t.Errorf("for %v expected %.2f, actual %f", tt.arr, tt.want, got)
			}
		})
	}
}
func Test_med(t *testing.T) {
	tests := []struct {
		name string
		arr  []float64
		want float64
	}{
		{
			name: "median [3, 5, 2] = 3",
			arr:  []float64{3, 5, 2},
			want: 3,
		},
		{
			name: "median [-3, 5, 2] = 2",
			arr:  []float64{-3, 5, 2},
			want: 2,
		},
		{
			name: "median [1, 3, 2, 4] =2.5",
			arr:  []float64{1, 3, 2, 4},
			want: 2.5,
		},
		{
			name: "median [1, 1, 1, 1] =1",
			arr:  []float64{1, 1, 1, 1},
			want: 1,
		},
		{
			name: "median [1, 0, 0, 0] = 0",
			arr:  []float64{1, 0, 0, 0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := med(tt.arr)
			assert.Equal(t, tt.want, got)
		})
	}
}
