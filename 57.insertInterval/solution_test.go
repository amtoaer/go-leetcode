package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=57 lang=golang
 *
 * [57] Insert Interval
 */

// @lc code=start
func insert(intervals [][]int, newInterval []int) [][]int {
	var res [][]int
	appended := false
	for _, interval := range intervals {
		if interval[1] < newInterval[0] {
			res = append(res, interval)
		} else if interval[0] > newInterval[1] {
			if !appended {
				res = append(res, newInterval)
			}
			appended = true
			res = append(res, interval)
		} else {
			newInterval[0] = min(interval[0], newInterval[0])
			newInterval[1] = max(interval[1], newInterval[1])
		}
	}
	if !appended {
		res = append(res, newInterval)
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input       [][]int
		newInterval []int
		output      [][]int
	}{
		{
			input:       [][]int{{1, 3}, {6, 9}},
			newInterval: []int{2, 5},
			output:      [][]int{{1, 5}, {6, 9}},
		},
		{
			input:       [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			newInterval: []int{4, 8},
			output:      [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, insert(tt.input, tt.newInterval))
	}
}
