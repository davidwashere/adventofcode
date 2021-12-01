package util

import "testing"

func TestStringStack(t *testing.T) {
	s := NewStringStack()

	s.Push("A")
	s.Push("B")
	s.Push("C")

	vf(t, s.IsEmpty(), false)
	vf(t, s.Pop(), "C")
	vf(t, s.Peek(), "B")
	vf(t, len(s), 2)
	vf(t, s.Pop(), "B")
	vf(t, s.Pop(), "A")
	vf(t, s.IsEmpty(), true)
	vf(t, s.Pop(), "")
}

func TestIntStack(t *testing.T) {
	s := NewIntStack()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	vf(t, s.IsEmpty(), false)
	vf(t, s.Pop(), 3)
	vf(t, s.Peek(), 2)
	vf(t, len(s), 2)
	vf(t, s.Pop(), 2)
	vf(t, s.Pop(), 1)
	vf(t, s.IsEmpty(), true)
	vf(t, s.Pop(), 0)
}

func TestStack(t *testing.T) {
	s := NewStack()

	type cat struct {
		name string
	}

	s.Push(cat{"Yuck"})
	s.Push(cat{"Smelly"})
	s.Push(cat{"Ugly"})

	vf(t, s.IsEmpty(), false)
	vf(t, s.Pop().(cat).name, "Ugly")
	vf(t, s.Peek().(cat).name, "Smelly")
	vf(t, len(s), 2)
	vf(t, s.Pop().(cat).name, "Smelly")
	vf(t, s.Pop().(cat).name, "Yuck")
	vf(t, s.IsEmpty(), true)
	vf(t, s.Pop(), nil)
}
