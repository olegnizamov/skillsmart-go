package main

import (
	"testing"
)

func TestCountEmpty(t *testing.T) {
	var ol = OrderedList[int]{}
	ol._ascending = true
	if ol.Count() != 0 {
		t.Error("Error")
	}
}

func TestCountNotEmpty(t *testing.T) {
	var ol = OrderedList[int]{}
	ol._ascending = false
	ol.Add(1)
	ol.Add(2)
	ol.Add(3)

	if ol.Count() != 3 {
		t.Error("Error")
	}
}

func TestClearEmpty(t *testing.T) {
	var ol = OrderedList[int]{}
	ol._ascending = true
	ol.Clear(false)
	if ol.Count() != 0 {
		t.Error("Error")
	}
	if ol._ascending != false {
		t.Error("Error")
	}

}

func TestClearNotEmpty(t *testing.T) {
	var ol = OrderedList[int]{}
	ol._ascending = true
	ol.Add(1)
	ol.Add(2)
	ol.Add(3)
	ol.Clear(false)

	if ol.Count() != 0 {
		t.Error("Error")
	}
	if ol._ascending != false {
		t.Error("Error")
	}
}

func TestAddInHead(t *testing.T) {
	var ol = OrderedList[int]{}
	ol._ascending = true
	ol.Add(1)

	if ol.Count() != 1 {
		t.Error("Error")
	}
	if ol.head.value != 1 {
		t.Error("Error")
	}

	if ol.tail.value != 1 {
		t.Error("Error")
	}
}

func TestDeleteTail(t *testing.T) {
	var ol = OrderedList[int]{}
	ol._ascending = true
	ol.Add(1)
	ol.Add(2)
	ol.Add(3)
	ol.Delete(3)

	if ol.Count() != 2 {
		t.Error("Error")
	}

	if ol.head.value != 1 {
		t.Error("Error")
	}

	if ol.tail.value != 2 {
		t.Error("Error")
	}
}

func TestDeleteHead1(t *testing.T) {
	var ol = OrderedList[int]{}
	ol._ascending = true
	ol.Add(1)
	ol.Add(2)
	ol.Add(3)
	ol.Delete(1)

	if ol.Count() != 2 {
		t.Error("Error")
	}

	if ol.head.value != 2 {
		t.Error("Error")
	}
}

func TestDeleteHead(t *testing.T) {
	var ol = OrderedList[int]{}
	ol._ascending = true
	ol.Add(1)
	ol.Delete(1)

	if ol.Count() != 0 {
		t.Error("Error")
	}

	if ol.head != nil {
		t.Error("Error")
	}
}

func TestDeleteOneTime(t *testing.T) {
	var ol = OrderedList[int]{}
	ol._ascending = true
	ol.Add(1)
	ol.Add(2)
	ol.Add(3)
	ol.Add(4)
	ol.Delete(2)

	if ol.Count() != 3 {
		t.Error("Error")
	}

	if ol.head.value != 1 {
		t.Error("Error")
	}

	if ol.head.next.value != 3 {
		t.Error("Error")
	}
}

func TestFindHas(t *testing.T) {
	var ol = OrderedList[int]{}
	ol._ascending = true
	ol.Add(1)
	ol.Add(2)
	ol.Add(3)
	ol.Add(4)
	var result, err = ol.Find(3)

	if err != nil {
		t.Error("Error")
	}

	if result.value != 3 {
		t.Error("Error")
	}
	if result.next.value != 4 {
		t.Error("Error")
	}

	if result.prev.value != 2 {
		t.Error("Error")
	}
}

func TestFindHasNot(t *testing.T) {
	var ol = OrderedList[int]{}
	ol._ascending = true
	ol.Add(1)
	ol.Add(2)
	ol.Add(3)
	ol.Add(4)
	var _, err = ol.Find(100)

	if err == nil {
		t.Error("Error")
	}
}

func TestEmpty(t *testing.T) {
	var ol = OrderedList[int]{}
	ol._ascending = true
	var _, err = ol.Find(10)

	if err == nil {
		t.Error("Error")
	}

}

func TestIn(t *testing.T) {
	var ol = OrderedList[int]{}
	ol._ascending = true
	ol.Add(0)
	ol.Add(2)
	ol.Add(1)
	ol.Add(1)
	ol.Add(2)
	ol.Add(1)
	ol.Add(0)

	if ol.Count() != 7 {
		t.Error("Error")
	}
	if ol.head.value != 0 {
		t.Error("Error")
	}

	if ol.tail.value != 2 {
		t.Error("Error")
	}

	ol.Delete(1)
	ol.Delete(2)
	ol.Delete(2)
	ol.Delete(1)
	ol.Delete(0)
	ol.Delete(1)

	if ol.Count() != 1 {
		t.Error("Error")
	}
	if ol.head.value != 0 {
		t.Error("Error")
	}
	if ol.tail.value != 0 {
		t.Error("Error")
	}
}
