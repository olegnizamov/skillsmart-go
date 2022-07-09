package main

import (
	"fmt"
	"os"
)

type Stack[T any] struct {
	stack []T
}

func (st *Stack[T]) Size() int {
	return len(st.stack)
}

// Peek - получить верхний элемент стека, но не удалять его.
func (st *Stack[T]) Peek() (T, error) {
	var result T

	if st.Size() < 1 {
		return result, fmt.Errorf("bad")
	}

	return st.stack[len(st.stack)-1], nil
}

// Pop - извлекает последний втолкнутый в стек элемент
func (st *Stack[T]) Pop() (T, error) {
	var result T
	if st.Size() < 1 {
		return result, fmt.Errorf("bad")
	}
	result = st.stack[len(st.stack)-1]
	st.stack = st.stack[:len(st.stack)-1]
	return result, nil
}

//Push - помещает элемент в этот вход -- говорят, на самый верх стека
func (st *Stack[T]) Push(itm T) {
	/** Пересоздаем стек */
	var netStack []T
	netStack = append(netStack, itm)
	var result = append(netStack, st.stack...)
	st.stack = result
}
