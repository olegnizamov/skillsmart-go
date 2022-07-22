package main

import "fmt"

type NativeDictionary[T any] struct {
	size   int
	slots  []string
	values []T
}

// создание экземпляра словаря
func Init[T any](sz int) NativeDictionary[T] {
	nd := NativeDictionary[T]{size: sz, slots: nil, values: nil}
	nd.slots = make([]string, sz)
	nd.values = make([]T, sz)
	return nd
}

func HashFun(value string) int {
	// всегда возвращает корректный индекс слота
	var index = len(value) % 100
	return index
}

func (nd *NativeDictionary[T]) IsKey(key string) bool {
	// возвращает true если ключ имеется
	return nd.slots[HashFun(key)] == key
}

func (nd *NativeDictionary[T]) Get(key string) (T, error) {
	// возвращает value для key,
	// или error если ключ не найден
	var result T

	if nd.IsKey(key) {
		return nd.values[HashFun(key)], nil
	}

	return result, fmt.Errorf("not found")
}

func (nd *NativeDictionary[T]) Put(key string, value T) {
	// гарантированно записываем
	// значение value по ключу key
	var index = HashFun(key)
	nd.slots[index] = key
	nd.values[index] = value
}
