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
// 0
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

// Front alias for Peek
func (s *IntQueue) Front() int {
	return s.Peek()
}

// Front Peeks at the front of the queue (oldest enqueued int)
func (s *IntQueue) Peek() int {
	return (*s)[0]
}

// Back Peeks at the back of the queue (most recently enqueued int)
func (s *IntQueue) Back() int {
	index := len(*s) - 1
	return (*s)[index]
}

// PopLargest will pop the largest int vlaue in the queue - O(n) time, if queue
// is empty returns 0
func (s *IntQueue) PopLargest() int {
	if s.IsEmpty() {
		return 0
	}

	index := -1
	max := MinInt
	for i, val := range *s {
		if val > max {
			index = i
			max = val
		}
	}

	(*s)[index] = 0 // Erase element (memory cleanup)
	(*s) = RemoveIndexFromIntSlice((*s), index)

	return max
}

// PopSmallest will pop the smallest int vlaue in the queue - O(n) time, if queue
// is empty returns 0
func (s *IntQueue) PopSmallest() int {
	if s.IsEmpty() {
		return 0
	}

	index := -1
	min := MaxInt
	for i, val := range *s {
		if val < min {
			index = i
			min = val
		}
	}

	(*s)[index] = 0 // Erase element (memory cleanup)
	(*s) = RemoveIndexFromIntSlice((*s), index)

	return min
}
