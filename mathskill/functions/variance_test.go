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
		{
			name: "Test with single number",
			args: args{
				numbers: []float64{10.0},
				average: 10.0,
			},
			want: 0.0,
		},
		{
			name: "Test with two numbers",
			args: args{
				numbers: []float64{5.0, 10.0},
				average: 7.5,
			},
			want: 6.25,
		},
		{
			name: "Test with multiple numbers",
			args: args{
				numbers: []float64{2.0, 4.0, 6.0, 8.0, 10.0},
				average: 6.0,
			},
			want: 8.0,
		},
		{
			name: "Test with zero average",
			args: args{
				numbers: []float64{-2.0, 0.0, 2.0},
				average: 0.0,
			},
			want: 2.6666666666666665,
		},
		{
			name: "Test with empty slice",
			args: args{
				numbers: []float64{},
				average: 0.0,
			},
			want: 0.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeVariance(tt.args.numbers, tt.args.average); got != tt.want {
				t.Errorf("ComputeVariance() = %v, want %v", got, tt.want)
			}
		})
	}
}
