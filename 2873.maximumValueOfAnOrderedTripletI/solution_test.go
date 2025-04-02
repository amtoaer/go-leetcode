/*
 * @lc app=leetcode.cn id=2873 lang=golang
 * @lcpr version=30104
 *
 * [2873] 有序三元组中的最大值 I
 */

package main

import (
	"testing"
)

// @lc code=start
func maximumTripletValue(nums []int) int64 {
	n := len(nums)
	var res, imax, dmax int64 = 0, 0, 0
	for k := 0; k < n; k++ {
		res = max(res, dmax*int64(nums[k]))
		dmax = max(dmax, imax-int64(nums[k]))
		imax = max(imax, int64(nums[k]))
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected int64
	}{
		{[]int{12, 6, 1, 2, 7}, 77},
		{[]int{1, 10, 3, 4, 19}, 133},
		{[]int{1, 2, 3}, 0},
	}

	for _, tc := range testCases {
		result := maximumTripletValue(tc.nums)
		if result != tc.expected {
			t.Errorf("For nums=%v, expected %d but got %d", tc.nums, tc.expected, result)
		}
	}
}

/*
// @lcpr case=start
// [12,6,1,2,7]\n
// @lcpr case=end

// @lcpr case=start
// [1,10,3,4,19]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

*/
