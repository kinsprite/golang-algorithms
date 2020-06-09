package cache

import "testing"

func TestLFUCache_Len(t *testing.T) {
	cache := LFUNew(2)
	cache.Set("abc", 111)
	cache.Set("DEF", 222)
	cache.Set("xyz", 333)

	tests := []struct {
		name  string
		cache *LFUCache
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
				t.Errorf("LFUCache.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}
