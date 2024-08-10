package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] Combination Sum
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	var (
		result    = [][]int{}
		tmpResult []int
		backtrack func(idx int, left int)
	)
	backtrack = func(idx int, left int) {
		if left == 0 {
			result = append(result, append([]int{}, tmpResult...))
		}
		for index := idx; index < len(candidates); index++ {
			candidate := candidates[index]
			if left < candidate {
				continue
			}
			tmpResult = append(tmpResult, candidate)
			backtrack(index, left-candidate)
			tmpResult = tmpResult[:len(tmpResult)-1]
		}
	}
	backtrack(0, target)
	return result
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		candidates []int
		target     int
		ans        [][]int
	}{
		{[]int{2, 3, 6, 7}, 7, [][]int{{2, 2, 3}, {7}}},
		{[]int{2, 3, 5}, 8, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
		{[]int{2}, 1, [][]int{}},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.ans, combinationSum(tt.candidates, tt.target))
	}
}
