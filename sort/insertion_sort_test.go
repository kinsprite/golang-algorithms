package sort

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	type args struct {
		items IntSlice
	}
	tests := []struct {
		name string
		args args
		want IntSlice
	}{
		{
			name: "5,2,4,6,1,3",
			args: args{IntSlice{5, 2, 4, 6, 1, 3}},
			want: IntSlice{1, 2, 3, 4, 5, 6},
		},
		{
			name: "5,2,4,6,1,3,3",
			args: args{IntSlice{5, 2, 4, 6, 1, 3, 3}},
			want: IntSlice{1, 2, 3, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertionSort(tt.args.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
