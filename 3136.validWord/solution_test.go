/*
 * @lc app=leetcode.cn id=3136 lang=golang
 * @lcpr version=30201
 *
 * [3136] 有效单词
 */

package main

import (
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

// @lc code=start
func isValid(word string) bool {
	if len(word) < 3 {
		return false
	}
	aeiou := map[rune]any{
		'a': struct{}{},
		'e': struct{}{},
		'i': struct{}{},
		'o': struct{}{},
		'u': struct{}{},
	}
	var hasAEIOU, hasNotAEIOU bool
	for _, r := range word {
		if unicode.IsLetter(r) {
			if _, ok := aeiou[unicode.ToLower(r)]; ok {
				hasAEIOU = true
			} else {
				hasNotAEIOU = true
			}
		} else if !unicode.IsNumber(r) {
			return false
		}
	}
	return hasAEIOU && hasNotAEIOU
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		word string
		want bool
	}{
		{"234Adas", true},
		{"b3", false},
		{"a3$e", false},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, isValid(tt.word))
	}
}

/*
// @lcpr case=start
// "234Adas"\n
// @lcpr case=end

// @lcpr case=start
// "b3"\n
// @lcpr case=end

// @lcpr case=start
// "a3$e"\n
// @lcpr case=end

*/
