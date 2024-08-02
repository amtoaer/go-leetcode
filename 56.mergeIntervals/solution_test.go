package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] Merge Intervals
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{}
	var cur []int
	for idx, interval := range intervals {
		if idx == 0 {
			cur = interval
		}
		if interval[0] > cur[1] {
			res = append(res, cur)
			cur = interval
		} else {
			cur[1] = max(cur[1], interval[1])
		}
	}
	res = append(res, cur)
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  [][]int
		output [][]int
	}{
		{
			input:  [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			output: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{input: [][]int{{1, 4}, {4, 5}}, output: [][]int{{1, 5}}},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, merge(tt.input))
	}
}
