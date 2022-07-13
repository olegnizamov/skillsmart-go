package main

import (
	"fmt"
	"testing"
)

func TestEmpty(t *testing.T) {
	var d = Deque[int]{}
	if d.Size() != 0 {
		t.Error("Error")
	}
}

func TestSizeNotEmpty(t *testing.T) {
	var d = Deque[int]{}
	d.AddTail(1)
	d.AddTail(2)
	d.AddTail(3)
	d.AddTail(4)
	d.AddTail(5)
	d.AddTail(6)

	if d.Size() != 6 {
		t.Error("Error")
	}
}
func TestAddFront(t *testing.T) {
	var d = Deque[int]{}
	d.AddFront(1)
	d.AddFront(2)

	if d.dequeue[0] != 2 {
		t.Error("Error")
	}

	if d.dequeue[1] != 1 {
		t.Error("Error")
	}
}

func TestAddTail(t *testing.T) {
	var d = Deque[int]{}
	d.AddTail(1)
	d.AddTail(2)

	if d.dequeue[0] != 1 {
		t.Error("Error")
	}

	if d.dequeue[1] != 2 {
		t.Error("Error")
	}
}

func TestRemoveFront(t *testing.T) {
	var d = Deque[int]{}
	d.AddTail(1)
	d.AddTail(2)

	var result, error = d.RemoveFront()

	if result != 1 {
		t.Error("Error")
	}

	if d.dequeue[0] != 2 {
		t.Error("Error")
	}

	if error != nil {
		t.Error("Error")
	}

}

func TestRemoveFrontEmpty(t *testing.T) {
	var d = Deque[int]{}

	var _, error = d.RemoveFront()

	if error == nil {
		t.Error("Error")
	}

}

func TestRemoveTailEmpty(t *testing.T) {
	var d = Deque[int]{}

	var _, error = d.RemoveTail()

	if error == nil {
		t.Error("Error")
	}

}

func TestRemoveTail(t *testing.T) {
	var d = Deque[int]{}
	d.AddTail(1)
	d.AddTail(2)

	var result, error = d.RemoveTail()

	if result != 2 {
		t.Error("Error")
	}

	if d.dequeue[0] != 1 {
		t.Error("Error")
	}

	if error != nil {
		t.Error("Error")
	}
}

func TestPolindrome(t *testing.T) {
	var is_string_polindrome = true
	var deque = Deque[string]{}
	var checkString = "aanabnaa"

	for j := 0; j < len(checkString); j++ {
		deque.AddTail(string(checkString[j]))
	}

	for {

		if deque.Size() <= 1 {
			break
		}
		var front, _ = deque.RemoveFront()
		var back, _ = deque.RemoveTail()

		if front != back {
			is_string_polindrome = false
			break
		}
	}

	fmt.Println(checkString, " = ", is_string_polindrome)
}
