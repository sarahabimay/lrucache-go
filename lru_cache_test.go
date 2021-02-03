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

func TestLRUCacheAddToEmptyCache(t *testing.T) {
	lru := lru_cache.NewLRU(2)
	ok := lru.Add(1, "Menu1")

	if !ok {
		t.Errorf("Add was not successful")
	}
	cacheContents := lru.Cache()
	switch {
	case len(cacheContents) != 1:
		t.Errorf("Failure - Expect cache length 1, got %v", len(cacheContents))
	case !reflect.DeepEqual(cacheContents, []string{"Menu1"}):
		t.Errorf("Failure, Cache contents: %v does not equal: %v", cacheContents, []string{"Menu1"})
	default:
		t.Logf("Success")
	}
}

func TestLRUCacheAddTwoMenus(t *testing.T) {
	lru := lru_cache.NewLRU(2)

	ok := lru.Add(1, "Menu1")
	if !ok {
		t.Errorf("Add menu: Menu1 was not successful")
	}
	ok = lru.Add(2, "Menu2")
	if !ok {
		t.Errorf("Add menu: %v was not successful", "Menu2")
	}

	cacheContents := lru.Cache()
	switch {
	case len(cacheContents) != 2:
		t.Errorf("Failure - Expect cache length 2, got %v", len(cacheContents))
	case !reflect.DeepEqual(cacheContents, []string{"Menu2", "Menu1"}):
		t.Errorf("Failure, Cache contents: %v does not equal: %v", cacheContents, []string{"Menu2", "Menu1"})
	default:
		t.Logf("Success")
	}
}

func TestLRUCacheEvictLRUMenuWhenCapacityExceeded(t *testing.T) {
	lru := lru_cache.NewLRU(2)

	ok := lru.Add(1, "Menu1")
	if !ok {
		t.Errorf("Add menu: Menu1 was not successful")
	}
	ok = lru.Add(2, "Menu2")
	if !ok {
		t.Errorf("Add menu: %v was not successful", "Menu2")
	}
	ok = lru.Add(3, "Menu3")
	if !ok {
		t.Errorf("Add menu: %v was not successful", "Menu3")
	}

	cacheContents := lru.Cache()
	switch {
	case len(cacheContents) != 2:
		t.Errorf("Failure - Expect cache length 2, got %v", len(cacheContents))
	case !reflect.DeepEqual(cacheContents, []string{"Menu3", "Menu2"}):
		t.Errorf("Failure, Cache contents: %v does not equal: %v", cacheContents, []string{"Menu3", "Menu2"})
	default:
		t.Logf("Success, cache contents: %v", cacheContents)
	}
}

func TestLRUCacheRequestInsertSameMenu(t *testing.T) {
	lru := lru_cache.NewLRU(3)

	ok := lru.Add(1, "Menu1")
	if !ok {
		t.Errorf("Add menu: Menu1 was not successful")
	}
	ok = lru.Add(2, "Menu2")
	if !ok {
		t.Errorf("Add menu: %v was not successful", "Menu2")
	}
	ok = lru.Add(3, "Menu1")
	if !ok {
		t.Errorf("Add menu: %v was not successful", "Menu1")
	}

	cacheContents := lru.Cache()
	switch {
	case len(cacheContents) != 2:
		t.Errorf("Failure - Expect cache length 2, got %v", len(cacheContents))
	case !reflect.DeepEqual(cacheContents, []string{"Menu1", "Menu2"}):
		t.Errorf("Failure, Cache contents: %v does not equal: %v", cacheContents, []string{"Menu1", "Menu2"})
	default:
		t.Logf("Success")
	}
}

func TestLRUCacheRequestNonexistantMenu(t *testing.T) {
	lru := lru_cache.NewLRU(3)

	item, ok := lru.Get(1)
	if !ok {
		t.Logf("Success")
	} else {
		t.Errorf("Failure, expected nothing to be found but got: %v", item)
	}
}

func TestLRUCacheRequestAMenu(t *testing.T) {
	lru := lru_cache.NewLRU(2)

	lru.Add(1, "Menu1")
	item, ok := lru.Get(1)
	if ok {
		t.Logf("Success")
	} else {
		t.Errorf("Failure, expected %v to be found but found: %v", "Menu2", item)
	}
}

func TestLRUCacheGetItemMostRecentlyUsed(t *testing.T) {
	lru := lru_cache.NewLRU(2)

	lru.Add(1, "Menu1")
	lru.Add(2, "Menu2")
	cacheContentsBeforeGet := lru.Cache()
	if !reflect.DeepEqual(cacheContentsBeforeGet, []string{"Menu2", "Menu1"}) {
		t.Errorf("Failure, Cache contents: %v does not equal: %v", cacheContentsBeforeGet, []string{"Menu2", "Menu1"})
	}

	item, ok := lru.Get(1)
	if ok {
		t.Logf("Success")
	} else {
		t.Errorf("Failure, expected %v to be found but found: %v", "Menu2", item)
	}
	cacheContentsAfterGet := lru.Cache()
	if !reflect.DeepEqual(cacheContentsAfterGet, []string{"Menu1", "Menu2"}) {
		t.Errorf("Failure, Cache contents: %v does not equal: %v", cacheContentsAfterGet, []string{"Menu1", "Menu2"})
	}
}
