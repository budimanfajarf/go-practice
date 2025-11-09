package word_ladder

import "testing"

func TestLadderLength(t *testing.T) {
	type args struct {
		beginWord string
		endWord   string
		wordList  []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				beginWord: "hit",
				endWord:   "cog",
				wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			},
			want: 5,
		},
		{
			name: "Case 2",
			args: args{
				beginWord: "hit",
				endWord:   "cog",
				wordList:  []string{"hot", "dot", "dog", "lot", "log"},
			},
			want: 0,
		},
		{
			name: "Case 3",
			args: args{
				beginWord: "a",
				endWord:   "c",
				wordList:  []string{"a", "b", "c"},
			},
			want: 2,
		},
		{
			name: "Case 4",
			args: args{
				beginWord: "hot",
				endWord:   "dog",
				wordList:  []string{"hot", "dog"},
			},
			want: 0,
		},
		{
			name: "Case 5",
			args: args{
				beginWord: "hot",
				endWord:   "dog",
				wordList:  []string{"hot", "dog", "cog", "pot", "dot"},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LadderLength(tt.args.beginWord, tt.args.endWord, tt.args.wordList); got != tt.want {
				t.Errorf("LadderLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
