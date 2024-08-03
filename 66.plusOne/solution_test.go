package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] Plus One
 */

// @lc code=start
func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += carry
		carry = 0
		if digits[i] < 10 {
			break
		}
		digits[i] -= 10
		carry = 1
	}
	if carry != 0 {
		return append([]int{1}, digits...)
	}
	return digits
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  []int
		output []int
	}{
		{input: []int{1, 2, 3}, output: []int{1, 2, 4}},
		{input: []int{4, 3, 2, 1}, output: []int{4, 3, 2, 2}},
		{input: []int{0}, output: []int{1}},
		{input: []int{9}, output: []int{1, 0}},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, plusOne(tt.input))
	}
}
