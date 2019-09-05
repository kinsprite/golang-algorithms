package sort

import (
	"reflect"
	"testing"
)

func TestShellSort(t *testing.T) {
	type args struct {
		arr IntSlice
	}
	tests := []struct {
		name string
		args args
		want IntSlice
	}{
		{
			name: "329, 457, 657, 839, 436, 720, 355",
			args: args{
				arr: IntSlice{329, 457, 657, 839, 436, 720, 355},
			},
			want: IntSlice{329, 355, 436, 457, 657, 720, 839},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShellSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShellSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
