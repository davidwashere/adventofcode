package util

// Stack holds any interface
type Stack []interface{}

// NewStack .
func NewStack() Stack {
	return []interface{}{}
}

// IsEmpty returns true if there is at least one item in the stack
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push adds ele to the top of the stack
func (s *Stack) Push(ele interface{}) {
	*s = append(*s, ele)
}

// Pop removes and returns top element from stack, if stack is empty returns
// empty string
func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}

	index := len(*s) - 1
	ele := (*s)[index]
	(*s)[index] = nil // Erase element (memory cleanup)
	*s = (*s)[:index]
	return ele
}

// Peek returns the top element of the stack
func (s *Stack) Peek() interface{} {
	index := len(*s) - 1
	return (*s)[index]
}
