package main

import (
	"context"
	"fmt"

	"github.com/ptsypyshev/gb-golang-level1-new/hw04/cache/internal/caches/exp"
	"github.com/ptsypyshev/gb-golang-level1-new/hw04/cache/internal/caches/lru"
	"github.com/ptsypyshev/gb-golang-level1-new/hw04/cache/internal/db"
)

func main() {
	c := exp.New(context.Background(), exp.DefaultTTL)
	db1 := db.New(c)

	fmt.Println(db1.Get("test"))  // read from dbs and cache result
	fmt.Println(db1.Get("hello")) // read from dbs and cache result

	fmt.Println(db1.Get("test"))  // read from cache
	fmt.Println(db1.Get("hello")) // read from cache

	db1.Set("new", "value")    // set value to dbs and cache
	db1.Set("another", "item") // set value to dbs and cache

	fmt.Println(db1.Get("new"))     // read from cache
	fmt.Println(db1.Get("another")) // read from cache

	lruCache := lru.New(3)
	db2 := db.New(lruCache)

	db2.Set("first", "one")
	db2.Set("second", "two")
	db2.Set("third", "three")

	db2.Get("second")
	fmt.Println(db2.Get("first"))
	fmt.Println(db2.Get("second"))
	fmt.Println(db2.Get("third"))
	db2.Get("second")

	db2.Set("fourth", "four")
	db2.Set("fifth", "five")

	fmt.Println(db2.Get("first")) // read from dbs
	fmt.Println(db2.Get("second")) // read from cache
	fmt.Println(db2.Get("third")) // read from dbs
	fmt.Println(db2.Get("fourth")) // read from cache
	fmt.Println(db2.Get("fifth")) // read from cache
}
