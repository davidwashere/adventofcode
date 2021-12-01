package util

import (
	"fmt"
	"testing"
)

func TestCountPaths(t *testing.T) {
	tree := NewTree()
	tree.AddChild(0, 1)
	tree.AddChild(0, 2)
	tree.AddChild(0, 3)
	tree.AddChild(1, 2)
	tree.AddChild(1, 3)
	tree.AddChild(1, 4)
	tree.AddChild(2, 3)
	tree.AddChild(2, 4)
	tree.AddChild(3, 4)
	tree.AddChild(4, 7)

	list := []int{0, 1, 2, 3, 4, 7}
	for i := range list {
		fmt.Println(tree.StringNode(i))
	}

	got := tree.CountPaths(0, 7)
	want := 7

	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}
