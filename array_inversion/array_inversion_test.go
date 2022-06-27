// package array_inversion
package main

import "testing"

func Test_count_inversions(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "first case",
			args: []int{},
			want: 0,
		},
		{
			name: "second case",
			args: []int{1},
			want: 0,
		},
		{
			name: "third case",
			args: []int{1, 1},
			want: 0,
		},
		{
			name: "fourth case",
			args: []int{1, 2},
			want: 0,
		},
		{
			name: "fifth case",
			args: []int{2, 1},
			want: 1,
		},
		{
			name: "sixth case",
			args: []int{0, 1, 2},
			want: 0,
		},
		{
			name: "seven case",
			args: []int{2, 1, 0},
			want: 3,
		},
		{
			name: "8 case",
			args: []int{1, 3, 5, 2, 4, 6},
			want: 3,
		},
		{
			name: "9 case",
			args: []int{10, 9, 8, 7, 6, 5},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, got := sort_and_count_inversions(tt.args); got != tt.want {
				t.Errorf("count_inversions() = %v, want %v", got, tt.want)
			}
		})
	}
}
