// https://youtu.be/Dd_NgYVOdLk?si=ld32mWM4eCUB_xhL
// https://gist.github.com/JyotinderSingh/d2bd0096e146aa3083442ceb48eab6b4
package edit_distance

import "testing"

func TestMinDistance(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				word1: "horse",
				word2: "ros",
			},
			want: 3,
		},
		{
			name: "Case 2",
			args: args{
				word1: "intention",
				word2: "execution",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinDistance(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("MinDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
