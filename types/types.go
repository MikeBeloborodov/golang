package types

import "fmt"

type Node[T comparable] struct {
	next  *Node[T]
	Value T
}

type TLinkedList[T comparable] struct {
	root *Node[T]
}

func (list *TLinkedList[T]) initializeList(values []T) *TLinkedList[T] {
	list.root = new(Node[T])
	node := list.root
	for index, value := range values {
		node.Value = value
		if index < len(values)-1 {
			node.next = new(Node[T])
			node = node.next
		}
	}
	return list
}

func NewLinkedList[T comparable](values []T) *TLinkedList[T] {
	return new(TLinkedList[T]).initializeList(values)
}

func (list *TLinkedList[T]) forEach(callback func(n, prevN *Node[T])) {
	node := list.root
	var prevNode *Node[T]
	for node != nil {
		callback(node, prevNode)
		prevNode = node
		node = node.next
	}
}

func (list *TLinkedList[T]) breakableForEach(callback func(n, prevN *Node[T]) bool) {
	node := list.root
	var prevNode *Node[T]
	for node != nil {
		toBreak := callback(node, prevNode)
		if toBreak {
			break
		}
		prevNode = node
		node = node.next
	}
}

func (list *TLinkedList[T]) SearchAndReplace(target, value T) bool {
	isFound := false
	list.breakableForEach(func(n, prevN *Node[T]) bool {
		if n.Value == target {
			n.Value = value
			isFound = true
			return true
		}
		return false
	})
	return isFound
}

func (list *TLinkedList[T]) Delete(target T) bool {
	isFound := false
	list.breakableForEach(func(n, prevN *Node[T]) bool {
		if n.Value == target {
			prevN.next = n.next
			isFound = true
			return true
		}
		return false
	})
	return isFound
}

func (list *TLinkedList[T]) Add(value T) {
	list.breakableForEach(func(n, prevN *Node[T]) bool {
		if n.next == nil {
			n.next = new(Node[T])
			n.next.Value = value
			return true
		}
		return false
	})
}

func (list TLinkedList[T]) PrintValues() {
	list.forEach(func(n, prevN *Node[T]) { fmt.Println(n.Value) })
}

// type ILinkedList[T any] interface {
// 	OperateOnList(callback func(n *Node[T]))
// 	GetRootNode() Node[T]
// 	PrintValues()
// }
