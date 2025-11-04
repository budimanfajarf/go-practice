package largest_good_int

import "testing"

func TestLargestGoodInteger(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Case 1",
			args: args{
				num: "6777133339",
			},
			want: "777",
		},
		{
			name: "Case 2",
			args: args{
				num: "2300019",
			},
			want: "000",
		},
		{
			name: "Case 3",
			args: args{
				num: "42352338",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LargestGoodInteger(tt.args.num); got != tt.want {
				t.Errorf("LargestGoodInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
