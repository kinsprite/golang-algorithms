package sort

import (
	"reflect"
	"testing"
)

func TestBucketSort(t *testing.T) {
	type args struct {
		A IntSlice
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
			},
			want: IntSlice{329, 355, 436, 457, 657, 720, 839},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BucketSort(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BucketSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
