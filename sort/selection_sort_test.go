package sort

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	type args struct {
		items SelectionSortInterface
	}
	tests := []struct {
		name string
		args args
		want SelectionSortInterface
	}{
		{
			name: "1,2,3,4",
			args: args{
				items: &IntSlice{1, 4, 3, 2},
			},
			want: &IntSlice{1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelectionSort(tt.args.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
