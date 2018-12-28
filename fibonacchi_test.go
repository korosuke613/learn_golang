package plactice

import (
	"reflect"
	"testing"
)

func Test_fibonacci(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		// Add test cases.
		{"1", 1},
		{"2", 1},
		{"3", 2},
		{"4", 3},
		{"5", 5},
	}
	f := fibonacci()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
