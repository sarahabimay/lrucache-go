package lru_cache_test

import (
	"lru_cache"
	"reflect"
	"testing"
)

func TestLRUCacheCreateWithLength(t *testing.T) {
	lru := lru_cache.NewLRU(2)
	if lru.Size() == 2 {
		t.Logf("Success")
	} else {
		t.Errorf("Failure")
	}
}

func TestLRUCacheCreateCacheWithLength(t *testing.T) {
	lru := lru_cache.NewLRU(2)
	if reflect.DeepEqual(lru.Cache(), []string{}) {
		t.Logf("Success")
	} else {
		t.Errorf("Failure")
	}
}
