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
	if cache.contains(newNode) {
		log.Printf("Cache contains: %v", newNode.value)
		cache.moveToHead(newNode)
	} else {
		cache.addToHead(newNode)
	}
	cache.evictOldest()
	return true
}

func (cache *LRU) contains(newNode *Node) bool {
	return cache.find(newNode) != nil
}

func (cache *LRU) find(newNode *Node) *Node {
	current := cache.head
	if current.value == newNode.value {
		return current
	}
	for current.next != nil {
		current = current.next
		if current.value == newNode.value {
			return current
		}
	}
	return nil
}

func (cache *LRU) moveToHead(newNode *Node) {
	// remove node in existing position
	current := cache.head
	if current.value != newNode.value {
		for current.next != nil {
			current = current.next
			if current.value == newNode.value {
				current.prev.next = current.next
			}
		}
	}
	cache.addToHead(newNode)
}

func (cache *LRU) addToHead(newNode *Node) {
	newNode.next = cache.head
	cache.head.prev = newNode
	cache.head = newNode
}

func (cache *LRU) evictOldest() {
	// evict if capacity exceeded
	if cache.capacityExceeded() {
		cache.tail = cache.tail.prev
		cache.tail.next = nil
	}
}

func (cache *LRU) capacityExceeded() bool {
	count := 1
	current := cache.head
	for current.next != nil {
		current = current.next
		count++
	}
	return count > cache.size
}
