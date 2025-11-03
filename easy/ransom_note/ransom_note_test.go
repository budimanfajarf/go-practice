package ransom_note

import "testing"

func TestCanConstruct(t *testing.T) {
	type args struct {
		ransomNote string
		magazine   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Case 1",
			args: args{
				ransomNote: "a",
				magazine:   "b",
			},
			want: false,
		},
		{
			name: "Case 2",
			args: args{
				ransomNote: "aa",
				magazine:   "ab",
			},
			want: false,
		},
		{
			name: "Case 3",
			args: args{
				ransomNote: "aa",
				magazine:   "aab",
			},
			want: true,
		},
		{
			name: "Case 4",
			args: args{
				ransomNote: "aab",
				magazine:   "baa",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanConstruct(tt.args.ransomNote, tt.args.magazine); got != tt.want {
				t.Errorf("CanConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
