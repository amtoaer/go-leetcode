package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] Reverse Integer
 */

// @lc code=start
func reverse(x int) int {
	var (
		max    = 1<<31 - 1
		min    = -1 << 31
		result = 0
	)
	for x != 0 {
		result = result*10 + x%10
		x /= 10
	}
	if result > max || result < min {
		return 0
	}
	return result
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(t, reverse(123), 321)
	assert.Equal(t, reverse(-123), -321)
	assert.Equal(t, reverse(120), 21)
}
