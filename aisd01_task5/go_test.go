package main

import (
	"testing"
)

func TestEmpty(t *testing.T) {
	var q = Queue[int]{}
	if q.Size() != 0 {
		t.Error("Error")
	}
}

func TestSizeNotEmpty(t *testing.T) {
	var q = Queue[int]{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	q.Enqueue(6)

	if q.Size() != 6 {
		t.Error("Error")
	}
}
func TestDequeueNotEmpty(t *testing.T) {
	var q = Queue[int]{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	q.Enqueue(6)

	var result, error = q.Dequeue()

	if error != nil {
		t.Error("Error")
	}
	if result != 1 {
		t.Error("Error")
	}

	if q.Size() != 5 {
		t.Error("Error")
	}
}

func TestDequeueEmpty(t *testing.T) {
	var q = Queue[int]{}
	var _, error = q.Dequeue()

	if error == nil {
		t.Error("Error")
	}
}
