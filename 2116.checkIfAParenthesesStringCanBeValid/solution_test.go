/*
 * @lc app=leetcode.cn id=2116 lang=golang
 * @lcpr version=30104
 *
 * [2116] 判断一个括号字符串是否有效
 */

package main

import (
	"testing"
)

// @lc code=start
func canBeValid(s string, locked string) bool {
	var minVal, maxVal int
	for i := 0; i < len(s); i++ {
		if locked[i] == '1' {
			d := diff(s[i])
			maxVal += d
			minVal = max(minVal+d, (i+1)%2)
		} else {
			maxVal++
			minVal = max(minVal-1, (i+1)%2)
		}
		if maxVal < minVal {
			return false
		}
	}
	return minVal == 0
}

func diff(b byte) int {
	if b == '(' {
		return 1
	}
	return -1
}

// @lc code=end

func Test(t *testing.T) {
	testCases := []struct {
		s      string
		locked string
		want   bool
	}{
		{
			s:      "))()))",
			locked: "010100",
			want:   true,
		},
		{
			s:      "()()",
			locked: "0000",
			want:   true,
		},
		{
			s:      ")",
			locked: "0",
			want:   false,
		},
		{
			s:      "(((())(((())",
			locked: "111111010111",
			want:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.s+"-"+tc.locked, func(t *testing.T) {
			got := canBeValid(tc.s, tc.locked)
			if got != tc.want {
				t.Errorf("canBeValid(%q, %q) = %v, want %v", tc.s, tc.locked, got, tc.want)
			}
		})
	}
}

/*
// @lcpr case=start
// "))()))"\n"010100"\n
// @lcpr case=end

// @lcpr case=start
// "()()"\n"0000"\n
// @lcpr case=end

// @lcpr case=start
// ")"\n"0"\n
// @lcpr case=end

// @lcpr case=start
// "(((())(((())"\n"111111010111"\n
// @lcpr case=end

*/
