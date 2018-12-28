package modules

import "testing"

func Test_add(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name  string
		args  args
		wantZ int
	}{
		// Add test cases.
		{"1", args{1, 1}, 2},
		{"1", args{2, 3}, 5},
		{"1", args{5, 6}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotZ := add(tt.args.x, tt.args.y); gotZ != tt.wantZ {
				t.Errorf("add() = %v, want %v", gotZ, tt.wantZ)
			}
		})
	}
}
