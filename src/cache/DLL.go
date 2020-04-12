package cache

import "fmt"

type DLLNode struct {
	key   uint
	value string
	next  *DLLNode
	prev  *DLLNode
}

type Queue struct {
	front *DLLNode
	rear  *DLLNode
	size  int
}

func (Q *Queue) add(id uint, val string) *DLLNode {
	temp := &DLLNode{
		key:   id,
		value: val,
		next:  nil,
		prev:  nil,
	}

	if Q.front == nil {
		Q.front = temp
		Q.rear = temp
	} else {
		temp.next = Q.front
		Q.front.prev = temp
		Q.front = temp
	}

	Q.size++
	return temp
}

func (Q *Queue) remove() (uint, string) {
	if Q.rear == nil {
		return 0, ""
	}

	value := Q.rear.value
	key := Q.rear.key

	if Q.rear == Q.front {
		Q.front = nil
		Q.rear = nil
	} else {
		temp := Q.rear.prev
		temp.next = nil
		Q.rear = temp
	}

	Q.size--
	// fmt.Printf("Inside remove, Size: %v\n", Q.size)
	return key, value
}

func (Q *Queue) removeNode(node *DLLNode) (uint, string) {
	if node == nil {
		return 0, ""
	}

	if node == Q.rear {
		return Q.remove()
	}

	if node == Q.front {
		node.next.prev = node.prev
		Q.front = node.next
		Q.size--
		return node.key, node.value
	}

	value := node.value
	key := node.key

	node.next.prev = node.prev
	node.prev.next = node.next

	Q.size--
	// fmt.Printf("Inside remove, Size: %v\n", Q.size)
	return key, value
}

func (Q *Queue) display() {
	temp := Q.front

	for temp != nil {
		fmt.Printf("--> %v ", temp.value)
		temp = temp.next
	}
}
