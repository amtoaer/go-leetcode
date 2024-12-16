package main

import (
	"strings"
	"testing"
)

/*
 * @lc app=leetcode.cn id=394 lang=golang
 * @lcpr version=20004
 *
 * [394] 字符串解码
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func decodeString(s string) string {
	var (
		nums []int
		strs []string
		i    int
		sb   strings.Builder
	)
	for i < len(s) {
		if s[i] >= '0' && s[i] <= '9' {
			num := 0
			for i < len(s) && s[i] >= '0' && s[i] <= '9' {
				num = num*10 + int(s[i]-'0')
				i++
			}
			nums = append(nums, num)
		} else if s[i] >= 'a' && s[i] <= 'z' {
			start := i
			for i < len(s) && s[i] >= 'a' && s[i] <= 'z' {
				i++
			}
			strs = append(strs, string(s[start:i]))
		} else if s[i] == '[' {
			strs = append(strs, "[")
			i++
		} else {
			left := len(strs) - 1
			for strs[left] != "[" {
				left--
			}
			for _, str := range strs[left+1:] {
				sb.WriteString(str)
			}
			num := nums[len(nums)-1]
			strs = strs[:left]
			nums = nums[:len(nums)-1]
			strs = append(strs, strings.Repeat(sb.String(), num))
			sb.Reset()
			i++
		}
	}
	for _, str := range strs {
		sb.WriteString(str)
	}
	return sb.String()
}

// @lc code=end

/*
// @lcpr case=start
// "3[a]2[bc]"\n
// @lcpr case=end

// @lcpr case=start
// "3[a2[c]]"\n
// @lcpr case=end

// @lcpr case=start
// "2[abc]3[cd]ef"\n
// @lcpr case=end

// @lcpr case=start
// "abc3[cd]xyz"\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"3[a]2[bc]", "aaabcbc"},
		{"3[a2[c]]", "accaccacc"},
		{"2[abc]3[cd]ef", "abcabccdcdcdef"},
		{"abc3[cd]xyz", "abccdcdcdxyz"},
		{"10[a]", "aaaaaaaaaa"},
		{"2[3[a]b]", "aaabaaab"},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := decodeString(test.input)
			if result != test.expected {
				t.Errorf("expected %s, got %s", test.expected, result)
			}
		})
	}
}
