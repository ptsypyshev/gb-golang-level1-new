package memory

import (
	"context"
	"sort"
	"strings"
	"time"

	"github.com/ptsypyshev/gb-golang-level1-new/hw03/urls/internal/app"
	"github.com/ptsypyshev/gb-golang-level1-new/hw03/urls/internal/models"
	"github.com/ptsypyshev/gb-golang-level1-new/hw03/urls/internal/storage"
)

const (
	factorTags = 2 << iota * 2 // factors used to score searched results
	factorDesc
	factorLink
)

var _ app.Storage = (*MemStor)(nil)

// MemStor is an implementation of app.Storage interface that stores data in memory (map).
type MemStor struct {
	Urls map[string]models.URL
}

// New is a constructor for MemStor.
func New() *MemStor {
	return &MemStor{
		Urls: make(map[string]models.URL),
	}
}

// Add implements app.Storage.
func (m *MemStor) Add(args []string) error {
	if len(args) < 3 {
		return storage.ErrBadArguments
	}

	m.Urls[args[0]] = models.URL{
		Date:        time.Now(),
		Link:        args[0],
		Description: args[1],
		Tags:        args[2:],
	}
	return nil
}

// Close implements app.Storage.
func (m *MemStor) Close(ctx context.Context) error {
	return nil
}

// List implements app.Storage.
func (m *MemStor) List() ([]models.URL, error) {
	lst := make([]models.URL, 0, len(m.Urls))
	for _, v := range m.Urls {
		lst = append(lst, v)
	}
	return lst, nil
}

// Remove implements app.Storage.
func (m *MemStor) Remove(url string) error {
	if _, ok := m.Urls[url]; !ok {
		return storage.ErrNotFound
	}
	delete(m.Urls, url)
	return nil
}

// Search implements app.Storage.
func (m *MemStor) Search(t string) ([]models.URL, error) {
	if len(t) < 3 {
		return nil, storage.ErrTooShortSearchWord
	}

	scoredMap := m.scoreUrls(t)
	keys, err := sortedKeys(scoredMap)
	if err != nil {
		return nil, err
	}

	res := make([]models.URL, len(keys))
	for i, k := range keys {
		res[i] = m.Urls[k]
	}
	return res, nil
}

// scoreUrls creates map of url:score.
func (m *MemStor) scoreUrls(t string) map[string]int {
	scoredMap := make(map[string]int)
	var counter int
	for k, v := range m.Urls {
		counter = strings.Count(v.Link, t) * factorLink
		counter += strings.Count(v.Description, t) * factorDesc
		for _, elem := range v.Tags {
			counter += strings.Count(elem, t) * factorTags
		}
		scoredMap[k] = counter
	}
	return scoredMap
}

// sortedKeys returns a sorted slice of scored urls.
func sortedKeys(m map[string]int) ([]string, error) {
	keys := make([]string, 0, len(m))
	for k, v := range m {
		if v > 0 {
			keys = append(keys, k)
		}
	}

	if len(keys) == 0 {
		return nil, storage.ErrSearchNotFound
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})
	return keys, nil
}
