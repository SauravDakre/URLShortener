package urlstore

import (
	"sync"
)

type URLStore struct {
	urls map[string]string
	mu   sync.RWMutex
}

type KeyGenerator interface {
	GenerateKey(url string) string
}

func NewURLStore() URLStore {
	u := URLStore{
		urls: make(map[string]string),
	}
	return u
}

func (u URLStore) set(key, url string) bool {
	u.mu.Lock()
	defer u.mu.Unlock()
	if _, present := u.urls[key]; present {
		return false
	}
	u.urls[key] = url
	return true
}

func (u URLStore) Get(key string) (url string) {
	u.mu.RLock()
	defer u.mu.RUnlock()
	return u.urls[key]
}

func (u URLStore) Put(url string, keygen KeyGenerator) string {
	for {
		k := keygen.GenerateKey(url)
		if success := u.set(k, url); success {
			return k
		}
	}
}
