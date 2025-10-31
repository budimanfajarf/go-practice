package easy

import "testing"

func TestNumberOfSteps(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				num: 14,
			},
			want: 6,
		},
		{
			name: "Case 2",
			args: args{
				num: 8,
			},
			want: 4,
		},
		{
			name: "Case 3",
			args: args{
				num: 123,
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumberOfSteps(tt.args.num); got != tt.want {
				t.Errorf("NumberOfSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
