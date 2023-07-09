package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] Permutations
 */

// @lc code=start
func permute(nums []int) [][]int {
	var (
		used      []bool = make([]bool, len(nums))
		res       [][]int
		tmp       []int
		backtrack func()
	)
	backtrack = func() {
		if len(tmp) == len(nums) {
			res = append(res, append([]int{}, tmp...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			tmp = append(tmp, nums[i])
			used[i] = true
			backtrack()
			used[i] = false
			tmp = tmp[:len(tmp)-1]
		}
	}
	backtrack()
	return res
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(t, [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}, permute([]int{1, 2, 3}))
	assert.Equal(t, [][]int{
		{0, 1},
		{1, 0},
	}, permute([]int{0, 1}))
	assert.Equal(t, [][]int{
		{1},
	}, permute([]int{1}))
}
