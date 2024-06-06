package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=2938 lang=golang
 *
 * [2938] Separate Black and White Balls
 */

// @lc code=start
func minimumSteps(s string) int64 {
	left, right := 0, len(s)-1
	var res int
	for left < right {
		if s[left] == '1' && s[right] == '0' {
			res += right - left
			left++
			right--
			continue
		}
		if s[left] != '1' {
			left++
			continue
		}
		if s[right] != '0' {
			right--
			continue
		}
	}
	return int64(res)
}

// @lc code=end

func Test(t *testing.T) {
	testcases := []struct {
		s      string
		output int64
	}{
		{"101", 1},
		{"100", 2},
		{"0111", 0},
	}
	for _, testcase := range testcases {
		assert.Equal(t, testcase.output, minimumSteps(testcase.s))
	}
}
