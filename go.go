package main

import (
	"fmt"
	"os"
	"reflect"
)

type Node struct {
	next  *Node
	value int
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (l *LinkedList) AddInTail(item Node) {
	if l.head == nil {
		l.head = &item
	} else {
		l.tail.next = &item
	}

	l.tail = &item
}

func (l *LinkedList) Count() int {
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
func (l *LinkedList) Find(n int) (Node, error) {
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

func (l *LinkedList) FindAll(n int) []Node {
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

func (l *LinkedList) Delete(n int, all bool) {

	var foundElement bool = false

	if l.head == nil {
		return
	}

	if l.head.value == n {
		l.head = l.head.next
		if l.head == nil {
			l.tail = nil
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

func (l *LinkedList) Insert(after *Node, add Node) {

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
			return
		}
		node = node.next
	}

}

func (l *LinkedList) InsertFirst(first Node) {

	if l.head == nil {
		l.head = &first
		l.tail = &first
		return
	}

	var oldHead *Node = l.head
	l.head = &first
	l.head.next = oldHead
}

func (l *LinkedList) Clean() {
	l.head = nil
	l.tail = nil
}

func (l *LinkedList) print() {
	var node = l.head

	for {

		if node == nil {
			break
		}
		fmt.Println(node.value)
		node = node.next
	}

}

func linkedListSum(list1 *LinkedList, list2 *LinkedList) *LinkedList {
	var result = new(LinkedList)
	if list1.Count() != list2.Count() {
		return &LinkedList{head: nil, tail: nil}
	}
	var node1 *Node = list1.head
	var node2 *Node = list2.head

	for {

		if node1 == nil {
			break
		}
		result.AddInTail(Node{value: node1.value + node2.value})
		node1 = node1.next
		node2 = node2.next
	}

	return result

}
