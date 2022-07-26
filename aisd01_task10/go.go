package main

import (
	//	"os"
	//	"strconv"
	"constraints"
)

type PowerSet[T constraints.Ordered] struct {
	// ваша реализация хранилища
	values []T
}

func (ps *PowerSet[T]) Size() int {
	// количество элементов в множестве
	return len(ps.values)
}

func (ps *PowerSet[T]) Put(value T) {
	if !ps.Get(value) {
		ps.values = append(ps.values, value)
	}
	// всегда срабатывает
}

func (ps *PowerSet[T]) Compare(v1 T, v2 T) int {
	if v1 < v2 {
		return -1
	}
	if v1 > v2 {
		return +1
	}
	return 0
}

func (ps *PowerSet[T]) Get(value T) bool {

	// возвращает true если value имеется в множестве
	var index = 0
	for {
		if index >= ps.Size() || ps.Size() < 1 {
			break
		}

		if ps.Compare(ps.values[index], value) == 0 {
			return true
		}

		index += 1
	}

	return false
}

func (ps *PowerSet[T]) Remove(value T) bool {
	// возвращает true если value удалено
	// возвращает true если value имеется в множестве
	var defaultTypeElement T

	var index = 0
	for {

		if index >= ps.Size() || ps.Size() < 1 {
			break
		}

		if ps.Compare(ps.values[index], value) == 0 {
			copy(ps.values[index:], ps.values[index+1:])
			ps.values[len(ps.values)-1] = defaultTypeElement
			ps.values = ps.values[:len(ps.values)-1]
			return true
		}

		index += 1
	}

	return false
}

func (ps *PowerSet[T]) Intersection(set2 PowerSet[T]) PowerSet[T] {
	// пересечение текущего множества и set2
	var result PowerSet[T]
	if ps.Size() >= set2.Size() {
		var index = 0
		for {
			if index >= ps.Size() || ps.Size() < 1 {
				break
			}

			if set2.Get(ps.values[index]) {
				result.Put(ps.values[index])
			}
			index += 1
		}
	} else {
		var index = 0
		for {
			if index >= set2.Size() || set2.Size() < 1 {
				break
			}

			if ps.Get(set2.values[index]) {
				result.Put(set2.values[index])
			}

			index += 1

		}
	}

	return result
}

func (ps *PowerSet[T]) Union(set2 PowerSet[T]) PowerSet[T] {
	// объединение текущего множества и set2
	var result PowerSet[T]
	var index = 0
	for {
		if index >= ps.Size() || ps.Size() < 1 {
			break
		}
		result.Put(ps.values[index])

		index += 1

	}

	index = 0
	for {
		if index >= set2.Size() || set2.Size() < 1 {
			break
		}

		result.Put(set2.values[index])

		index += 1
	}

	return result
}

func (ps *PowerSet[T]) Difference(set2 PowerSet[T]) PowerSet[T] {
	// разница текущего множества и set2
	var result PowerSet[T]

	var index = 0
	for {

		if index >= ps.Size() || ps.Size() < 1 {
			break
		}

		if !set2.Get(ps.values[index]) {
			result.Put(ps.values[index])
		}

		index += 1

	}
	return result
}

func (ps *PowerSet[T]) IsSubset(set2 PowerSet[T]) bool {
	// возвращает true, если set2 есть
	// подмножество текущего множества

	var index = 0
	for {

		if index >= set2.Size() || set2.Size() < 1 {
			break
		}

		if !ps.Get(set2.values[index]) {
			return false
		}

		index += 1

	}

	return true
}
