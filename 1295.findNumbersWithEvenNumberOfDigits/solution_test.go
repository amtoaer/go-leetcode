/*
 * @lc app=leetcode.cn id=1295 lang=golang
 * @lcpr version=30200
 *
 * [1295] 统计位数为偶数的数字
 */

package main

import (
	"testing"
)

// @lc code=start
func findNumbers(nums []int) int {
	var res int
	for _, num := range nums {
		var cnt int
		for num > 0 {
			num /= 10
			cnt++
		}
		if cnt&1 == 0 {
			res++
		}
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	type testCase struct {
		nums     []int
		expected int
	}

	cases := []testCase{
		{nums: []int{12, 345, 2, 6, 7896}, expected: 2},
		{nums: []int{555, 901, 482, 1771}, expected: 1},
	}

	for _, c := range cases {
		result := findNumbers(c.nums)
		if result != c.expected {
			t.Errorf("findNumbers(%v) = %d, expected %d", c.nums, result, c.expected)
		}
	}
}

/*
// @lcpr case=start
// [12,345,2,6,7896]\n
// @lcpr case=end

// @lcpr case=start
// [555,901,482,1771]\n
// @lcpr case=end

*/
