package leetcode

import "testing"

func TestDistributeCandies(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		limit int
		want  int64
	}{
		{
			name:  "test 1",
			n:     5,
			limit: 2,
			want:  3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistributeCandies(tt.n, tt.limit); got != tt.want {
				t.Errorf("DistributeCandies() = %v, want %v", got, tt.want)
			}
		})
	}
}
