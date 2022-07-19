package main

import (
	"testing"
)

func TestHashFun(t *testing.T) {
	var result = HashFun("aaa")
	if result != 3 {
		t.Error("Error")
	}

	result = HashFun("aa")
	if result != 2 {
		t.Error("Error")
	}

	result = HashFun("aab")
	if result != 3 {
		t.Error("Error")
	}

	result = HashFun("a")
	if result != 1 {
		t.Error("Error")
	}
}

func TestPut(t *testing.T) {
	var ht = Init(17, 3)
	var result1 = ht.Put("aaa")
	var result2 = ht.Put("aa")

	if result1 != 3 {
		t.Error("Error")
	}

	if result2 != 2 {
		t.Error("Error")
	}
}

func TestSeekSlot(t *testing.T) {

	var ht = Init(17, 3)
	var result = ht.SeekSlot("aaa")
	if result != 3 {
		t.Error("Error")
	}

	ht = Init(17, 3)
	result = HashFun("aa")
	if result != 2 {
		t.Error("Error")
	}

}

func TestSeekSlotAndPut(t *testing.T) {

	var ht = Init(17, 3)
	_ = ht.Put("aaa")
	_ = ht.Put("aa")
	_ = ht.Put("bb")
	_ = ht.Put("bbb")
	_ = ht.Put("cc")
	_ = ht.Put("ccc")
	_ = ht.Put("dd")
	_ = ht.Put("ddd")
	_ = ht.Put("ee")
	_ = ht.Put("eee")
	_ = ht.Put("ff")
	_ = ht.Put("fff")
	_ = ht.Put("gg")
	_ = ht.Put("ggg")
	_ = ht.Put("vv")
	_ = ht.Put("vvv")

	var result1 = ht.Put("ww")
	var result2 = ht.Put("www")

	if result1 != 16 {
		t.Error("Error")
	}

	if result2 != -1 {
		t.Error("Error")
	}

}

func TestFind(t *testing.T) {

	var ht = Init(17, 3)
	_ = ht.Put("aaa")
	_ = ht.Put("aa")
	_ = ht.Put("bb")
	_ = ht.Put("bbb")
	_ = ht.Put("cc")
	_ = ht.Put("ccc")
	_ = ht.Put("dd")
	_ = ht.Put("ddd")
	_ = ht.Put("ee")
	_ = ht.Put("eee")
	_ = ht.Put("ff")
	_ = ht.Put("fff")
	_ = ht.Put("gg")
	_ = ht.Put("ggg")
	_ = ht.Put("vv")
	_ = ht.Put("vvv")

	var result1 = ht.Find("vv")
	var result2 = ht.Find("www")

	if result1 == -1 {
		t.Error("Error")
	}

	if result2 != -1 {
		t.Error("Error")
	}

}
