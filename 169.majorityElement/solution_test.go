package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] Majority Element
 */

// @lc code=start
func majorityElement(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
		if m[num] > len(nums)/2 {
			return num
		}
	}
	return -1
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  []int
		output int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, majorityElement(tt.input))
	}
}
