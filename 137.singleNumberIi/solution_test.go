package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=137 lang=golang
 *
 * [137] Single Number II
 */

// @lc code=start
func singleNumber(nums []int) int {
	var x, y int
	for _, num := range nums {
		x = (x ^ num) & ^y
		y = (y ^ num) & ^x
	}
	return x
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		nums []int
		want int
	}{
		{[]int{2, 2, 3, 2}, 3},
		{[]int{0, 1, 0, 1, 0, 1, 99}, 99},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, singleNumber(tt.nums))
	}
}
