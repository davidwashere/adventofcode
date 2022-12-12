package util

type Queue[T any] []T

// NewIntQueue .
func NewQueue[T any]() Queue[T] {
	return []T{}
}

// IsEmpty returns true if there is at least one item in the stack
func (s *Queue[T]) IsEmpty() bool {
	return len(*s) == 0
}

// Enqueue adds to end of queu
func (s *Queue[T]) Enqueue(ele T) {
	*s = append(*s, ele)
}

// Dequeue removes from beginning of queue
func (s *Queue[T]) Dequeue() T {
	result := (*s)[0]
	(*s)[0] = *new(T)
	(*s) = (*s)[1:]
	return result
}

// Pop removes from end of queue (as if it were a stack)
func (s *Queue[T]) Pop() T {
	if s.IsEmpty() {
		panic("pop on empty queue")
	}

	index := len(*s) - 1
	ele := (*s)[index]
	(*s)[index] = *new(T) // Erase element (memory cleanup)
	*s = (*s)[:index]
	return ele
}

// Front peek at front of queue
func (s *Queue[T]) Front() T {
	return s.Peek()
}

// Front peek at front of queue
func (s *Queue[T]) Peek() T {
	return (*s)[0]
}

// Back Peeks at the back of the queue (most recently enqueued item)
func (s *Queue[T]) Back() T {
	index := len(*s) - 1
	return (*s)[index]
}
