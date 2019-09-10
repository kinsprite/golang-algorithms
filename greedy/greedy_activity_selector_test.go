package greedy

import (
	"reflect"
	"testing"
)

func TestActivitySelector(t *testing.T) {
	type args struct {
		activities []Activity
	}
	tests := []struct {
		name string
		args args
		want []Activity
	}{
		{
			name: "Activities",
			args: args{
				activities: []Activity{
					{"a1", 1, 4},
					{"a2", 3, 5},
					{"a3", 0, 6},
					{"a4", 5, 7},
					{"a5", 3, 9},
					{"a6", 5, 9},
					{"a7", 6, 10},
					{"a8", 8, 11},
					{"a9", 8, 12},
					{"a10", 2, 14},
					{"a11", 12, 16},
				},
			},
			want: []Activity{
				{"a1", 1, 4},
				{"a4", 5, 7},
				{"a8", 8, 11},
				{"a11", 12, 16},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ActivitySelector(tt.args.activities); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ActivitySelector() = %v, want %v", got, tt.want)
			}
		})
	}
}
