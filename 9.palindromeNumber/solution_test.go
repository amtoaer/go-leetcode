package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var (
		tmp     = x
		reverse = 0
	)
	for tmp > 0 {
		reverse = reverse*10 + tmp%10
		tmp /= 10
	}
	return reverse == x
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(t, true, isPalindrome(121))
	assert.Equal(t, false, isPalindrome(-121))
	assert.Equal(t, false, isPalindrome(10))
}
