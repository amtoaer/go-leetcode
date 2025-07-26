/*
 * @lc app=leetcode.cn id=734 lang=golang
 * @lcpr version=30202
 *
 * [734] 句子相似性
 */

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// @lc code=start
func areSentencesSimilar(sentence1 []string, sentence2 []string, similarPairs [][]string) bool {
	if len(sentence1) != len(sentence2) {
		return false
	}
	m := make(map[string]struct{})
	for _, pair := range similarPairs {
		m[pair[0]+"/"+pair[1]] = struct{}{}
	}
	for i, word := range sentence1 {
		_, ok1 := m[word+"/"+sentence2[i]]
		_, ok2 := m[sentence2[i]+"/"+word]
		if !(word == sentence2[i] || ok1 || ok2) {
			return false
		}
	}
	return true
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		sentence1    []string
		sentence2    []string
		similarPairs [][]string
		want         bool
	}{
		{[]string{"great", "acting", "skills"}, []string{"fine", "drama", "talent"}, [][]string{{"great", "fine"}, {"drama", "acting"}, {"skills", "talent"}}, true},
		{[]string{"great"}, []string{"great"}, [][]string{}, true},
		{[]string{"great"}, []string{"doubleplus", "good"}, [][]string{{"great", "doubleplus"}}, false},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, areSentencesSimilar(tt.sentence1, tt.sentence2, tt.similarPairs))
	}
}

/*
// @lcpr case=start
// ["great","acting","skills"]\n["fine","drama","talent"]\n\n[["great","fine"],["drama","acting"],["skills","talent"]]\n
// @lcpr case=end

// @lcpr case=start
// ["great"]\n["great"]\n[]\n
// @lcpr case=end

// @lcpr case=start
// ["great"]\n["doubleplus","good"]\n[["great","doubleplus"]]\n
// @lcpr case=end

*/
