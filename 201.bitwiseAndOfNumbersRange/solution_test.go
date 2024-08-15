package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=201 lang=golang
 *
 * [201] Bitwise AND of Numbers Range
 */

// @lc code=start
func rangeBitwiseAnd(left int, right int) int {
	cnt := 0
	for left != right {
		left >>= 1
		right >>= 1
		cnt++
	}
	return left << cnt
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		left  int
		right int
		want  int
	}{
		{5, 7, 4},
		{0, 0, 0},
		{1, 2147483647, 0},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, rangeBitwiseAnd(tt.left, tt.right))
	}
}
