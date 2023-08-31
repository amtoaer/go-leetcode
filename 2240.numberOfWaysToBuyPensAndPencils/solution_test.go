package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=2240 lang=golang
 *
 * [2240] Number of Ways to Buy Pens and Pencils
 */

// @lc code=start
func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {
	var res, cnt int
	if cost1 < cost2 {
		cost1, cost2 = cost2, cost1
	}
	for cnt*cost1 <= total {
		res += (total-cnt*cost1)/cost2 + 1
		cnt++
	}
	return int64(res)
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(t, int64(9), waysToBuyPensPencils(20, 10, 5))
	assert.Equal(t, int64(1), waysToBuyPensPencils(5, 10, 10))
}
