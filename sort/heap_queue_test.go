package sort

import (
	"reflect"
	"testing"
)

func TestHeapExtractMax(t *testing.T) {
	type args struct {
		A IntSlice
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 IntSlice
	}{
		{
			name: "4, 2, 3, 1",
			args: args{
				A: IntSlice{4, 2, 3, 1},
			},
			want:  4,
			want1: IntSlice{3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := HeapExtractMax(tt.args.A)
			if got != tt.want {
				t.Errorf("HeapExtractMax() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("HeapExtractMax() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestHeapIncreaseKey(t *testing.T) {
	type args struct {
		A   IntSlice
		i   int
		key int
	}
	tests := []struct {
		name string
		args args
		want IntSlice
	}{
		{
			name: "4, 2, 3, 1",
			args: args{
				A:   IntSlice{4, 2, 3, 1},
				i:   2,
				key: 5,
			},
			want: IntSlice{5, 2, 4, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HeapIncreaseKey(tt.args.A, tt.args.i, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HeapIncreaseKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxHeapInsert(t *testing.T) {
	type args struct {
		A   IntSlice
		key int
	}
	tests := []struct {
		name string
		args args
		want IntSlice
	}{
		{
			name: "4, 2, 3, 1",
			args: args{
				A:   IntSlice{4, 2, 3, 1},
				key: 5,
			},
			want: IntSlice{5, 4, 3, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxHeapInsert(tt.args.A, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxHeapInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
