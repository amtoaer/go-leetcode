package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=2544 lang=golang
 *
 * [2544] Alternating Digit Sum
 */

// @lc code=start
func alternateDigitSum(n int) int {
	flag, sum := 1, 0
	for n > 0 {
		sum += n % 10 * flag
		n /= 10
		flag = -flag
	}
	return -sum * flag
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(t, 0, alternateDigitSum(0))
	assert.Equal(t, 1, alternateDigitSum(1))
	assert.Equal(t, 1, alternateDigitSum(10))
	assert.Equal(t, 2, alternateDigitSum(101))
}
