package leetcode

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  bool
	}{
		{
			name:  "test true palindrone",
			input: 1221,
			want:  true,
		},
		{
			name:  "test negative is false",
			input: -121,
			want:  false,
		},
		{
			name:  "test 10",
			input: 10,
			want:  false,
		},
		{
			name:  "test 1",
			input: 1,
			want:  true,
		},
		{
			name:  "test 0",
			input: 0,
			want:  true,
		},
		{
			name:  "test 1000000001",
			input: 1000000001,
			want:  true,
		},
		{
			name:  "test 1000030001",
			input: 1000030001,
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.input); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
