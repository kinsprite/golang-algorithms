package sort

import (
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {
	type args struct {
		A IntSlice
	}
	tests := []struct {
		name string
		args args
		want IntSlice
	}{
		{
			name: "1,4,3,2",
			args: args{
				A: IntSlice{1, 4, 3, 2},
			},
			want: IntSlice{1, 2, 3, 4},
		},
		{
			name: "2, 8, 7, 1, 3, 5, 6, 4",
			args: args{
				A: IntSlice{2, 8, 7, 1, 3, 5, 6, 4},
			},
			want: IntSlice{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HeapSort(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HeapSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
