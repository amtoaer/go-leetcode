package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] String to Integer (atoi)
 */

// @lc code=start
func myAtoi(s string) int {
	var (
		max                 = 1<<31 - 1
		min                 = -1 << 31
		currentIdx          = 0
		gotSign             = false
		isLeadingWhiteSpace = true
		sign                = 1
		result              = 0
	)
	for currentIdx < len(s) {
		if s[currentIdx] == ' ' && isLeadingWhiteSpace {
			// 跳过前缀空格
			currentIdx += 1
			continue
		}
		isLeadingWhiteSpace = false
		if !gotSign && (s[currentIdx] == '-' || s[currentIdx] == '+') {
			// 读取符号
			gotSign = true
			if s[currentIdx] == '-' {
				sign = -1
			}
			currentIdx += 1
			continue
		}
		if int(s[currentIdx]-'0') >= 0 && int(s[currentIdx]-'0') <= 9 {
			// 已经有数字了，也认为是读过了符号
			gotSign = true
			result = result*10 + int(s[currentIdx]-'0')
			if result >= max && sign == 1 {
				// 超过最大值
				return max
			}
			if result >= -min && sign == -1 {
				// 超过最小值
				return min
			}
			currentIdx += 1
			continue
		}
		// 遇到非数字，直接返回
		return result * sign
	}
	return result * sign
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(t, myAtoi("42"), 42)
	assert.Equal(t, myAtoi("   -42"), -42)
	assert.Equal(t, myAtoi("4193 with words"), 4193)
	assert.Equal(t, myAtoi("words and 987"), 0)
	assert.Equal(t, myAtoi("-91283472332"), -2147483648)
	assert.Equal(t, myAtoi("18446744073709551617"), 2147483647)
}
