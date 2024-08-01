package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] Valid Palindrome
 */

// @lc code=start
func isPalindrome(s string) bool {
	isOk := func(b byte) bool {
		return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9')
	}
	toLower := func(b byte) byte {
		if b >= 'A' && b <= 'Z' {
			return b + 32
		}
		return b
	}
	left, right := 0, len(s)-1
	for {
		for ; left < right && !isOk(s[left]); left++ {
		}
		for ; left < right && !isOk(s[right]); right-- {
		}
		if left >= right {
			break
		}
		if toLower(s[left]) != toLower(s[right]) {
			return false
		}
		left++
		right--
	}
	return true
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  string
		output bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true},
		{"0P", false},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, isPalindrome(tt.input))
	}
}
