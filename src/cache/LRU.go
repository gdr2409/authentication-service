package cache

import "fmt"

type LRU struct {
	hashMap  map[int]*DLLNode
	lruQueue *Queue
	size     int
}

func New(size int) LRU {
	lruQueue := &Queue{
		size:  0,
		front: nil,
		rear:  nil,
	}

	return LRU{
		hashMap:  make(map[int]*DLLNode),
		lruQueue: lruQueue,
		size:     size,
	}
}

func (lru LRU) Get(key int) int {
	node, exists := lru.hashMap[key]

	if exists {
		value := lru.lruQueue.removeNode(node)
		lru.hashMap[key] = lru.lruQueue.add(value)
		fmt.Printf("Inside Get, Size: %v Max Size: %v\n", lru.lruQueue.size, lru.size)
		return value
	}

	return -1
}

func (lru LRU) Put(value int) bool {
	if lru.lruQueue.size == lru.size {
		// Need to drop LRU
		prevValue := lru.lruQueue.remove()
		delete(lru.hashMap, prevValue)
	}

	lru.hashMap[value] = lru.lruQueue.add(value)

	for id, node := range lru.hashMap {
		fmt.Println("Key, value", id, node.value)
	}

	fmt.Printf("Inside Put, Size: %v Max Size: %v\n", lru.lruQueue.size, lru.size)
	return true
}

func (lru LRU) Display() {
	lru.lruQueue.display()
}
