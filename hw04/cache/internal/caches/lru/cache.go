package lru

import (
	"math"

	"github.com/ptsypyshev/gb-golang-level1-new/hw04/cache/internal/caches"
)

const DefaultSize = 5

var _ caches.Cache = (*cacheLRUImpl)(nil)

// Доработает конструктор и методы кеша, так чтобы они соответствовали интерфейсу Cache
func New(size int) *cacheLRUImpl {
	if size == 0 {
		size = DefaultSize
	}
	c := &cacheLRUImpl{
		cache: make(map[string]string),
		size:  size,
		lastHitsCount: make(map[string]int),
	}
	return c
}

type cacheLRUImpl struct {
	cache         map[string]string
	lastHitsCount map[string]int
	size          int
}

func (c *cacheLRUImpl) Get(k string) (string, bool) {
	v, ok := c.cache[k]
	if ok {
		c.lastHitsCount[k] = c.lastHitsCount[k] + 1
	}
	return v, ok
}

func (c *cacheLRUImpl) Set(k, v string) {
	if len(c.cache) == c.size {
		c.dropLeastRecentlyUsed()
	}
	c.cache[k] = v
}

func (c *cacheLRUImpl) dropLeastRecentlyUsed() {
	var leastKey string
	leastCount := math.MaxInt
	for k, v := range c.lastHitsCount {
		if v < leastCount {
			leastKey = k
			leastCount = v
		}
	}
	if leastKey == "" {
		for k := range c.cache {
			leastKey = k // drop any key randomly
			break
		}
	}
	delete(c.cache, leastKey)
}
