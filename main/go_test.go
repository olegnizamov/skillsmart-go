package main

import (
	"testing"
)

func TestInit(t *testing.T) {
	var da = DynArray[int]{}
	da.count = 10
	da.Init()
	if da.capacity != 16 {
		t.Error("Expected 16, got ", da.capacity)
	}
	if da.count != 0 {
		t.Error("Expected 0, got ", da.count)
	}
}

func TestMakeArray(t *testing.T) {

	var da = new(DynArray[int])
	da.Init()
	da.MakeArray(20)
	if da.capacity != 20 {
		t.Error("Expected 20, got ", da.capacity)
	}
	if da.count != 0 {
		t.Error("Expected 0, got ", da.count)
	}
	da.Append(1)
	da.Append(2)
	da.Append(3)
	da.Append(4)
	da.Append(5)
	da.MakeArray(50)
	if da.capacity != 50 {
		t.Error("Expected 50, got ", da.capacity)
	}
	if da.count != 5 {
		t.Error("Expected 5, got ", da.count)
	}
}

func TestAppend(t *testing.T) {

	var da = new(DynArray[int])
	da.Init()
	da.MakeArray(20)
	if da.capacity != 20 {
		t.Error("Expected 20, got ", da.capacity)
	}
	if da.count != 0 {
		t.Error("Expected 0, got ", da.count)
	}
	da.Append(1)
	da.Append(2)
	da.Append(3)
	da.Append(4)
	da.Append(5)

	for i := 0; i < da.count; i++ {
		if da.array[i] != (i + 1) {
			t.Error("got ", da.array[i])
		}
	}
}

func TestGetItem(t *testing.T) {

	var da = new(DynArray[int])
	da.Init()
	da.MakeArray(20)
	if da.capacity != 20 {
		t.Error("Expected 20, got ", da.capacity)
	}
	if da.count != 0 {
		t.Error("Expected 0, got ", da.count)
	}
	da.Append(1)
	da.Append(2)
	da.Append(3)
	da.Append(4)
	da.Append(5)

	for i := 0; i < da.count; i++ {
		var result, err = da.GetItem(i)
		if result != (i + 1) {
			t.Error("got ", da.array[i])
		}
		if err != nil {
			t.Error("got error")
		}
	}
	var _, err = da.GetItem(100)
	if err == nil {
		t.Error("error has not")
	}
}

func TestRemove(t *testing.T) {

	var da = new(DynArray[int])
	da.Init()
	da.MakeArray(20)

	if da.capacity != 20 {
		t.Error("Expected 20, got ", da.capacity)
	}

	if da.count != 0 {
		t.Error("Expected 0, got ", da.count)
	}

	da.Append(1)
	da.Append(2)
	da.Append(3)
	da.Append(4)
	da.Append(5)

	var err = da.Remove(0)

	for i := 0; i < da.count; i++ {
		var result, err = da.GetItem(i)
		if result != (i + 2) {
			t.Error("got ", da.array[i])
		}
		if err != nil {
			t.Error("got error")
		}
	}

	err = da.Remove(100)
	if err == nil {
		t.Error("error has not")
	}
}

func TestInsert(t *testing.T) {

	var da = new(DynArray[int])
	da.Init()
	da.MakeArray(20)
	if da.capacity != 20 {
		t.Error("Expected 20, got ", da.capacity)
	}
	if da.count != 0 {
		t.Error("Expected 0, got ", da.count)
	}
	var _ = da.Insert(1, 0)
	_ = da.Insert(2, 1)
	_ = da.Insert(3, 2)
	_ = da.Insert(4, 3)
	_ = da.Insert(5, 4)

	for i := 0; i < da.count; i++ {
		if da.array[i] != (i + 1) {
			t.Error("got ", da.array[i])
		}
	}

	var err = da.Insert(5, 100)
	if err == nil {
		t.Error("error has not")
	}
}
