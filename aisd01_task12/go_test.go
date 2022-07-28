package main

import (
	"testing"
)

func TestPut(t *testing.T) {
	var result = Init[string](10)
	result.Put("a", "100")
	result.Put("aa", "101")
	result.Put("aaa", "102")
	result.Put("aaaa", "103")

	if result.size != 10 {
		t.Error("Error")
	}

	if result.slots[1] != "a" {
		t.Error("Error")
	}

	if result.values[1] != "100" {
		t.Error("Error")
	}

	if result.slots[2] != "aa" {
		t.Error("Error")
	}

	if result.values[2] != "101" {
		t.Error("Error")
	}

	if result.slots[3] != "aaa" {
		t.Error("Error")
	}

	if result.values[3] != "102" {
		t.Error("Error")
	}

	if result.slots[4] != "aaaa" {
		t.Error("Error")
	}

	if result.values[4] != "103" {
		t.Error("Error")
	}
	result.Put("b", "1000")
	result.Put("c", "2000")
	result.Put("d", "2000")
	result.Put("e", "2000")
	result.Put("f", "2000")
	result.Put("g", "2000")
	result.Put("z", "100000")

}

func TestIsKey(t *testing.T) {
	var result = Init[string](10)
	result.Put("a", "100")
	result.Put("aa", "101")
	result.Put("aaa", "102")
	result.Put("aaaa", "103")

	if result.IsKey("a") != true {
		t.Error("Error")
	}

	if result.IsKey("aa") != true {
		t.Error("Error")
	}

	if result.IsKey("aaa") != true {
		t.Error("Error")
	}

	if result.IsKey("aaaa") != true {
		t.Error("Error")
	}

	if result.IsKey("b") == true {
		t.Error("Error")
	}

	if result.IsKey("aaaaa") == true {
		t.Error("Error")
	}
}

func TestGet(t *testing.T) {
	var result = Init[string](10)
	result.Put("a", "100")
	result.Put("aa", "101")
	result.Put("aaa", "102")
	result.Put("aaaa", "103")

	var answer, err = result.Get("a")

	if answer != "100" {
		t.Error("Error")
	}
	if err != nil {
		t.Error("Error")
	}

	answer, err = result.Get("aa")

	if answer != "101" {
		t.Error("Error")
	}
	if err != nil {
		t.Error("Error")
	}

	answer, err = result.Get("aaa")

	if answer != "102" {
		t.Error("Error")
	}
	if err != nil {
		t.Error("Error")
	}

	answer, err = result.Get("aaaa")

	if answer != "103" {
		t.Error("Error")
	}
	if err != nil {
		t.Error("Error")
	}

	answer, err = result.Get("bbb")

	if err == nil {
		t.Error("Error")
	}
}
