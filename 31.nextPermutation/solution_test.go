package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] Next Permutation
 */

// @lc code=start
func reverse(nums []int, i, j int) {
	if i > j {
		i, j = j, i
	}
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func nextPermutation(nums []int) {
	left, right := 0, 0
	for left = len(nums) - 2; left >= 0; left-- {
		if nums[left] < nums[left+1] {
			break
		}
	}
	if left >= 0 {
		for right = len(nums) - 1; right >= 0; right-- {
			if nums[left] < nums[right] {
				break
			}
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	reverse(nums, left+1, len(nums)-1)
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  []int
		output []int
	}{
		{[]int{1, 2, 3}, []int{1, 3, 2}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 1, 5}, []int{1, 5, 1}},
		{[]int{1}, []int{1}},
	}
	for _, tt := range tc {
		nextPermutation(tt.input)
		assert.Equal(t, tt.output, tt.input)
	}
}
