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

func TestIn(t *testing.T) {
	var st = Queue[int]{}

	for j := 1; j <= 1000; j++ {
		st.Enqueue(j)
	}

	if st.Size() != 1000 {
		t.Error("Error")
	}

	for j := 1; j <= 1000; j++ {
		var result, error = st.Dequeue()
		if result != j {
			t.Error("Error")
		}
		if error != nil {
			t.Error("Error")
		}
	}

	if st.Size() != 0 {
		t.Error("Error")
	}

	for j := 1; j <= 1000; j++ {
		st.Enqueue(j)
	}

	if st.Size() != 1000 {
		t.Error("Error")
	}

	for j := 1; j <= 1000; j++ {
		var result, error = st.Dequeue()
		if result != j {
			t.Error("Error")
		}
		if error != nil {
			t.Error("Error")
		}
		st.Enqueue(j)
	}

	var result, error = st.Dequeue()

	if result != 1 {
		t.Error("Error")
	}
	if error != nil {
		t.Error("Error")
	}
}
