/*
 * @lc app=leetcode.cn id=3330 lang=golang
 * @lcpr version=30201
 *
 * [3330] 找到初始输入字符串 I
 */

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// @lc code=start
func possibleStringCount(word string) int {
	res := 1
	for i := range word {
		if i == 0 {
			continue
		}
		if word[i-1] == word[i] {
			res++
		}
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		word string
		want int
	}{
		{"abbcccc", 5},
		{"abcd", 1},
		{"aaaa", 4},
	}

	for _, tt := range tc {
		t.Run(tt.word, func(t *testing.T) {
			assert.Equal(t, tt.want, possibleStringCount(tt.word))
		})
	}
}

/*
// @lcpr case=start
// "abbcccc"\n
// @lcpr case=end

// @lcpr case=start
// "abcd"\n
// @lcpr case=end

// @lcpr case=start
// "aaaa"\n
// @lcpr case=end

*/
