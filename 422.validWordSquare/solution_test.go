/*
 * @lc app=leetcode.cn id=422 lang=golang
 * @lcpr version=30202
 *
 * [422] 有效的单词方块
 */

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// @lc code=start
func validWordSquare(words []string) bool {
	for i, word := range words {
		for j := range word {
			if !(j < len(words) && i < len(words[j])) {
				return false
			}
			if words[j][i] != word[j] {
				return false
			}
		}
	}
	return true
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		words []string
		want  bool
	}{
		{[]string{"abcd", "bnrt", "crmy", "dtye"}, true},
		{[]string{"abcd", "bnrt", "crm", "dt"}, true},
		{[]string{"ball", "area", "read", "lady"}, false},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, validWordSquare(tt.words))
	}
}

/*
// @lcpr case=start
// ["abcd","bnrt","crmy","dtye"]\n
// @lcpr case=end

// @lcpr case=start
// ["abcd","bnrt","crm","dt"]\n
// @lcpr case=end

// @lcpr case=start
// ["ball","area","read","lady"]\n
// @lcpr case=end

*/
