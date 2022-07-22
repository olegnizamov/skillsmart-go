package main

import (
	"strconv"
	"os"
)

type HashTable struct {
	size  int
	step  int
	slots []string
}

func Init(sz int, stp int) HashTable {
	ht := HashTable{size: sz, step: stp, slots: nil}
	ht.slots = make([]string, sz)
	return ht
}

func HashFun(value string) int {
	// всегда возвращает корректный индекс слота
	var index = len(value) % 100
	return index
}

func (ht *HashTable) SeekSlot(value string) int {
	// находит индекс пустого слота для значения,
	// или -1
	var index = HashFun(value)
	if ht.slots[index] == "" {
		return index
	}
	var i = 0
	for {
		if i >= ht.size {
			break
		}
		index += ht.step

		if index >= ht.size {
			index = index - ht.size
		}
		if ht.slots[index] == "" {
			return index
		} else {
			i += 1
		}
	}

	return -1
}

func (ht *HashTable) Put(value string) int {
	// записываем значение по хэш-функции
	var slot = ht.SeekSlot(value)
	if slot != -1 {
		ht.slots[slot] = value
	}
	return slot
	// возвращается индекс слота или -1
	// если из-за коллизий элемент не удаётся разместить
}

func (ht *HashTable) Find(value string) int {
	// находит индекс слота со значением, или -1
	var hashIndex = HashFun(value)

	if ht.slots[hashIndex] == value {
		return hashIndex
	}

	var i = 0
	for {
		if i >= ht.size {
			break
		}
		hashIndex += ht.step

		if hashIndex >= ht.size {
			hashIndex = hashIndex - ht.size
		}
		if ht.slots[hashIndex] == value {
			return hashIndex
		} else {
			i += 1
		}
	}

	// находит индекс слота со значением, или -1
	return -1
}
