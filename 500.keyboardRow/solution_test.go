/*
 * @lc app=leetcode.cn id=500 lang=golang
 * @lcpr version=30201
 *
 * [500] 键盘行
 */

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// @lc code=start
func findWords(words []string) []string {
	var bytes [26]int
	for idx, slice := range []string{
		"qwertyuiop",
		"asdfghjkl",
		"zxcvbnm",
	} {
		for sIdx := range slice {
			bytes[slice[sIdx]-'a'] = idx
		}
	}
	var res []string
	for _, word := range words {
		label := bytes[lower(word[0])-'a']
		flag := true
		for idx := range word {
			if bytes[lower(word[idx])-'a'] != label {
				flag = false
				break
			}
		}
		if flag {
			res = append(res, word)
		}
	}
	return res
}

func lower(b byte) byte {
	if b >= 'a' && b <= 'z' {
		return b
	}
	return b + 32
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		words []string
		want  []string
	}{
		{[]string{"Hello", "Alaska", "Dad", "Peace"}, []string{"Alaska", "Dad"}},
		{[]string{"omk"}, []string(nil)},
		{[]string{"adsdf", "sfd"}, []string{"adsdf", "sfd"}},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, findWords(tt.words))
	}
}

/*
// @lcpr case=start
// ["Hello","Alaska","Dad","Peace"]\n
// @lcpr case=end

// @lcpr case=start
// ["omk"]\n
// @lcpr case=end

// @lcpr case=start
// ["adsdf","sfd"]\n
// @lcpr case=end

*/
