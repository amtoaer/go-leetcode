package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] Length of Last Word
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	var count int
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if count > 0 {
				return count
			}
		} else {
			count++
		}
	}
	return count
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(t, 5, lengthOfLastWord("Hello World"))
	assert.Equal(t, 0, lengthOfLastWord(" "))
	assert.Equal(t, 1, lengthOfLastWord("a"))
	assert.Equal(t, 1, lengthOfLastWord("a "))
}
