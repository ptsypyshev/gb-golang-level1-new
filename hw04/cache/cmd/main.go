package main

import (
	"fmt"
)

type Cache interface {
	Get(k string) (string, bool)
	Set(k, v string)
}

var _ Cache = (*cacheImpl)(nil)

// Доработает конструктор и методы кеша, так чтобы они соответствовали интерфейсу Cache
func newCacheImpl() *cacheImpl {
	return &cacheImpl{cache: make(map[string]string)}
}

type cacheImpl struct {
	cache map[string]string
}

func (c *cacheImpl) Get(k string) (string, bool) {
	v, ok := c.cache[k]
	return v, ok
}

func (c *cacheImpl) Set(k, v string) {
	c.cache[k] = v
}

func newDbImpl(cache Cache) *dbImpl {
	return &dbImpl{cache: cache, dbs: map[string]string{"hello": "world", "test": "test"}}
}

type dbImpl struct {
	cache Cache
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

func main() {
	c := newCacheImpl()
	db := newDbImpl(c)

	fmt.Println(db.Get("test"))  // read from dbs and cache result
	fmt.Println(db.Get("hello")) // read from dbs and cache result

	fmt.Println(db.Get("test"))  // read from cache
	fmt.Println(db.Get("hello")) // read from cache

	db.Set("new", "value")    // set value to dbs and cache
	db.Set("another", "item") // set value to dbs and cache

	fmt.Println(db.Get("new"))     // read from cache
	fmt.Println(db.Get("another")) // read from cache
}
