package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	next := func(i int) int {
		for j := i; j < len(nums); j++ {
			if nums[j] == nums[i] {
				continue
			}
			return j
		}
		return len(nums)
	}
	prev := func(i int) int {
		for j := i; j >= 0; j-- {
			if nums[j] == nums[i] {
				continue
			}
			return j
		}
		return -1
	}
	var res [][]int
	for i := 0; i < len(nums)-2; i = next(i) {
		target := -nums[i]
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[left] + nums[right]
			if sum > target {
				right = prev(right)
			} else if sum < target {
				left = next(left)
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left = next(left)
			}
		}
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(t, [][]int{{-1, -1, 2}, {-1, 0, 1}}, threeSum([]int{-1, 0, 1, 2, -1, -4}))
	assert.Equal(t, [][]int{{0, 0, 0}}, threeSum([]int{0, 0, 0, 0}))
	assert.Equal(t, [][]int{{-2, 0, 2}}, threeSum([]int{-2, 0, 0, 2, 2}))
	assert.Equal(t, [][]int{{-2, 0, 2}}, threeSum([]int{-2, 0, 0, 2, 2}))
}
