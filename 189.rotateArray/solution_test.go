package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] Rotate Array
 */

// @lc code=start
func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-i-1] = a[n-i-1], a[i]
	}
}

func rotate(nums []int, k int) {
	l := len(nums)
	k = k % l
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
	}
	for _, tt := range tc {
		rotate(tt.nums, tt.k)
		assert.Equal(t, tt.want, tt.nums)
	}
}
