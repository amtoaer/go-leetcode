package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=1679 lang=golang
 *
 * [1679] Max Number of K-Sum Pairs
 */

// @lc code=start
func maxOperations(nums []int, k int) int {
	m := make(map[int]int)
	var res int
	for _, num := range nums {
		if v, ok := m[k-num]; ok && v > 0 {
			m[k-num] = v - 1
			res += 1
		} else {
			m[num]++
		}
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 2, 3, 4}, 5, 2},
		{[]int{3, 1, 3, 4, 3}, 6, 1},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, maxOperations(tt.nums, tt.k))
	}
}
