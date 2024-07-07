package functions

import "testing"

func TestComputeVariance(t *testing.T) {
	type args struct {
		numbers []float64
		average float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeVariance(tt.args.numbers, tt.args.average); got != tt.want {
				t.Errorf("ComputeVariance() = %v, want %v", got, tt.want)
			}
		})
	}
}
