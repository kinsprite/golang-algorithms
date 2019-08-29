package sort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	type args struct {
		items IntSlice
	}
	tests := []struct {
		name string
		args args
		want IntSlice
	}{
		{
			name: "1,2,3,4",
			args: args{
				items: IntSlice{1, 4, 3, 2},
			},
			want: IntSlice{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSort(tt.args.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
