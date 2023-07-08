package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=167 lang=golang
 *
 * [167] Two Sum II - Input Array Is Sorted
 */

// @lc code=start
func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum > target {
			right--
		} else if sum < target {
			left++
		} else {
			return []int{left + 1, right + 1}
		}
	}
	return nil
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(t, []int{1, 2}, twoSum([]int{2, 7, 11, 15}, 9))
	assert.Equal(t, []int{1, 3}, twoSum([]int{2, 3, 4}, 6))
}
