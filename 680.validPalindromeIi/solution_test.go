package main

import (
	"testing"
)

/*
 * @lc app=leetcode.cn id=680 lang=golang
 * @lcpr version=30002
 *
 * [680] 验证回文串 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func validPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] == s[right] {
			left++
			right--
		} else {
			flag1, flag2 := true, true
			for newLeft, newRight := left, right-1; newLeft < newRight; newLeft, newRight = newLeft+1, newRight-1 {
				if s[newLeft] != s[newRight] {
					flag1 = false
					break
				}
			}
			for newLeft, newRight := left+1, right; newLeft < newRight; newLeft, newRight = newLeft+1, newRight-1 {
				if s[newLeft] != s[newRight] {
					flag2 = false
					break
				}
			}
			return flag1 || flag2
		}
	}
	return true
}

// @lc code=end

/*
// @lcpr case=start
// "aba"\n
// @lcpr case=end

// @lcpr case=start
// "abca"\n
// @lcpr case=end

// @lcpr case=start
// "abc"\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect bool
	}{
		{"already palindrome", "aba", true},
		{"palindrome after removal", "abca", true},
		{"non-palindrome", "abc", false},
		{"empty string", "", true},
		{"single character", "a", true},
		{"two identical characters", "aa", true},
		{"two different characters", "ab", true}, // Remove one char to get a palindrome
		{"longer palindrome", "racecar", true},
		{"remove first char", "ebacab", true},   // remove 'e' at beginning gives "bacab"
		{"remove middle char", "abeceba", true}, // remove 'c' gives "abeeba"
		{"complex non-palindrome", "abcdef", false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := validPalindrome(tc.input); got != tc.expect {
				t.Errorf("validPalindrome(%q) = %v, want %v", tc.input, got, tc.expect)
			}
		})
	}
}
