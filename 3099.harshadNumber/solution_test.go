package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=3099 lang=golang
 *
 * [3099] Harshad Number
 */

// @lc code=start
func sumOfTheDigitsOfHarshadNumber(x int) int {
	cur, sum := x, 0
	for cur > 0 {
		sum += cur % 10
		cur /= 10
	}
	if x%sum == 0 {
		return sum
	}
	return -1
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		x   int
		ans int
	}{
		{18, 9},
		{23, -1},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.ans, sumOfTheDigitsOfHarshadNumber(tt.x))
	}
}
