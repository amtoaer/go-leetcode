package main

import (
	"math"
	"testing"
)

/*
 * @lc app=leetcode.cn id=633 lang=golang
 * @lcpr version=20002
 *
 * [633] 平方数之和
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func judgeSquareSum(c int) bool {
	i, j := 0, int(math.Sqrt(float64(c)))
	for i <= j {
		tmp := i*i + j*j
		if tmp == c {
			return true
		} else if tmp > c {
			j--
		} else {
			i++
		}
	}
	return false
}

// @lc code=end

/*
// @lcpr case=start
// 5\n
// @lcpr case=end

// @lcpr case=start
// 3\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{5, true},
		{3, false},
		{4, true},
		{2, true},
		{1, true},
		{0, true},
		{10, true},
		{25, true},
		{26, true},
		{100, true},
		{101, true},
	}

	for _, test := range tests {
		if result := judgeSquareSum(test.input); result != test.expected {
			t.Errorf("For input %d, expected %v, but got %v", test.input, test.expected, result)
		}
	}
}
