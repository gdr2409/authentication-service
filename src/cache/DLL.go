package cache

import "fmt"

type DLLNode struct {
	value int
	next  *DLLNode
	prev  *DLLNode
}

type Queue struct {
	front *DLLNode
	rear  *DLLNode
	size  int
}

func (Q *Queue) add(val int) *DLLNode {
	temp := &DLLNode{
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

func (Q *Queue) remove() int {
	if Q.rear == nil {
		return -1
	}

	value := Q.rear.value

	if Q.rear == Q.front {
		Q.front = nil
		Q.rear = nil
	} else {
		temp := Q.rear.prev
		temp.next = nil
		Q.rear = temp
	}

	Q.size--
	fmt.Printf("Inside remove, Size: %v\n", Q.size)
	return value
}

func (Q *Queue) removeNode(node *DLLNode) int {
	if node == nil {
		return -1
	}

	if node == Q.front {
		node.next.prev = node.prev
		Q.front = node.next
		Q.size--
		return node.value
	}

	if node == Q.rear {
		return Q.remove()
	}

	value := node.value

	node.next.prev = node.prev
	node.prev.next = node.next

	Q.size--
	fmt.Printf("Inside remove, Size: %v\n", Q.size)
	return value
}

func (Q *Queue) display() {
	temp := Q.front

	for temp != nil {
		fmt.Printf("--> %v ", temp.value)
		temp = temp.next
	}
}
