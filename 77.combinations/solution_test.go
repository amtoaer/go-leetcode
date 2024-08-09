package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] Combinations
 */

// @lc code=start
func combine(n int, k int) [][]int {
	var res [][]int
	var tmpRes []int
	var backtrack func(idx int)
	backtrack = func(idx int) {
		for i := idx; i <= n; i++ {
			tmpRes = append(tmpRes, i)
			if len(tmpRes) == k {
				res = append(res, append([]int{}, tmpRes...))
			} else {
				backtrack(i + 1)
			}
			tmpRes = tmpRes[:len(tmpRes)-1]
		}
	}
	backtrack(1)
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		n    int
		k    int
		want [][]int
	}{
		{4, 2, [][]int{{2, 4}, {3, 4}, {2, 3}, {1, 2}, {1, 3}, {1, 4}}},
		{1, 1, [][]int{{1}}},
	}
	sorted := func(input [][]int) [][]int {
		sort.Slice(input, func(i, j int) bool {
			if input[i][0] != input[j][0] {
				return input[i][0] < input[j][0]
			}
			return input[i][1] < input[j][1]
		})
		return input
	}
	for _, tt := range tc {
		assert.Equal(t, sorted(tt.want), sorted(combine(tt.n, tt.k)))
	}
}
