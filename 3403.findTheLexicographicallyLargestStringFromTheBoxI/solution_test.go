/*
 * @lc app=leetcode.cn id=3403 lang=golang
 * @lcpr version=30201
 *
 * [3403] 从盒子中找出字典序最大的字符串 I
 */

package main

import (
	"testing"
)

// @lc code=start
func answerString(word string, numFriends int) string {
	if numFriends == 1 {
		return word
	}
	var bytes = []byte(word)
	numLength := len(bytes) - numFriends + 1
	maxBytes := bytes[:numLength]
	for start := 1; start < len(bytes); start++ {
		current := bytes[start:min(start+numLength, len(bytes))]
		if greater(current, maxBytes) {
			maxBytes = current
		}
	}
	return string(maxBytes)
}

func greater(a, b []byte) bool {
	for i := range min(len(a), len(b)) {
		if a[i] > b[i] {
			return true
		}
		if a[i] < b[i] {
			return false
		}
	}
	return len(a) > len(b)
}

// @lc code=end

func Test(t *testing.T) {
	tests := []struct {
		word       string
		numFriends int
		expected   string
	}{
		{"dbca", 2, "dbc"},
		{"gggg", 4, "g"},
		{"bif", 2, "if"},
		{"gh", 1, "gh"},
	}

	for _, tt := range tests {
		result := answerString(tt.word, tt.numFriends)
		if result != tt.expected {
			t.Errorf("answerString(%q, %d) = %q, want %q", tt.word, tt.numFriends, result, tt.expected)
		}
	}
}

/*
// @lcpr case=start
// "dbca"\n2\n
// @lcpr case=end

// @lcpr case=start
// "gggg"\n4\n
// @lcpr case=end

*/
