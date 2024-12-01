package main

import (
	"testing"
)

/*
 * @lc app=leetcode.cn id=165 lang=golang
 * @lcpr version=20004
 *
 * [165] 比较版本号
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func compareVersion(version1 string, version2 string) int {
	var cur1, cur2 int
	for cur1 < len(version1) || cur2 < len(version2) {
		var left, right int
		for cur1 < len(version1) && version1[cur1] != '.' {
			left = left*10 + int(version1[cur1]-'0')
			cur1++
		}
		for cur2 < len(version2) && version2[cur2] != '.' {
			right = right*10 + int(version2[cur2]-'0')
			cur2++
		}
		if left < right {
			return -1
		}
		if left > right {
			return 1
		}
		cur1++
		cur2++
	}
	return 0
}

// @lc code=end

/*
// @lcpr case=start
// "1.2"\n"1.10"\n
// @lcpr case=end

// @lcpr case=start
// "1.01"\n"1.001"\n
// @lcpr case=end

// @lcpr case=start
// "1.0"\n"1.0.0.0"\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		version1 string
		version2 string
		expected int
	}{
		{"1.2", "1.10", -1},
		{"1.01", "1.001", 0},
		{"1.0", "1.0.0.0", 0},
		{"0.1", "1.1", -1},
		{"1.0.1", "1", 1},
		{"7.5.2.4", "7.5.3", -1},
		{"1.01", "1.001", 0},
		{"1.0", "1.0.0", 0},
	}

	for _, test := range tests {
		result := compareVersion(test.version1, test.version2)
		if result != test.expected {
			t.Errorf("compareVersion(%s, %s) = %d; expected %d", test.version1, test.version2, result, test.expected)
		}
	}
}
