package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] Contains Duplicate II
 */

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i, n := range nums {
		if j, ok := m[n]; ok && (i-j) <= k {
			return true
		}
		m[n] = i
	}
	return false
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		nums   []int
		k      int
		output bool
	}{
		{[]int{1, 2, 3, 1}, 3, true},
		{[]int{1, 0, 1, 1}, 1, true},
		{[]int{1, 2, 3, 1, 2, 3}, 2, false},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, containsNearbyDuplicate(tt.nums, tt.k))
	}
}
