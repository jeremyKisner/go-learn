package leetcode

import (
	"testing"
)

func Test_evaluate(t *testing.T) {
	type args struct {
		s         string
		knowledge [][]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test 1",
			args: args{
				s: "(name)is(age)yearsold",
				knowledge: [][]string{
					{"name", "bob"},
					{"age", "two"},
				},
			},
			want: "bobistwoyearsold",
		},
		{
			name: "test 2",
			args: args{
				s: "hi(name)",
				knowledge: [][]string{
					{"hi(name)"},
				},
			},
			want: "hi?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evaluate(tt.args.s, tt.args.knowledge); got != tt.want {
				t.Errorf("evaluate() got %v, want %v", got, tt.want)
			}
		})
	}
}
