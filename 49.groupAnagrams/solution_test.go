package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] Group Anagrams
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	l := make(map[string][]string)
	var sb strings.Builder
	for _, str := range strs {
		tmp := [26]int{}
		for _, c := range str {
			tmp[c-'a']++
		}
		for i, v := range tmp {
			if v != 0 {
				sb.WriteRune(rune(i))
				sb.WriteRune(rune(v))
			}
		}
		key := sb.String()
		l[key] = append(l[key], str)
		sb.Reset()
	}
	res := make([][]string, 0, len(l))
	for _, strs := range l {
		res = append(res, strs)
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  []string
		output [][]string
	}{
		{
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
		{[]string{""}, [][]string{{""}}},
		{[]string{"a"}, [][]string{{"a"}}},
	}
	for i := range tc {
		assert.Equal(t, tc[i].output, groupAnagrams(tc[i].input))
	}
}
