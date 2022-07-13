package main

import (
	//"os"
	"fmt"
)

type Deque[T any] struct {
	dequeue []T
}

func (d *Deque[T]) Size() int {
	return len(d.dequeue)
}

func (d *Deque[T]) AddFront(itm T) {
	var netQ []T
	netQ = append(netQ, itm)
	var result = append(netQ, d.dequeue...)
	d.dequeue = result
}

func (d *Deque[T]) AddTail(itm T) {
	d.dequeue = append(d.dequeue, itm)
}

func (d *Deque[T]) RemoveFront() (T, error) {
	var result T
	if d.Size() < 1 {
		return result, fmt.Errorf("bad")
	}
	result = d.dequeue[0]
	copy(d.dequeue[0:], d.dequeue[1:])
	d.dequeue = d.dequeue[:len(d.dequeue)-1]
	return result, nil
}

func (d *Deque[T]) RemoveTail() (T, error) {
	var result T
	if d.Size() < 1 {
		return result, fmt.Errorf("bad")
	}
	result = d.dequeue[len(d.dequeue)-1]
	d.dequeue = d.dequeue[:len(d.dequeue)-1]
	return result, nil
}
