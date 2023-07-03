package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 4Sum
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var res [][]int
	nextM := make(map[int]int)
	prevM := make(map[int]int)
	next := func(i int) int {
		var (
			j  int
			ok bool
		)
		if j, ok = nextM[i]; ok {
			return j
		}
		for j = i; j < len(nums); j++ {
			if nums[j] == nums[i] {
				continue
			}
			break
		}
		nextM[i] = j
		return j
	}
	prev := func(i int) int {
		var (
			j  int
			ok bool
		)
		if j, ok = prevM[i]; ok {
			return j
		}
		for j = i; j >= 0; j-- {
			if nums[j] == nums[i] {
				continue
			}
			break
		}
		prevM[i] = j
		return j
	}
	for i := 0; i < len(nums)-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i = next(i) {
		for j := i + 1; j < len(nums)-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j = next(j) {
			tmpTarget := target - nums[i] - nums[j]
			left, right := j+1, len(nums)-1
			for left < right {
				sum := nums[left] + nums[right]
				if sum == tmpTarget {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					left = next(left)
					right = prev(right)
				} else if sum < tmpTarget {
					left = next(left)
				} else {
					right = prev(right)
				}
			}
		}
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(t, [][]int{{0, 1, 2, 3}}, fourSum([]int{0, 1, 2, 3}, 6))
	assert.Equal(t, [][]int{{0, 0, 0, 0}}, fourSum([]int{0, 0, 0, 0}, 0))
	assert.Equal(
		t,
		[][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		fourSum([]int{1, 0, -1, 0, -2, 2}, 0),
	)
}
