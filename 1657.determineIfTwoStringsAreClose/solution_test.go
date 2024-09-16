package main

import (
	"reflect"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=1657 lang=golang
 *
 * [1657] Determine if Two Strings Are Close
 */

// @lc code=start
func closeStrings(word1 string, word2 string) bool {
	count1, count2 := make([]int, 26), make([]int, 26)
	for _, c := range word1 {
		count1[c-'a']++
	}
	for _, c := range word2 {
		count2[c-'a']++
	}
	for i := 0; i < 26; i++ {
		if count1[i] > 0 && count2[i] == 0 || count1[i] == 0 && count2[i] > 0 {
			return false
		}
	}
	sort.Ints(count1)
	sort.Ints(count2)
	return reflect.DeepEqual(count1, count2)
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		word1, word2 string
		want         bool
	}{
		{"abc", "bca", true},
		{"a", "aa", false},
		{"cabbba", "abbccc", true},
		{"uau", "ssx", false},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, closeStrings(tt.word1, tt.word2))
	}
}
