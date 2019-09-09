package dynamicProgramming

import (
	"testing"
)

func TestFoolishCutRod(t *testing.T) {
	type args struct {
		price []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "4m",
			args: args{
				price: RodPrice,
				n:     4,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FoolishCutRod(tt.args.price, tt.args.n); got != tt.want {
				t.Errorf("FoolishCutRod() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMemorialCutRod(t *testing.T) {
	type args struct {
		price []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "4m",
			args: args{
				price: RodPrice,
				n:     4,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MemorialCutRod(tt.args.price, tt.args.n); got != tt.want {
				t.Errorf("MemorialCutRod() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBottomUpCutRod(t *testing.T) {
	type args struct {
		price []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "4m",
			args: args{
				price: RodPrice,
				n:     4,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BottomUpCutRod(tt.args.price, tt.args.n); got != tt.want {
				t.Errorf("BottomUpCutRod() = %v, want %v", got, tt.want)
			}
		})
	}
}
