package main

import "fmt"

type Number interface {
	int64 | float64
}

func Max[T any](values []T, less func(a, b T) bool) T {
	if len(values) == 0 {
		var zero T
		return zero
	}
	max := values[0]
	for _, value := range values[1:] {
		if less(value, max) {
			max = value
		}
	}
	return max
}

type Stack[T any] struct {
	elements []T
}

// NewStack creates a new stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Push adds an element to the stack
func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

// Pop removes and returns the top element from the stack
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, true
}

// Peek returns the top element without removing it
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	return s.elements[len(s.elements)-1], true
}

// Size returns the number of elements in the stack
func (s *Stack[T]) Size() int {
	return len(s.elements)
}

func main() {
	intStack := NewStack[int]()
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)

	fmt.Println("Int Stack Size:", intStack.Size())

	if elem, ok := intStack.Pop(); ok {
		fmt.Println("Popped from int stack:", elem)
	}

	strStack := NewStack[string]()
	strStack.Push("hello")
	strStack.Push("world")

	fmt.Println("Str Stack Size:", strStack.Size())

	if elem, ok := strStack.Peek(); ok {
		fmt.Println("Top of str stack:", elem)
	}
}

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// SumIntsOrFloats sums the values of map m. It supports both floats and integers
// as map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// SumNumbers sums the values of map m. Its supports both integers
// and floats as map values.
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
