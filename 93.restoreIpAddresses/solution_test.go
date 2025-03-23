/*
 * @lc app=leetcode.cn id=93 lang=golang
 * @lcpr version=30104
 *
 * [93] 复原 IP 地址
 */

package main

import (
	"strings"
	"testing"
)

// @lc code=start
func restoreIpAddresses(s string) []string {
	var res []string
	var helper func(s string, items []string, idx int)
	helper = func(s string, items []string, idx int) {
		if len(items) == 4 && idx == len(s) {
			res = append(res, strings.Join(items, "."))
		}
		if len(items) == 4 || idx == len(s) {
			return
		}
		cur := 0
		cnt := 3
		if s[idx] == '0' {
			cnt = 1
		}
		for i := idx; i < min(len(s), idx+cnt); i++ {
			cur = cur*10 + int(s[i]-'0')
			if cur >= 0 && cur <= 0xff {
				helper(s, append(items, s[idx:i+1]), i+1)
			} else {
				return
			}
		}
	}
	helper(s, []string{}, 0)
	return res
}

// @lc code=end

func Test(t *testing.T) {
	testCases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "25525511135",
			expected: []string{"255.255.11.135", "255.255.111.35"},
		},
		{
			input:    "0000",
			expected: []string{"0.0.0.0"},
		},
		{
			input:    "101023",
			expected: []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"},
		},
	}

	for _, tc := range testCases {
		t.Run("input="+tc.input, func(t *testing.T) {
			result := restoreIpAddresses(tc.input)
			if len(result) != len(tc.expected) {
				t.Errorf("Expected %d results, got %d for input %s", len(tc.expected), len(result), tc.input)
				t.Errorf("Expected: %v, Got: %v", tc.expected, result)
				return
			}

			// Create map for easier comparison since order may differ
			expectedMap := make(map[string]bool)
			for _, ip := range tc.expected {
				expectedMap[ip] = true
			}

			for _, ip := range result {
				if !expectedMap[ip] {
					t.Errorf("Unexpected IP address: %s", ip)
				}
			}
		})
	}
}

/*
// @lcpr case=start
// "25525511135"\n
// @lcpr case=end

// @lcpr case=start
// "0000"\n
// @lcpr case=end

// @lcpr case=start
// "101023"\n
// @lcpr case=end

*/
