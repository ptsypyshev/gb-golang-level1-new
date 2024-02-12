package memory

import (
	"context"
	"time"

	"github.com/ptsypyshev/gb-golang-level1-new/hw03/urls/internal/app"
	"github.com/ptsypyshev/gb-golang-level1-new/hw03/urls/internal/models"
	"github.com/ptsypyshev/gb-golang-level1-new/hw03/urls/internal/storage"
)

var _ app.Storage = (*MemStor)(nil)

// MemStor is an implementation of app.Storage interface that stores data in memory (map)
type MemStor struct {
	urls map[string]models.URL
}

// New is a constructor for MemStor
func New() *MemStor {
	return &MemStor{
		urls: make(map[string]models.URL),
	}
}

// Add implements app.Storage.
func (a *MemStor) Add(args []string) error {
	if len(args) < 3 {
		return storage.ErrBadArguments
	}

	a.urls[args[0]] = models.URL{
		Date:        time.Now(),
		Link:        args[0],
		Description: args[1],
		Tags:        args[2:],
	}
	return nil
}

// Close implements app.Storage.
func (a *MemStor) Close(ctx context.Context) error {
	return nil
}

// List implements app.Storage.
func (a *MemStor) List() ([]models.URL, error) {
	lst := make([]models.URL, 0, len(a.urls))
	for _, v := range a.urls {
		lst = append(lst, v)
	}
	return lst, nil
}

// Remove implements app.Storage.
func (a *MemStor) Remove(url string) error {
	if _, ok := a.urls[url]; !ok {
		return storage.ErrNotFound
	}
	delete(a.urls, url)
	return nil
}
