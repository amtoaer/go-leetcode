package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */

// @lc code=start
func letterCombinations(digits string) []string {
	letterMap := map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}
	var bf bytes.Buffer
	res := make([]string, 0)
	var backtrack func(int)
	backtrack = func(idx int) {
		for _, choice := range letterMap[digits[idx]] {
			bf.WriteByte(choice)
			if idx == len(digits)-1 {
				res = append(res, bf.String())
			} else {
				backtrack(idx + 1)
			}
			bf.Truncate(bf.Len() - 1)
		}
	}
	if len(digits) > 0 {
		backtrack(0)
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(
		t,
		[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		letterCombinations("23"),
	)
	assert.Equal(t, []string{}, letterCombinations(""))
	assert.Equal(t, []string{"a", "b", "c"}, letterCombinations("2"))
}
