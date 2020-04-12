package cache

import "fmt"

type LRU struct {
	hashMap  map[uint]*DLLNode
	lruQueue *Queue
	size     int
}

func New(size int) *LRU {
	lruQueue := &Queue{
		size:  0,
		front: nil,
		rear:  nil,
	}

	return &LRU{
		hashMap:  make(map[uint]*DLLNode),
		lruQueue: lruQueue,
		size:     size,
	}
}

func (lru LRU) Get(key uint) string {
	node, exists := lru.hashMap[key]

	if exists {
		key, value := lru.lruQueue.removeNode(node)
		lru.hashMap[key] = lru.lruQueue.add(key, value)
		fmt.Printf("Inside Get, Size: %v Max Size: %v\n", lru.lruQueue.size, lru.size)
		return value
	}

	return ""
}

func (lru LRU) Put(id uint, value string) bool {
	if lru.lruQueue.size == lru.size {
		// Need to drop LRU
		key, _ := lru.lruQueue.remove()
		delete(lru.hashMap, key)
	}

	lru.hashMap[id] = lru.lruQueue.add(id, value)

	fmt.Printf("Inside Put, Size: %v Max Size: %v\n", lru.lruQueue.size, lru.size)
	return true
}

func (lru LRU) Display() {
	lru.lruQueue.display()
}
