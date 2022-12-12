package util

import "testing"

func TestStack(t *testing.T) {

	type cat struct {
		name string
	}
	s := NewStack[cat]()

	s.Push(cat{"Yuck"})
	s.Push(cat{"Smelly"})
	s.Push(cat{"Ugly"})

	vf(t, s.IsEmpty(), false)
	vf(t, s.Pop().name, "Ugly")
	vf(t, s.Peek().name, "Smelly")
	vf(t, len(s), 2)
	vf(t, s.Pop().name, "Smelly")
	vf(t, s.Pop().name, "Yuck")
	vf(t, s.IsEmpty(), true)

	s2 := NewStack[string]()

	s2.Push("A")
	s2.Push("B")
	s2.Push("C")

	vf(t, s2.IsEmpty(), false)
	vf(t, s2.Pop(), "C")
	vf(t, s2.Peek(), "B")
	vf(t, len(s2), 2)
	vf(t, s2.Pop(), "B")
	vf(t, s2.Pop(), "A")
	vf(t, s2.IsEmpty(), true)

	s3 := NewStack[int]()

	s3.Push(1)
	s3.Push(2)
	s3.Push(3)

	vf(t, s3.IsEmpty(), false)
	vf(t, s3.Pop(), 3)
	vf(t, s3.Peek(), 2)
	vf(t, len(s3), 2)
	vf(t, s3.Pop(), 2)
	vf(t, s3.Pop(), 1)
	vf(t, s3.IsEmpty(), true)
}

func TestStackPanic(t *testing.T) {
	s := NewStack[int]()
	defer func() { _ = recover() }()
	s.Pop()
	t.Errorf("did not panic")
}
