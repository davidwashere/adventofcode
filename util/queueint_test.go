package util

import "testing"

func TestPopLargest(t *testing.T) {
	q := IntQueue{}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(10)
	q.Enqueue(3)
	q.Enqueue(8)

	var want int

	want = 10
	got := q.PopLargest()
	vf(t, got, want)

	want = 8
	got = q.PopLargest()
	vf(t, got, want)

	want = 3
	got = q.PopLargest()
	vf(t, got, want)

	want = 2
	got = q.PopLargest()
	vf(t, got, want)

	want = 1
	got = q.PopLargest()
	vf(t, got, want)
}

func TestPopSmallest(t *testing.T) {
	q := IntQueue{}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(10)
	q.Enqueue(3)
	q.Enqueue(8)

	var want int

	want = 1
	got := q.PopSmallest()
	vf(t, got, want)

	want = 2
	got = q.PopSmallest()
	vf(t, got, want)

	want = 3
	got = q.PopSmallest()
	vf(t, got, want)

	want = 8
	got = q.PopSmallest()
	vf(t, got, want)

	want = 10
	got = q.PopSmallest()
	vf(t, got, want)
}
