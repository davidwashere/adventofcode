package util

import "testing"

func TestQueueInt(t *testing.T) {

	q := NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	vf(t, q.Dequeue(), 1)
	vf(t, q.Peek(), 2)
	vf(t, q.Front(), 2)
	vf(t, q.Back(), 3)
	vf(t, q.IsEmpty(), false)
	vf(t, q.Pop(), 3)
	vf(t, len(q), 1)
	vf(t, q.Pop(), 2)
	vf(t, q.IsEmpty(), true)
}

func TestQueueString(t *testing.T) {

	q := NewQueue[string]()
	q.Enqueue("A")
	q.Enqueue("B")
	q.Enqueue("C")

	vf(t, q.Dequeue(), "A")
	vf(t, q.Peek(), "B")
	vf(t, q.Front(), "B")
	vf(t, q.Back(), "C")
	vf(t, q.IsEmpty(), false)
	vf(t, q.Pop(), "C")
	vf(t, len(q), 1)
	vf(t, q.Pop(), "B")
	vf(t, q.IsEmpty(), true)
}

func TestQueuePanic(t *testing.T) {
	s := NewQueue[int]()
	defer func() { _ = recover() }()
	s.Pop()
	t.Errorf("did not panic")
}
