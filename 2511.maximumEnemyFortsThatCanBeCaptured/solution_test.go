package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=2511 lang=golang
 *
 * [2511] Maximum Enemy Forts That Can Be Captured
 */

// @lc code=start
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func captureForts(forts []int) int {
	var target, start, res int
	for idx, fort := range forts {
		if fort == 0 {
			continue
		}
		if fort == target {
			res = max(res, idx-start-1)
		}
		start = idx
		target = -1 / fort
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(t, 4, captureForts([]int{1, 0, 0, -1, 0, 0, 0, 0, 1}))
	assert.Equal(t, 0, captureForts([]int{0, 0, 1, -1}))
}
