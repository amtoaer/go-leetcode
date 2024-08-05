package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] Move Zeroes
 */

// @lc code=start
func moveZeroes(nums []int) {
	x := 0
	for i, v := range nums {
		if v != 0 {
			if i != x {
				nums[x], nums[i] = nums[i], nums[x]
			}
			x++
		}
	}
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  []int
		output []int
	}{
		{input: []int{0, 1, 0, 3, 12}, output: []int{1, 3, 12, 0, 0}},
		{input: []int{0, 0, 0, 0, 0}, output: []int{0, 0, 0, 0, 0}},
		{input: []int{1, 2, 3, 4, 5}, output: []int{1, 2, 3, 4, 5}},
		{input: []int{0, 0, 0, 0, 1}, output: []int{1, 0, 0, 0, 0}},
		{input: []int{1, 0, 0, 0, 0}, output: []int{1, 0, 0, 0, 0}},
	}
	for _, tt := range tc {
		moveZeroes(tt.input)
		assert.True(t, reflect.DeepEqual(tt.input, tt.output))
	}
}
