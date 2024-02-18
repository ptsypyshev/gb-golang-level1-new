package main

import (
	"context"
	"fmt"

	"github.com/ptsypyshev/gb-golang-level1-new/hw04/cache/internal/cache"
	"github.com/ptsypyshev/gb-golang-level1-new/hw04/cache/internal/db"
)

func main() {
	c := cache.New(context.Background(), cache.DefaultTTL)
	db := db.New(c)

	fmt.Println(db.Get("test"))  // read from dbs and cache result
	fmt.Println(db.Get("hello")) // read from dbs and cache result

	fmt.Println(db.Get("test"))  // read from cache
	fmt.Println(db.Get("hello")) // read from cache

	db.Set("new", "value")    // set value to dbs and cache
	db.Set("another", "item") // set value to dbs and cache

	fmt.Println(db.Get("new"))     // read from cache
	fmt.Println(db.Get("another")) // read from cache
}
