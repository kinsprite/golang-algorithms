package dynamicProgramming

import (
	"testing"
)

func TestDoMatrixChainOrder(t *testing.T) {
	type args struct {
		p []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "A1[30x35],A2[35x15],A3[15,5],A4[5x10],A5[10x20],A6[20,25]",
			args: args{
				p: []int{30, 35, 15, 5, 10, 20, 25},
			},
			want: "((A1(A2A3))((A4A5)A6))",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DoMatrixChainOrder(tt.args.p); got != tt.want {
				t.Errorf("DoMatrixChainOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
