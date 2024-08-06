package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] Triangle
 */

// @lc code=start
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	var dp []int
	for i, row := range triangle {
		if i == 0 {
			dp = row
			continue
		}
		new := make([]int, len(row))
		for i := range new {
			new[i] = math.MaxInt
		}
		for j, item := range dp {
			new[j] = min(new[j], row[j]+item)
			new[j+1] = min(new[j+1], row[j+1]+item)
		}
		dp = new
	}
	res := math.MaxInt
	for _, v := range dp {
		res = min(res, v)
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  [][]int
		output int
	}{
		{[][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}, 11},
		{[][]int{{-10}}, -10},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, minimumTotal(tt.input))
	}
}
