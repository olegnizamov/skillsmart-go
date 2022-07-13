package main

import (
	"os"
	"fmt"
)

type Queue[T any] struct {
	queue []T
}

func (q *Queue[T]) Size() int {
	return len(q.queue)
}

// Dequeue которая возвращает элемент из головы очереди, удаляя его.
func (q *Queue[T]) Dequeue() (T, error) {
	var result T
	if q.Size() < 1 {
		return result, fmt.Errorf("bad")
	}
	result = q.queue[0]
	copy(q.queue[0:], q.queue[1:])
	q.queue = q.queue[:len(q.queue)-1]
	return result, nil
}

// Enqueue добавить элемент в хвост очереди
func (q *Queue[T]) Enqueue(itm T) {
	q.queue = append(q.queue, itm)
}
