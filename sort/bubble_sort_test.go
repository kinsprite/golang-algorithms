package sort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	type args struct {
		items BubbleSortInterface
	}
	tests := []struct {
		name string
		args args
		want BubbleSortInterface
	}{
		{
			name: "4,3,2,1,0",
			args: args{
				items: &IntSlice{4, 3, 2, 1, 0},
			},
			want: &IntSlice{0, 1, 2, 3, 4},
		},
		{
			name: "4,2,1,3,0",
			args: args{
				items: &IntSlice{4, 2, 1, 3, 0},
			},
			want: &IntSlice{0, 1, 2, 3, 4},
		},
		{
			name: "0,2,1,3,4,5",
			args: args{
				items: &IntSlice{0, 2, 1, 3, 4, 5},
			},
			want: &IntSlice{0, 1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BubbleSort(tt.args.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
