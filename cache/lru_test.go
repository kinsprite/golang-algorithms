package cache

import (
	"reflect"
	"testing"
)

func TestLRUCache_GetEmpty(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name  string
		cache *LRUCache
		args  args
		want  interface{}
		want1 bool
	}{
		{
			name:  "Empty cache get",
			cache: LRUNew(2),
			args:  args{key: "abc"},
			want:  nil,
			want1: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.cache.Get(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LRUCache.Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("LRUCache.Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLRUCache_GetExist(t *testing.T) {
	cache := LRUNew(2)
	cache.Set("abc", 111)
	cache.Set("DEF", 222)
	cache.Set("xyz", 333)

	type args struct {
		key string
	}
	tests := []struct {
		name  string
		cache *LRUCache
		args  args
		want  interface{}
		want1 bool
	}{
		{
			name:  "Exist cache get",
			cache: cache,
			args:  args{key: "DEF"},
			want:  222,
			want1: true,
		},
		{
			name:  "Outdate cache get",
			cache: cache,
			args:  args{key: "abc"},
			want:  nil,
			want1: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.cache.Get(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LRUCache.Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("LRUCache.Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLRUCache_Len(t *testing.T) {
	cache := LRUNew(2)
	cache.Set("abc", 111)
	cache.Set("DEF", 222)
	cache.Set("xyz", 333)

	tests := []struct {
		name  string
		cache *LRUCache
		want  int
	}{
		{
			name:  "cache Len",
			cache: cache,
			want:  2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cache.Len(); got != tt.want {
				t.Errorf("LRUCache.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}
