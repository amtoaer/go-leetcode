package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=3143 lang=golang
 *
 * [3143] Maximum Points Inside the Square
 */

// @lc code=start
func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func maxPointsInsideSquare(points [][]int, s string) int {
	edge := math.MaxInt
	m := make([]int, 26)
	for i := range m {
		m[i] = math.MaxInt
	}
	for i := 0; i < len(s); i++ {
		x, y := points[i][0], points[i][1]
		idx := int(s[i] - 'a')
		d := max(abs(x), abs(y))
		if d < m[idx] {
			edge = min(edge, m[idx])
			m[idx] = d
		} else {
			edge = min(edge, d)
		}
	}
	cnt := 0
	for _, d := range m {
		if d < edge {
			cnt++
		}
	}
	return cnt
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		points [][]int
		s      string
		expect int
	}{
		{[][]int{{2, 2}, {-1, -2}, {-4, 4}, {-3, 1}, {3, -3}}, "abcda", 2},
		{[][]int{{1, 1}, {-2, -2}, {-2, 2}}, "abb", 1},
	}
	for _, tc := range tc {
		assert.Equal(t, tc.expect, maxPointsInsideSquare(tc.points, tc.s))
	}
}
