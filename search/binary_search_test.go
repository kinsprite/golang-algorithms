package search

import "testing"

func TestBinarySearch(t *testing.T) {
	type args struct {
		items BinarySearchInterface
		value interface{}
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 int
	}{
		{
			name: "1,2,3,4",
			args: args{
				items: &IntSlice{1, 2, 3, 4},
				value: 3,
			},
			want:  true,
			want1: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := BinarySearch(tt.args.items, tt.args.value)
			if got != tt.want {
				t.Errorf("BinarySearch() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("BinarySearch() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		items := &IntSlice{1, 2, 3, 4, 5}
		value := 3
		BinarySearch(items, value)
	}
}

func BenchmarkBinarySearchParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			items := &IntSlice{1, 2, 3, 4, 5}
			value := 3
			BinarySearch(items, value)
		}
	})
}
