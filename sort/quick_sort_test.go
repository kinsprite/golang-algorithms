package sort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	type args struct {
		items IntSlice
	}
	tests := []struct {
		name string
		args args
		want IntSlice
	}{
		{
			name: "1,4,3,2",
			args: args{
				items: IntSlice{1, 4, 3, 2},
			},
			want: IntSlice{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort(tt.args.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuickSort2(t *testing.T) {
	type args struct {
		items IntSlice
	}
	tests := []struct {
		name string
		args args
		want IntSlice
	}{
		{
			name: "1,4,3,2",
			args: args{
				items: IntSlice{1, 4, 3, 2},
			},
			want: IntSlice{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort2(tt.args.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort2() = %v, want %v", got, tt.want)
			}
		})
	}
}
