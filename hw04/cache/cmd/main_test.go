package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newCacheImpl(t *testing.T) {
	got := newCacheImpl()
	assert.NotNil(t, got.cache)
	assert.IsType(t, map[string]string{}, got.cache)
}

func Test_cacheImpl_Get(t *testing.T) {
	t.Parallel()
	cache := &cacheImpl{
		cache: map[string]string{"exist": "true"},
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
			c:    newCacheImpl(),
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
			c:    newCacheImpl(),
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

func Test_dbImpl_Get(t *testing.T) {
	t.Parallel()

	c := newCacheImpl()
	db := newDbImpl(c)
	db.Set("exist", "true")

	tests := []struct {
		name  string
		d     *dbImpl
		key   string
		value string
		ok    bool
	}{
		{
			name:  "Get exist value from db (cached)",
			d:     db,
			key:   "exist",
			value: "answer from cache: key: exist, val: true",
			ok:    true,
		},
		{
			name:  "Get exist value from db (not cached)",
			d:     db,
			key:   "hello",
			value: "answer from dbs: key: hello, val: world",
			ok:    true,
		},		
		{
			name:  "Get exist value from db (now it is cached)",
			d:     db,
			key:   "hello",
			value: "answer from cache: key: hello, val: world",
			ok:    true,
		},
		{
			name:  "Get exist value from db (not exist)",
			d:     db,
			key:   "not_exist",
			value: "answer from dbs: key: not_exist, val: ",
			ok:    false,
		},
		{
			name:  "Get exist value from db (now it is cached but absent in dbs)",
			d:     db,
			key:   "not_exist",
			value: "answer from cache: key: not_exist, val: ",
			ok:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			val, ok := tt.d.Get(tt.key)

			assert.Equal(t, tt.value, val)
			assert.Equal(t, tt.ok, ok)
		})
	}
}
