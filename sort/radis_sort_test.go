package sort

import (
	"reflect"
	"testing"
)

func TestRadixSort(t *testing.T) {
	type args struct {
		A IntSlice
		d int
	}
	tests := []struct {
		name string
		args args
		want IntSlice
	}{
		{
			name: "329, 457, 657, 839, 436, 720, 355",
			args: args{
				A: IntSlice{329, 457, 657, 839, 436, 720, 355},
				d: 3,
			},
			want: IntSlice{329, 355, 436, 457, 657, 720, 839},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RadixSort(tt.args.A, tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RadixSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
