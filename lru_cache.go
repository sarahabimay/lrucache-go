package lru_cache

type Node struct {
	prev  *Node
	next  *Node
	value string
}

type LRU struct {
	size int
	head *Node
}

type LRUCache interface {
	Size() int
	Cache() []string
}

func NewLRU(cacheSize int) *LRU {
	return &LRU{
		size: cacheSize,
		head: nil,
	}
}

func (cache *LRU) Size() int {
	return cache.size
}

func (cache *LRU) Cache() []string {
	current := cache.head
	if current == nil {
		return []string{}
	}
	cacheContents := make([]string, cache.size)
	for current.next != nil {
		cacheContents = append(cacheContents, current.value)
		current = current.next
	}
	return cacheContents
}
