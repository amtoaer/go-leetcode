package main

import (
	"testing"
)

/*
 * @lc app=leetcode.cn id=2506 lang=golang
 * @lcpr version=30005
 *
 * [2506] 统计相似字符串对的数目
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func similarPairs(words []string) int {
	m := make(map[[26]bool]int)
	for _, word := range words {
		m[hash(word)]++
	}
	var res int
	for _, v := range m {
		res += v * (v - 1) / 2
	}
	return res
}

func hash(word string) [26]bool {
	var helper [26]bool
	for _, w := range word {
		helper[int(w-'a')] = true
	}
	return helper
}

// @lc code=end

/*
// @lcpr case=start
// ["aba","aabb","abcd","bac","aabc"]\n
// @lcpr case=end

// @lcpr case=start
// ["aabb","ab","ba"]\n
// @lcpr case=end

// @lcpr case=start
// ["nba","cba","dba"]\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{
			name:  "Test Case 1",
			input: []string{"aba", "aabb", "abcd", "bac", "aabc"},
			want:  2,
		},
		{
			name:  "Test Case 2",
			input: []string{"aabb", "ab", "ba"},
			want:  3,
		},
		{
			name:  "Test Case 3",
			input: []string{"nba", "cba", "dba"},
			want:  0,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := similarPairs(tt.input); got != tt.want {
				t.Errorf("similarPairs(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}
