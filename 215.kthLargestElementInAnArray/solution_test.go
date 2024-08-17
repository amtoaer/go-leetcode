package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] Kth Largest Element in an Array
 */

// @lc code=start
func siftDown(nums []int, start, end int) {
	cur, next := start, start*2+1
	for next <= end {
		if next+1 <= end && nums[next+1] > nums[next] {
			next++
		}
		if nums[cur] >= nums[next] {
			return
		}
		nums[cur], nums[next] = nums[next], nums[cur]
		cur, next = next, next*2+1
	}
}

func findKthLargest(nums []int, k int) int {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		siftDown(nums, i, len(nums)-1)
	}
	i := len(nums)
	for j := 0; j < k; j++ {
		i--
		nums[i], nums[0] = nums[0], nums[i]
		siftDown(nums, 0, i-1)
	}
	return nums[i]
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  []int
		k      int
		output int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, 2, 5},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, findKthLargest(tt.input, tt.k))
	}
}
