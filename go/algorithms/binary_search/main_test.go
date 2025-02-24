package main

import "testing"

func Test_binSearch(t *testing.T) {
	type args struct {
		source []int64
		target int64
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Target in middle",
			args: args{source: []int64{1, 3, 5, 7, 9}, target: 5},
			want: 2,
		},
		{
			name: "Target at beginning",
			args: args{source: []int64{1, 3, 5, 7, 9}, target: 1},
			want: 0,
		},
		{
			name: "Target at end",
			args: args{source: []int64{1, 3, 5, 7, 9}, target: 9},
			want: 4,
		},
		{
			name: "Target not in list",
			args: args{source: []int64{1, 3, 5, 7, 9}, target: 6},
			want: -1,
		},
		{
			name: "Empty list",
			args: args{source: []int64{}, target: 5},
			want: -1,
		},
		{
			name: "Single element - found",
			args: args{source: []int64{5}, target: 5},
			want: 0,
		},
		{
			name: "Single element - not found",
			args: args{source: []int64{5}, target: 3},
			want: -1,
		},
		{
			name: "Target in even-sized list",
			args: args{source: []int64{2, 4, 6, 8, 10, 12}, target: 8},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := binSearch(tt.args.source, tt.args.target); got != tt.want {
				t.Errorf("binSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
