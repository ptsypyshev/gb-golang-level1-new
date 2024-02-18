package db

import (
	"fmt"

	"github.com/ptsypyshev/gb-golang-level1-new/hw04/cache/internal/cache"
)

func New(cache cache.Cache) *dbImpl {
	return &dbImpl{cache: cache, dbs: map[string]string{"hello": "world", "test": "test"}}
}

type dbImpl struct {
	cache cache.Cache
	dbs   map[string]string
}

func (d *dbImpl) Get(k string) (string, bool) {
	v, ok := d.cache.Get(k)
	if ok {
		return fmt.Sprintf("answer from cache: key: %s, val: %s", k, v), ok
	}

	v, ok = d.dbs[k]
	d.cache.Set(k, v)
	return fmt.Sprintf("answer from dbs: key: %s, val: %s", k, v), ok
}

func (d *dbImpl) Set(k, v string) {
	d.dbs[k] = v
	d.cache.Set(k, v)
	fmt.Printf("value is set up to dbs and cache: key: %s, val: %s\n", k, v)
}
