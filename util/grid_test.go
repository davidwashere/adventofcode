package util

import (
	"fmt"
	"testing"
)

func TestBlah(t *testing.T) {
	g := NewInfGrid()

	g.Set("a", 0, 0)
	g.Set("b", 0, 1)

	fmt.Println(g.Get(0, 0))
	fmt.Println(g.Get(0, 1))
	fmt.Println(g.Get(0, 2))
}
