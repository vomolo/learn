package functions

import "testing"

func TestComputeMedian(t *testing.T) {
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
			name: "Single element slice",
			args: args{
				numbers: []float64{42.0},
			},
			want: 42.0,
		},
		{
			name: "Odd-length slice",
			args: args{
				numbers: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
			},
			want: 3.0,
		},
		{
			name: "Even-length slice",
			args: args{
				numbers: []float64{1.0, 2.0, 3.0, 4.0},
			},
			want: 2.5,
		},
		{
			name: "Negative numbers",
			args: args{
				numbers: []float64{-5.0, -2.0, 0.0, 2.0, 5.0},
			},
			want: 0.0,
		},
		{
			name: "Duplicate numbers",
			args: args{
				numbers: []float64{1.0, 1.0, 2.0, 2.0, 3.0},
			},
			want: 2.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeMedian(tt.args.numbers); got != tt.want {
				t.Errorf("ComputeMedian() = %v, want %v", got, tt.want)
			}
		})
	}
}
