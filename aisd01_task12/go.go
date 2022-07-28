package main

import (
	"fmt"
)

type NativeCache[T any] struct {
	size   int
	slots  []string
	values []T
	hits   []int
}

// создание экземпляра словаря
func Init[T any](sz int) NativeCache[T] {
	nd := NativeCache[T]{size: sz, slots: nil, values: nil}
	nd.slots = make([]string, sz)
	nd.values = make([]T, sz)
	nd.hits = make([]int, sz)
	return nd
}

func HashFun(value string) int {
	// всегда возвращает корректный индекс слота
	var index = len(value) % 100
	return index
}

func (nd *NativeCache[T]) IsKey(key string) bool {
	// возвращает true если ключ имеется
	if nd.slots[HashFun(key)] == key {
		nd.hits[HashFun(key)] += 1
		return true
	}
	return false
}

func (nd *NativeCache[T]) Get(key string) (T, error) {
	// возвращает value для key,
	// или error если ключ не найден
	var result T

	if nd.IsKey(key) {
		return nd.values[HashFun(key)], nil
	}

	return result, fmt.Errorf("not found")
}

func (nd *NativeCache[T]) SeekSlot(value string) int {
	// находит индекс пустого слота для значения,
	// или -1
	var index = HashFun(value)
	if nd.slots[index] == "" {
		return index
	}
	var i = 0
	for {
		if i >= nd.size {
			break
		}
		index += 1

		if index >= nd.size {
			index = index - nd.size
		}
		if nd.slots[index] == "" {
			return index
		} else {
			i += 1
		}
	}

	return -1
}

func (nd *NativeCache[T]) MinSeekSlot() int {
	// находит индекс пустого слота для значения,
	// или -1
	var min = nd.hits[0]
	var minKey = 0

	for key, val := range nd.hits {
		if min > val {
			min = val
			minKey = key
		}
	}

	return minKey
}

func (nd *NativeCache[T]) Put(key string, value T) {
	// записываем значение по хэш-функции
	var slot = nd.SeekSlot(key)
	if slot == -1 {
		// нет пустых - тогда ищем минимальное значение
		slot = nd.MinSeekSlot()
		nd.hits[slot] = 0
	}
	nd.slots[slot] = key
	nd.values[slot] = value
}
