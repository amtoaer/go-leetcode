package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=2951 lang=golang
 *
 * [2951] Find the Peaks
 */

// @lc code=start
func findPeaks(mountain []int) []int {
	result := []int{}
	for i := 1; i < len(mountain)-1; i++ {
		if mountain[i] > mountain[i-1] && mountain[i] > mountain[i+1] {
			result = append(result, i)
		}
	}
	return result
}

// @lc code=end

func Test(t *testing.T) {
	testcases := [][2][]int{
		{
			{2, 4, 4},
			{},
		},
		{
			{1, 4, 3, 8, 5},
			{1, 3},
		},
	}
	for _, testcase := range testcases {
		assert.Equal(t, testcase[1], findPeaks(testcase[0]))
	}
}
