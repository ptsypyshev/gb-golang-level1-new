package lru

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		size int
		want *cacheLRUImpl
	}{
		{
			name: "LRU Cache with size 4",
			size: 4,
			want: &cacheLRUImpl{
				cache: make(map[string]string),
				size:  4,
			},
		},
		{
			name: "LRU Cache with size 10",
			size: 4,
			want: &cacheLRUImpl{
				cache: make(map[string]string),
				size:  10,
			},
		},
		{
			name: "LRU Cache with size 0",
			size: 0,
			want: &cacheLRUImpl{
				cache: make(map[string]string),
				size:  DefaultSize,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := New(tt.size)
			assert.Equal(t, map[string]string{}, c.cache)
			assert.Equal(t, c.size, c.size)
		})
	}
}

func Test_cacheLRUImpl_Get(t *testing.T) {
	t.Parallel()

	cache := &cacheLRUImpl{
		cache:         map[string]string{"exist": "true"},
		lastHitsCount: make(map[string]int),
		size:          2,
	}

	tests := []struct {
		name string
		c    *cacheLRUImpl
		k    string
		v    string
		ok   bool
	}{
		{
			name: "Get existing key from cache",
			c:    cache,
			k:    "exist",
			v:    "true",
			ok:   true,
		},
		{
			name: "Get not existing key from cache",
			c:    cache,
			k:    "next",
			v:    "",
			ok:   false,
		},
		{
			name: "Get any key from empty cache",
			c:    New(DefaultSize),
			k:    "any",
			v:    "",
			ok:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, ok := tt.c.Get(tt.k)
			assert.Equal(t, tt.v, v)
			assert.Equal(t, tt.ok, ok)
		})
	}
}

func Test_cacheLRUImpl_Set(t *testing.T) {
	t.Parallel()

	c := New(2)
	c.Set("First", "One")
	assert.Equal(t, map[string]string{"First": "One"}, c.cache)

	c.Set("Second", "Two")
	assert.Equal(t, map[string]string{"First": "One", "Second": "Two"}, c.cache)

	c.Set("Third", "Three")
	assert.Equal(t, 2, len(c.cache)) // Cannot check values because one of previous will be dropped randomly

	_, _ = c.Get("First")
	_, _ = c.Get("Third")
	_, _ = c.Get("Second")
	_, _ = c.Get("Third")

	c.Set("Fourth", "Four")
	assert.Equal(t, map[string]string{"Third": "Three", "Fourth": "Four"}, c.cache) // Third is more frequently used than First or Second
}
