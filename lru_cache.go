package lru_cache

type Node struct {
	prev  *Node
	next  *Node
	value string
}

type LRU struct {
	size int
	head *Node
	tail *Node
}

type LRUCache interface {
	Size() int
	Cache() []string
	Add(value string) bool
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
	cacheContents := make([]string, 0)
	current := cache.head
	if current == nil {
		return []string{}
	} else {
		cacheContents = append(cacheContents, current.value)
	}
	for current.next != nil {
		current = current.next
		cacheContents = append(cacheContents, current.value)
	}
	return cacheContents
}

func (cache *LRU) Add(value string) bool {
	newNode := &Node{value: value}
	if cache.head == nil {
		cache.head = newNode
		cache.tail = newNode
		return true
	}
	// insert at the head
	newNode.next = cache.head
	cache.head.prev = newNode
	cache.head = newNode

	// evict if capacity exceeded
	count := 1
	current := cache.head
	for current.next != nil {
		current = current.next
		count++
	}
	if count > cache.size {
		cache.tail = cache.tail.prev
		cache.tail.next = nil
	}
	return true
}
