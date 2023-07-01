package main

import (
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 3Sum Closest
 */

// @lc code=start
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func threeSumClosest(nums []int, target int) int {
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
	var res int
	diff := math.MaxInt32
	for i := 0; i < len(nums)-2; i = next(i) {
		tmpTarget := target - nums[i]
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[left] + nums[right]
			curDiff := sum - tmpTarget
			if curDiff == 0 {
				return target
			} else if curDiff < 0 {
				left = next(left)
			} else {
				right = prev(right)
			}
			if abs(curDiff) < diff {
				diff = abs(curDiff)
				res = target + curDiff
			}
		}
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(t, 2, threeSumClosest([]int{-1, 2, 1, -4}, 1))
	assert.Equal(t, 0, threeSumClosest([]int{0, 0, 0}, 1))
	assert.Equal(t, 2, threeSumClosest([]int{1, 1, 1, 0}, -100))
	assert.Equal(t, 3, threeSumClosest([]int{1, 1, 1, 0}, 100))
}
