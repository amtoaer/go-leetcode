/*
 * @lc app=leetcode.cn id=249 lang=golang
 * @lcpr version=30202
 *
 * [249] 移位字符串分组
 */

package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// @lc code=start
func groupStrings(strings []string) [][]string {
	var (
		latest int
		res    [][]string
		group  = make(map[string]int)
	)
	for _, s := range strings {
		key := hash(s)
		if idx, ok := group[key]; ok {
			res[idx] = append(res[idx], s)
		} else {
			res = append(res, []string{s})
			group[key] = latest
			latest++
		}
	}
	return res
}

func hash(s string) string {
	var sb strings.Builder
	prev := s[0]
	for i := range len(s) {
		diff := int(s[i]) - int(prev)
		if diff < 0 {
			diff += 26
		}
		sb.WriteByte('a' + byte(diff))
		prev = s[i]
	}
	return sb.String()
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		strings []string
		want    [][]string
	}{
		{
			[]string{"abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"},
			[][]string{{"abc", "bcd", "xyz"}, {"acef"}, {"az", "ba"}, {"a", "z"}},
		},
		{
			[]string{"a"},
			[][]string{{"a"}},
		},
	}
	for _, tt := range tc {
		result := groupStrings(tt.strings)
		assert.Equal(t, result, tt.want)
	}
}

/*
// @lcpr case=start
// ["abc","bcd","acef","xyz","az","ba","a","z"]\n
// @lcpr case=end

// @lcpr case=start
// ["a"]\n
// @lcpr case=end

*/
