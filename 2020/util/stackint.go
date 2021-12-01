package util

// IntStack holds ints
type IntStack []int

// NewIntStack .
func NewIntStack() IntStack {
	return []int{}
}

// IsEmpty returns true if there is at least one item in the stack
func (s *IntStack) IsEmpty() bool {
	return len(*s) == 0
}

// Push adds ele to the top of the stack
func (s *IntStack) Push(ele int) {
	*s = append(*s, ele)
}

// Pop removes and returns top element from stack, if stack is empty returns
// empty string
func (s *IntStack) Pop() int {
	if s.IsEmpty() {
		return 0
	}

	index := len(*s) - 1
	ele := (*s)[index]
	(*s)[index] = 0 // Erase element (memory cleanup)
	*s = (*s)[:index]
	return ele
}

// Peek returns the top element of the stack
func (s *IntStack) Peek() int {
	index := len(*s) - 1
	return (*s)[index]
}
