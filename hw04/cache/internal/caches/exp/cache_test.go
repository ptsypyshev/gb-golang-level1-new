package exp

import (
	"context"
	"testing"
	"time"

	"github.com/ptsypyshev/gb-golang-level1-new/hw04/cache/internal/llist"
	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	got := New(context.Background(), DefaultTTL)
	assert.NotNil(t, got.cache)
	assert.IsType(t, map[string]string{}, got.cache)
}

func Test_cacheImpl_Get(t *testing.T) {
	t.Parallel()
	cache := &cacheImpl{
		cache: map[string]string{"exist": "true"},
		timeList: &llist.LinkedListImpl{},
		ttl: DefaultTTL,
	}

	type args struct {
		key   string
		value string
		ok    bool
	}

	tests := []struct {
		name string
		c    *cacheImpl
		args args
	}{
		{
			name: "Get exist value from cache",
			c:    cache,
			args: args{
				key:   "exist",
				value: "true",
				ok:    true,
			},
		},
		{
			name: "Get non-exist value from cache",
			c:    cache,
			args: args{
				key:   "not_exist",
				value: "",
				ok:    false,
			},
		},
		{
			name: "Get any value from empty cache",
			c:    New(context.Background(), DefaultTTL),
			args: args{
				key:   "any",
				value: "",
				ok:    false,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, got1 := tt.c.Get(tt.args.key)
			assert.Equal(t, tt.args.value, got)
			assert.Equal(t, tt.args.ok, got1)
		})
	}
}

func Test_cacheImpl_Set(t *testing.T) {
	cache := &cacheImpl{
		cache: map[string]string{"exist": "true"},
		timeList: &llist.LinkedListImpl{},
		ttl: DefaultTTL,
	}

	type args struct {
		k string
		v string
	}
	tests := []struct {
		name string
		c    *cacheImpl
		args args
	}{
		{
			name: "Set to empty cache",
			c:    New(context.Background(), DefaultTTL),
			args: args{
				k: "first",
				v: "item",
			},
		},
		{
			name: "Set to filled cache with same key",
			c:    cache,
			args: args{
				k: "exist",
				v: "changed",
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.c.Set(tt.args.k, tt.args.v)
			v, ok := tt.c.cache[tt.args.k]
			assert.Equal(t, tt.args.v, v)
			assert.Equal(t, true, ok)
		})
	}
}

func Test_cacheImpl_dropExpired(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	c := New(ctx, 30 * time.Millisecond)

	c.Set("first", "one")
	time.Sleep(20 * time.Millisecond) // 20ms
	c.Set("second", "two")
	time.Sleep(20 * time.Millisecond) // 40ms
	c.Set("third", "three")
	time.Sleep(20 * time.Millisecond) // 60ms
	c.Set("fourh", "four")
	time.Sleep(20 * time.Millisecond) // 80ms

	assert.Equal(t, map[string]string{"third": "three", "fourh": "four"}, c.cache)
}
