package util

// Stack holds any interface
type Stack[T any] []T

// NewStack .
func NewStack[T any]() Stack[T] {
	return []T{}
}

// IsEmpty returns true if there is at least one item in the stack
func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

// Push adds ele to the top of the stack
func (s *Stack[T]) Push(ele T) {
	*s = append(*s, ele)
}

// Pop removes and returns top element from stack. Popping an empty stack will
// cause panic
func (s *Stack[T]) Pop() T {
	if s.IsEmpty() {
		panic("pop on empty stack")
	}

	index := len(*s) - 1
	ele := (*s)[index]
	(*s)[index] = *new(T) // Erase element (memory cleanup).. or does this ad memory
	*s = (*s)[:index]
	return ele
}

// Peek returns the top element of the stack
func (s *Stack[T]) Peek() T {
	index := len(*s) - 1
	return (*s)[index]
}
