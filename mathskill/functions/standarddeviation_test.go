package functions

import "testing"

func TestComputeStandardDeviation(t *testing.T) {
	type args struct {
		variance float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Positive variance",
			args: args{
				variance: 4.0,
			},
			want: 2.0,
		},
		{
			name: "Zero variance",
			args: args{
				variance: 0.0,
			},
			want: 0.0,
		},
		{
			name: "Negative variance",
			args: args{
				variance: -2.0,
			},
			want: 0.0, // Standard deviation is always non-negative
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeStandardDeviation(tt.args.variance); got != tt.want {
				t.Errorf("ComputeStandardDeviation() = %v, want %v", got, tt.want)
			}
		})
	}
}
