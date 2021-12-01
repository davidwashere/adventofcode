package util

// StringStack holds strings
type StringStack []string

// NewStringStack .
func NewStringStack() StringStack {
	return []string{}
}

// IsEmpty returns true if there is at least one item in the stack
func (s *StringStack) IsEmpty() bool {
	return len(*s) == 0
}

// Push adds ele to the top of the stack
func (s *StringStack) Push(ele string) {
	*s = append(*s, ele)
}

// Pop removes and returns top element from stack, if stack is empty returns
// empty string
func (s *StringStack) Pop() string {
	if s.IsEmpty() {
		return ""
	}

	index := len(*s) - 1
	ele := (*s)[index]
	(*s)[index] = "" // Erase element (memory cleanup)
	*s = (*s)[:index]
	return ele
}

// Peek returns the top element of the stack
func (s *StringStack) Peek() string {
	index := len(*s) - 1
	return (*s)[index]
}
