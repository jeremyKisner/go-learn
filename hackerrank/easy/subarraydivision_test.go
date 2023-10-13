package hackerrank

import "testing"

func TestBirthday(t *testing.T) {
	type args struct {
		s []int32
		d int32
		m int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Birthday(tt.args.s, tt.args.d, tt.args.m); got != tt.want {
				t.Errorf("Birthday() = %v, want %v", got, tt.want)
			}
		})
	}
}
