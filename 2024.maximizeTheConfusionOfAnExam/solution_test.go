package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=2024 lang=golang
 *
 * [2024] Maximize the Confusion of an Exam
 */

// @lc code=start

func key(r byte) int {
	if r == 'T' {
		return 1
	}
	return 0
}

func maxConsecutiveAnswers(answerKey string, k int) int {
	var cnt [2]int
	var left, res int
	for right := 0; right < len(answerKey); right++ {
		cnt[key(answerKey[right])]++
		for !(cnt[0] <= k || cnt[1] <= k) {
			cnt[key(answerKey[left])]--
			left++
		}
		res = max(res, right-left+1)
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		answerKey string
		k         int
		expect    int
	}{
		{"TTFF", 2, 4},
		{"TFFT", 1, 3},
		{"TTFTTFTT", 1, 5},
	}
	for _, tt := range tc {
		assert.Equal(t, maxConsecutiveAnswers(tt.answerKey, tt.k), tt.expect)
	}
}
