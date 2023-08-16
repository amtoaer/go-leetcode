package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=2682 lang=golang
 *
 * [2682] Find the Losers of the Circular Game
 */

// @lc code=start
func circularGameLosers(n int, k int) []int {
	var (
		flag    = make([]bool, n)
		count   = 0
		current = 0
	)
	for i := 0; ; i++ {
		current = (current + i*k) % n
		if flag[current] {
			res := make([]int, 0, n-count)
			for j := 0; j < n; j++ {
				if !flag[j] {
					res = append(res, j+1)
				}
			}
			return res
		}
		flag[current] = true
		count += 1
	}
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(t, []int{4, 5}, circularGameLosers(5, 2))
	assert.Equal(t, []int{2, 3, 4}, circularGameLosers(4, 4))
}
