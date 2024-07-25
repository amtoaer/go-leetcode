package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=2844 lang=golang
 *
 * [2844] Minimum Operations to Make a Special Number
 */

// @lc code=start
func minimumOperations(num string) int {
	has_0, has_5 := false, false
	n := len(num)
	for i := len(num) - 1; i >= 0; i-- {
		if num[i] == '0' || num[i] == '5' {
			if has_0 {
				return n - i - 2
			}
			if num[i] == '0' {
				has_0 = true
			} else {
				has_5 = true
			}
		} else if num[i] == '2' || num[i] == '7' {
			if has_5 {
				return n - i - 2
			}
		}
	}
	if has_0 {
		return n - 1
	}
	return n
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  string
		output int
	}{
		{"2245047", 2},
		{"2908305", 3},
		{"10", 1},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, minimumOperations(tt.input))
	}
}
