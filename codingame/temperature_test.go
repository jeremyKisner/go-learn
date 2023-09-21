package main

import (
	"reflect"
	"testing"
)

func TestGetLowestTemperature(t *testing.T) {
	type args struct {
		n      int
		inputs []string
	}
	tests := []struct {
		name      string
		args      args
		wantValue string
	}{
		{
			name: "test variety",
			args: args{
				n:      5,
				inputs: []string{"-2", "-8", "4", "5", "1"},
			},
			wantValue: "1",
		},
		{
			name: "test only negative",
			args: args{
				n:      5,
				inputs: []string{"-2", "-8", "-4", "-5", "-1"},
			},
			wantValue: "-1",
		},
		{
			name: "test only postive",
			args: args{
				n:      5,
				inputs: []string{"2", "8", "4", "5", "1"},
			},
			wantValue: "1",
		},
		{
			name: "test only postive with 0",
			args: args{
				n:      5,
				inputs: []string{"2", "8", "4", "5", "1", "0"},
			},
			wantValue: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetClosestTemperatureToZero(tt.args.n, tt.args.inputs); !reflect.DeepEqual(got, tt.wantValue) {
				t.Errorf("GetLowestTemperature() got %v, want %v", got, tt.wantValue)
			}
		})
	}
}
