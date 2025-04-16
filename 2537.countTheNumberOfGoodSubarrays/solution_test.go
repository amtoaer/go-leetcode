/*
 * @lc app=leetcode.cn id=2537 lang=golang
 * @lcpr version=30104
 *
 * [2537] 统计好子数组的数目
 */

package main

import (
	"testing"
)

// @lc code=start
func countGood(nums []int, k int) int64 {
	var (
		l, r, pair int
		res        int64
		count      = make(map[int]int)
	)
	for r < len(nums) {
		pair += count[nums[r]]
		count[nums[r]]++
		r++
		for l < r && pair >= k {
			res += int64(len(nums) - r + 1)
			count[nums[l]]--
			pair -= count[nums[l]]
			l++
		}
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		k        int
		expected int64
	}{
		{
			name:     "example 1",
			nums:     []int{1, 1, 1, 1, 1},
			k:        10,
			expected: 1,
		},
		{
			name:     "example 2",
			nums:     []int{3, 1, 4, 3, 2, 2, 4},
			k:        2,
			expected: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := countGood(tc.nums, tc.k)
			if result != tc.expected {
				t.Errorf("countGood(%v, %d) = %d; want %d", tc.nums, tc.k, result, tc.expected)
			}
		})
	}
}

/*
// @lcpr case=start
// [1,1,1,1,1]\n10\n
// @lcpr case=end

// @lcpr case=start
// [3,1,4,3,2,2,4]\n2\n
// @lcpr case=end

*/
