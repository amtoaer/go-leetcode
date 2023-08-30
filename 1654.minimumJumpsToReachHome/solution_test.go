package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=1654 lang=golang
 *
 * [1654] Minimum Jumps to Reach Home
 */

// @lc code=start
type Stat struct {
	direction int
	location  int
	step      int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minimumJumps(forbidden []int, a int, b int, x int) int {
	maxPos := 0
	forbiddenSet := make(map[int]struct{})
	for _, pos := range forbidden {
		maxPos = max(maxPos, pos)
		forbiddenSet[pos] = struct{}{}
	}
	lower, upper := 0, max(maxPos+a, x)+b
	visited := make(map[int]struct{})
	queue := []Stat{
		{1, 0, 0},
	}
	for len(queue) > 0 {
		queueLen := len(queue)
		for _, stat := range queue[:queueLen] {
			if stat.location == x {
				return stat.step
			}
			_, isVisited := visited[stat.location*stat.direction]
			_, isForbidden := forbiddenSet[stat.location]
			if isVisited || isForbidden {
				continue
			}
			visited[stat.location*stat.direction] = struct{}{}
			if lower <= stat.location+a && stat.location+a <= upper {
				queue = append(queue, Stat{1, stat.location + a, stat.step + 1})
			}
			if stat.direction == 1 && lower <= stat.location-b && stat.location-b <= upper {
				queue = append(queue, Stat{-1, stat.location - b, stat.step + 1})
			}
		}
		queue = queue[queueLen:]
	}
	return -1
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(t, 3, minimumJumps([]int{14, 4, 18, 1, 15}, 3, 15, 9))
	assert.Equal(t, -1, minimumJumps([]int{8, 3, 16, 6, 12, 20}, 15, 13, 11))
	assert.Equal(t, 20, minimumJumps([]int{3}, 14, 5, 90))
}
