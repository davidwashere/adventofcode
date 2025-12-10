package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoint3Dist(t *testing.T) {
	tcs := []struct {
		p1   Point3
		p2   Point3
		dist float64
	}{
		{Point3{7, 4, 3}, Point3{17, 6, 2}, 10.246950765959598},
		{Point3{5, 6, 2}, Point3{-7, 11, -13}, 19.849433241279208},
		{Point3{1, 1, 0}, Point3{2, 1, 2}, 2.23606797749979},
		{Point3{0, 0, 0}, Point3{1, 0, 0}, 1},
		{Point3{0, 0, 0}, Point3{0, 1, 0}, 1},
		{Point3{0, 0, 0}, Point3{0, 0, 1}, 1},
		{Point3{1, 1, 1}, Point3{0, 1, 1}, 1},
		{Point3{1, 1, 1}, Point3{1, 0, 1}, 1},
		{Point3{1, 1, 1}, Point3{1, 1, 0}, 1},
	}
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("%v->%v", tc.p1, tc.p2), func(t *testing.T) {
			dist := tc.p1.Dist(tc.p2)
			assert.Equal(t, tc.dist, dist)
		})
	}
}
