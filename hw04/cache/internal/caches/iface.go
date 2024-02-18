package caches

type Cache interface {
	Get(k string) (string, bool)
	Set(k, v string)
}