package most_frequent

import "testing"

func TestMostFrequent(t *testing.T) {
	type args struct {
		nums []int
		key  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				nums: []int{1, 100, 200, 1, 100},
				key:  1,
			},
			want: 100,
		},
		{
			name: "Case 2",
			args: args{
				nums: []int{2, 2, 2, 2, 3},
				key:  2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MostFrequent(tt.args.nums, tt.args.key); got != tt.want {
				t.Errorf("MostFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
