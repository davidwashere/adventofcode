package util

// IntQueue holds ints
type IntQueue []int

// NewIntQueue .
func NewIntQueue() IntQueue {
	return []int{}
}

// IsEmpty returns true if there is at least one item in the stack
func (s *IntQueue) IsEmpty() bool {
	return len(*s) == 0
}

// Enqueue .
func (s *IntQueue) Enqueue(ele int) {
	*s = append(*s, ele)
}

// Dequeue .
func (s *IntQueue) Dequeue() int {
	result := (*s)[0]
	(*s)[0] = 0
	(*s) = (*s)[1:]
	return result
}

// Pop removes and returns top element from stack, if stack is empty returns
// empty string
func (s *IntQueue) Pop() int {
	if s.IsEmpty() {
		return 0
	}

	index := len(*s) - 1
	ele := (*s)[index]
	(*s)[index] = 0 // Erase element (memory cleanup)
	*s = (*s)[:index]
	return ele
}

// Front .
func (s *IntQueue) Front() int {
	// index := len(*s) - 1
	return (*s)[0]
}

// Back .
func (s *IntQueue) Back() int {
	index := len(*s) - 1
	return (*s)[index]
}
