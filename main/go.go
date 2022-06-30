package main

import (
	"fmt"
)

type Node struct {
	prev  *Node
	next  *Node
	value int
}

type LinkedList2 struct {
	head *Node
	tail *Node
}

func (l *LinkedList2) AddInTail(item Node) {
	if l.head == nil {
		l.head = &item
		l.head.next = nil
		l.head.prev = nil
	} else {
		l.tail.next = &item
		item.prev = l.tail
	}

	l.tail = &item
	l.tail.next = nil
}

func (l *LinkedList2) Count() int {
	var node = l.head
	var count int = 0
	for {
		if node == nil {
			break
		}
		count++
		node = node.next
	}

	return count
}

// error не nil, если узел не найден
func (l *LinkedList2) Find(n int) (Node, error) {
	var node = l.head

	for {

		if node == nil {
			break
		}

		if node.value == n {
			return Node{value: node.value, next: node.next}, nil
		}

		node = node.next
	}

	return Node{value: -1, next: nil}, fmt.Errorf("not found")
}

func (l *LinkedList2) FindAll(n int) []Node {

	var nodes []Node
	var node = l.head
	for {

		if node == nil {
			break
		}

		if node.value == n {
			nodes = append(nodes, Node{value: node.value, next: node.next})
		}

		node = node.next
	}

	return nodes
}

func (l *LinkedList2) Delete(n int, all bool) {
	var foundElement bool = false

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
		foundElement = true
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
					foundElement = true
					break
				}

				prevNode.next = currentNode.next
				currentNode.next.prev = prevNode
				foundElement = true
				break
			}

			prevNode = currentNode
			currentNode = currentNode.next
		}
	}

	if all == true && foundElement == true {
		l.Delete(n, all)
	}
}

func (l *LinkedList2) Insert(after *Node, add Node) {

	if after == nil {
		l.InsertFirst(add)
		return
	}

	var node = l.head
	for {

		if node == nil {
			break
		}

		if node.value == after.value {
			if node == l.tail {
				l.tail = &add
			}
			add.next = node.next
			node.next = &add
			add.prev = node
			return
		}
		node = node.next
	}

}

func (l *LinkedList2) InsertFirst(first Node) {

	if l.head == nil {
		l.head = &first
		l.tail = &first
		first.next = nil
		first.prev = nil
		return
	}

	var oldHead *Node = l.head
	oldHead.prev = &first
	l.head = &first
	l.head.next = oldHead
}

func (l *LinkedList2) Clean() {
	l.head = nil
	l.tail = nil
}

func (l *LinkedList2) print() {
	var node = l.head

	for {

		if node == nil {
			break
		}
		fmt.Println(node.value)
		node = node.next
	}

}
