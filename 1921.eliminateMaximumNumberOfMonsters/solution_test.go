package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=1921 lang=golang
 *
 * [1921] Eliminate Maximum Number of Monsters
 */

// @lc code=start
func eliminateMaximum(dist []int, speed []int) int {
	for i := 0; i < len(dist); i++ {
		dist[i] = (dist[i]-1)/speed[i] + 1
	}
	sort.Ints(dist)
	for idx, minute := range dist {
		if minute <= idx {
			return idx
		}
	}
	return len(dist)
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(t, 3, eliminateMaximum([]int{1, 3, 4}, []int{1, 1, 1}))
	assert.Equal(t, 1, eliminateMaximum([]int{1, 1, 2, 3}, []int{1, 1, 1, 1}))
}
