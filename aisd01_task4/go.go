package main

import "fmt"

type Stack[T any] struct {
	stack []T
}

func (st *Stack[T]) Size() int {
	return len(st.stack)
}

func (st *Stack[T]) Peek() (T, error) {
	var result T

	if st.Size() < 1 {
		return result, fmt.Errorf("bad")
	}

	return st.stack[len(st.stack)-1], nil
}

func (st *Stack[T]) Pop() (T, error) {
	var result T

	if st.Size() < 1 {
		return result, fmt.Errorf("bad")
	}
	var last T = st.stack[len(st.stack)-1]
	st.stack[len(st.stack)-1] = nil
	return last, nil
}

func (st *Stack[T]) Push(itm T) {
	st.stack = append(st.stack, itm)
}
