/*
 * @lc app=leetcode.cn id=1165 lang=golang
 * @lcpr version=30202
 *
 * [1165] 单行键盘
 */

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// @lc code=start
func calculateTime(keyboard string, word string) int {
	var m [26]int
	for i := range len(keyboard) {
		m[keyboard[i]-'a'] = i
	}
	var res, prev int
	for i := range len(word) {
		newPos := m[word[i]-'a']
		res += abs(prev - newPos)
		prev = newPos
	}
	return res
}

func abs(i int) int {
	if i > 0 {
		return i
	}
	return -i
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		keyboard string
		word     string
		want     int
	}{
		{"abcdefghijklmnopqrstuvwxyz", "cba", 4},
		{"pqrstuvwxyzabcdefghijklmno", "leetcode", 73},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, calculateTime(tt.keyboard, tt.word))
	}
}

/*
// @lcpr case=start
// "abcdefghijklmnopqrstuvwxyz"\n"cba"\n
// @lcpr case=end

// @lcpr case=start
// "pqrstuvwxyzabcdefghijklmno"\n"leetcode"\n
// @lcpr case=end

*/
