// https://youtu.be/dOonV4byDEg?si=Jd_gCsaqGFcfy1sB
package main

import "testing"

func Test_maxWindowSumSubarrayOfK(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				nums: []int{8, 3, -2, 4, 5, -1, 0, 5, 3, 9, 6},
				k:    5,
			},
			want: 23,
		},
		{
			name: "Case 2",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
				k:    2,
			},
			want: 9,
		},
		{
			name: "Case 3",
			args: args{
				nums: []int{-1, -2, -3, -4, -5},
				k:    3,
			},
			want: -6,
		},
		{
			name: "Case 4",
			args: args{
				nums: []int{5, 5, 5, 5, 5},
				k:    1,
			},
			want: 5,
		},
		{
			name: "Case 5",
			args: args{
				nums: []int{10, -1, 2, 3, -4, 5, 6},
				k:    4,
			},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxWindowSumSubarrayOfK(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maxWindowSumSubarrayOfK() = %v, want %v", got, tt.want)
			}
		})
	}
}
