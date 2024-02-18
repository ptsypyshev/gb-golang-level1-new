package exp

import (
	"context"
	"time"

	"github.com/ptsypyshev/gb-golang-level1-new/hw04/cache/internal/caches"
	"github.com/ptsypyshev/gb-golang-level1-new/hw04/cache/internal/llist"
)

const DefaultTTL = 5 * time.Second

var _ caches.Cache = (*cacheImpl)(nil)

// Доработает конструктор и методы кеша, так чтобы они соответствовали интерфейсу Cache
func New(ctx context.Context, ttl time.Duration) *cacheImpl {
	c := &cacheImpl{
		cache:    make(map[string]string),
		ttl:      ttl,
		timeList: &llist.LinkedListImpl{},
	}
	go c.dropExpired(ctx)
	return c
}

type cacheImpl struct {
	cache    map[string]string
	timeList llist.LinkedList
	ttl      time.Duration
}

func (c *cacheImpl) Get(k string) (string, bool) {
	v, ok := c.cache[k]
	return v, ok
}

func (c *cacheImpl) Set(k, v string) {
	c.cache[k] = v
	n := llist.NewNode(time.Now(), k)
	c.timeList.Append(n)
}

func (c *cacheImpl) dropExpired(ctx context.Context) {
	t := time.NewTicker(c.ttl)
	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			oldest := c.timeList.GetHead()
			for oldest != nil && time.Since(oldest.Updated()) > c.ttl {
				c.timeList.PopHead()
				delete(c.cache, oldest.Key())
				oldest = c.timeList.GetHead()
			}
		}
	}
}
