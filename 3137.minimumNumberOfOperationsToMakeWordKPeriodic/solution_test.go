package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=3137 lang=golang
 *
 * [3137] Minimum Number of Operations to Make Word K-Periodic
 */

// @lc code=start
func minimumOperationsToMakeKPeriodic(word string, k int) int {
	m := make(map[string]int)
	var maxCnt int
	for i := 0; i != len(word); i += k {
		subStr := word[i : i+k]
		m[subStr]++
		maxCnt = max(maxCnt, m[subStr])
	}
	return len(word)/k - maxCnt
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		word string
		k    int
		want int
	}{
		{word: "leetcodeleet", k: 4, want: 1},
		{word: "leetcoleet", k: 2, want: 3},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, minimumOperationsToMakeKPeriodic(tt.word, tt.k))
	}
}
