package singleton

import (
	"fmt"
)

type cacheStore struct{}

var store *cacheStore

func init() {
	fmt.Println("initialize cache store")
	store = new(cacheStore)
}

func GetCacheStore() *cacheStore {
	return store
}
