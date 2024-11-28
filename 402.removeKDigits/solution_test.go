package main

import (
	"testing"
)

/*
 * @lc app=leetcode.cn id=402 lang=golang
 * @lcpr version=20003
 *
 * [402] 移掉 K 位数字
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func removeKdigits(num string, k int) string {
	var nums []rune
	for _, r := range num {
		for len(nums) != 0 && nums[len(nums)-1] > r && k > 0 {
			k--
			nums = nums[:len(nums)-1]
		}
		if len(nums) == 0 && r == '0' {
			continue
		}
		nums = append(nums, r)
	}
	if k >= len(nums) {
		return "0"
	}
	if k > 0 {
		nums = nums[:len(nums)-k]
	}
	return string(nums)
}

// @lc code=end

/*
// @lcpr case=start
// "1432219"\n3\n
// @lcpr case=end

// @lcpr case=start
// "10200"\n1\n
// @lcpr case=end

// @lcpr case=start
// "10"\n2\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		num      string
		k        int
		expected string
	}{
		{"1432219", 3, "1219"},
		{"10200", 1, "200"},
		{"10", 2, "0"},
		{"1234567890", 9, "0"},
		{"112", 1, "11"},
		{"9", 1, "0"},
	}

	for _, test := range tests {
		result := removeKdigits(test.num, test.k)
		if result != test.expected {
			t.Errorf("removeKdigits(%v, %v) = %v; expected %v", test.num, test.k, result, test.expected)
		}
	}
}
