package db

import (
	"context"
	"testing"

	"github.com/ptsypyshev/gb-golang-level1-new/hw04/cache/internal/cache"
	"github.com/stretchr/testify/assert"
)

func Test_dbImpl_Get(t *testing.T) {
	t.Parallel()

	c := cache.New(context.Background(), cache.DefaultTTL)
	db := New(c)
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
