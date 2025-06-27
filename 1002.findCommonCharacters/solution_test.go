/*
 * @lc app=leetcode.cn id=1002 lang=golang
 * @lcpr version=30201
 *
 * [1002] 查找共用字符
 */

package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// @lc code=start
func commonChars(words []string) []string {
	var cnt [26]int
	for i := range cnt {
		cnt[i] = math.MaxInt
	}
	for _, word := range words {
		var tmpCnt [26]int
		for i := range word {
			tmpCnt[word[i]-'a']++
		}
		for i, v := range tmpCnt {
			cnt[i] = min(cnt[i], v)
		}
	}
	var res []string
	for i, v := range cnt {
		for range v {
			res = append(res, string('a'+byte(i)))
		}
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		words []string
		want  []string
	}{
		{[]string{"bella", "label", "roller"}, []string{"e", "l", "l"}},
		{[]string{"cool", "lock", "cook"}, []string{"c", "o"}},
	}
	for _, tt := range tc {
		assert.ElementsMatch(t, tt.want, commonChars(tt.words))
	}
}

/*
// @lcpr case=start
// ["bella","label","roller"]\n
// @lcpr case=end

// @lcpr case=start
// ["cool","lock","cook"]\n
// @lcpr case=end

*/
