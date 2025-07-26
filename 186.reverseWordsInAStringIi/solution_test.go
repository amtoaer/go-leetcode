/*
 * @lc app=leetcode.cn id=186 lang=golang
 * @lcpr version=30202
 *
 * [186] 反转字符串中的单词 II
 */

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// @lc code=start
func reverseWords(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	start := 0
	for end := range s {
		if end == len(s)-1 || s[end+1] == ' ' {
			for i, j := start, end; i < j; i, j = i+1, j-1 {
				s[i], s[j] = s[j], s[i]
			}
			start = end + 2
		}
	}
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		s    []byte
		want []byte
	}{
		{[]byte{'t', 'h', 'e', ' ', 's', 'k', 'y', ' ', 'i', 's', ' ', 'b', 'l', 'u', 'e'}, []byte{'b', 'l', 'u', 'e', ' ', 'i', 's', ' ', 's', 'k', 'y', ' ', 't', 'h', 'e'}},
		{[]byte{'a'}, []byte{'a'}},
	}
	for _, tt := range tc {
		input := make([]byte, len(tt.s))
		copy(input, tt.s)
		reverseWords(input)
		assert.Equal(t, tt.want, input)
	}
}

/*
// @lcpr case=start
// ["t","h","e"," ","s","k","y"," ","i","s"," ","b","l","u","e"]\n
// @lcpr case=end

// @lcpr case=start
// ["a"]\n
// @lcpr case=end

*/
