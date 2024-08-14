package algebra

import "testing"

func TestOrderedPair_Quadrant(t *testing.T) {
	tests := []struct {
		name string
		o    *OrderedPair
		want string
	}{
		{
			"test 1",
			&OrderedPair{
				X: 0,
				Y: 0,
			},
			"Origin",
		},
		{
			"test 2",
			&OrderedPair{
				X: -4,
				Y: 1,
			},
			"II",
		},
		{
			"test 3",
			&OrderedPair{
				X: -2,
				Y: 3,
			},
			"II",
		},
		{
			"test 4",
			&OrderedPair{
				X: 2,
				Y: -5,
			},
			"IV",
		},
		{
			"test 5",
			&OrderedPair{
				X: -2,
				Y: 5,
			},
			"II",
		},
		{
			"test 6",
			&OrderedPair{
				X: -3,
				Y: (5 / 2),
			},
			"II",
		},
		{
			"test 6",
			&OrderedPair{
				X: 1,
				Y: 1,
			},
			"I",
		},
		{
			"test 7",
			&OrderedPair{
				X: 1,
				Y: 0,
			},
			"I and IV",
		},
		{
			"test 8",
			&OrderedPair{
				X: 0,
				Y: 1,
			},
			"I and II",
		},
		{
			"test 9",
			&OrderedPair{
				X: -1,
				Y: 0,
			},
			"II and III",
		},
		{
			"test 10",
			&OrderedPair{
				X: 0,
				Y: -1,
			},
			"III and IV",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.Quadrant(); got != tt.want {
				t.Errorf("OrderedPair.Quadrant() = %v, want %v", got, tt.want)
			}
		})
	}
}
