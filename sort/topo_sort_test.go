package sort

import (
	"reflect"
	"testing"
)

func TestTopoSort(t *testing.T) {
	type args struct {
		vs []GraphVertex
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "A-B, B-C, A-C, B-D, C-D, B-E",
			args: args{
				vs: []GraphVertex{
					{Idx: 0, Name: "A", EdgeVertexIdx: []int{1, 2}},
					{Idx: 1, Name: "B", EdgeVertexIdx: []int{2, 3, 4}},
					{Idx: 2, Name: "C", EdgeVertexIdx: []int{3}},
					{Idx: 3, Name: "D", EdgeVertexIdx: []int{}},
					{Idx: 4, Name: "E", EdgeVertexIdx: []int{}},
				},
			},
			want: []int{0, 1, 4, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TopoSort(tt.args.vs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TopoSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
