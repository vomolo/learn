package functions

import "testing"

func TestComputeAverage(t *testing.T) {
	type args struct {
		numbers []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Empty slice",
			args: args{
				numbers: []float64{},
			},
			want: 0.0,
		},
		{
			name: "Single element",
			args: args{
				numbers: []float64{10.0},
			},
			want: 10.0,
		},
		{
			name: "Multiple elements",
			args: args{
				numbers: []float64{5.0, 10.0, 15.0},
			},
			want: 10.0,
		},
		{
			name: "Negative numbers",
			args: args{
				numbers: []float64{-5.0, 0.0, 5.0},
			},
			want: 0.0,
		},
		{
			name: "Floating-point numbers",
			args: args{
				numbers: []float64{2.5, 3.7, 4.2},
			},
			want: 3.4666666666666667,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeAverage(tt.args.numbers); got != tt.want {
				t.Errorf("ComputeAverage() = %v, want %v", got, tt.want)
			}
		})
	}
}
