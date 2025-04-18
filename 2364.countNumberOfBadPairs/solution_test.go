/*
 * @lc app=leetcode.cn id=2364 lang=golang
 * @lcpr version=30104
 *
 * [2364] 统计坏数对的数目
 */

package main

import (
	"testing"
)

// @lc code=start
func countBadPairs(nums []int) int64 {
	var notRes int64
	m := make(map[int]int)
	for idx, num := range nums {
		notRes += int64(m[num-idx])
		m[num-idx]++
	}
	return int64(len(nums)*(len(nums)-1)/2) - notRes
}

// @lc code=end

func Test(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int64
	}{
		{[]int{4, 1, 3, 3}, 5},
		{[]int{1, 2, 3, 4, 5}, 0},
	}

	for _, test := range tests {
		result := countBadPairs(test.nums)
		if result != test.expected {
			t.Errorf("For nums %v, expected %d but got %d", test.nums, test.expected, result)
		}
	}
}

/*
// @lcpr case=start
// [4,1,3,3]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5]\n
// @lcpr case=end

*/
