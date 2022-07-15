package main

import (
	"constraints"
	"fmt"
	"os"
)

type Node[T constraints.Ordered] struct {
	prev  *Node[T]
	next  *Node[T]
	value T
}

type OrderedList[T constraints.Ordered] struct {
	head       *Node[T]
	tail       *Node[T]
	_ascending bool
}

func (l *OrderedList[T]) Count() int {
	var node = l.head
	var count = 0
	for {
		if node == nil {
			break
		}
		count++
		node = node.next
	}

	return count
}

func (l *OrderedList[T]) Clear(asc bool) {
	l.head = nil
	l.tail = nil
	l._ascending = asc
}

func (l *OrderedList[T]) Compare(v1 T, v2 T) int {
	if v1 < v2 {
		return -1
	}
	if v1 > v2 {
		return +1
	}
	return 0
}

func (l *OrderedList[T]) Add(item T) {

	var itemInsert = Node[T]{}
	itemInsert.value = item

	if l.head == nil {
		l.head = &itemInsert
		l.tail = &itemInsert
		itemInsert.next = nil
		itemInsert.prev = nil
		return
	}

	//Вставка в начало списка
	if l.Compare(l.head.value, item) > -1 && l._ascending ||
		l.Compare(l.head.value, item) < 1 && !l._ascending {
		itemInsert.next = l.head
		l.head.prev = &itemInsert
		itemInsert.prev = nil
		l.head = &itemInsert
		l.head.prev = nil
		return
	}

	//Вставка в конец списка
	if l.Compare(l.tail.value, item) < 1 && l._ascending ||
		l.Compare(l.tail.value, item) > -1 && !l._ascending {
		l.tail.next = &itemInsert
		itemInsert.prev = l.tail
		l.tail = &itemInsert
		l.tail.next = nil
		return
	}

	var node = l.head
	for {
		if node == nil {
			break
		}
		if l.Compare(node.value, item) > -1 && l._ascending ||
			l.Compare(node.value, item) < 1 && !l._ascending {
			itemInsert.next = node
			itemInsert.prev = node.prev

			node.prev.next = &itemInsert
			node.prev = &itemInsert
			return
		}
		node = node.next
	}
}

func (l *OrderedList[T]) Find(n T) (Node[T], error) {

	var node = l.head
	for {
		if node == nil {
			break
		}
		if l.Compare(node.value, n) == 0 {
			return Node[T]{value: n, next: node.next, prev: node.prev}, nil
		}

		if l.Compare(node.value, n) == 1 && l._ascending ||
			l.Compare(node.value, n) == -1 && !l._ascending {
			break
		}
		node = node.next
	}

	return Node[T]{value: n, next: nil, prev: nil}, fmt.Errorf("bad")
}

func (l *OrderedList[T]) Delete(n T) {

	if l.head == nil {
		return
	}

	if l.head.value == n {
		l.head = l.head.next
		if l.head == nil {
			l.tail = nil
		} else {
			l.head.prev = nil
		}
		return
	} else {
		var currentNode = l.head.next
		var prevNode = l.head

		for {
			if currentNode == nil {
				break
			}

			if currentNode.value == n {
				if l.tail == currentNode {
					l.tail = prevNode
					l.tail.next = nil
					break
				}

				prevNode.next = currentNode.next
				currentNode.next.prev = prevNode
				break
			}

			prevNode = currentNode
			currentNode = currentNode.next
		}
	}
}
