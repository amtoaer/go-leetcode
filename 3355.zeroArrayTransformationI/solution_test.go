/*
 * @lc app=leetcode.cn id=3355 lang=golang
 * @lcpr version=30201
 *
 * [3355] 零数组变换 I
 */

package main

import (
	"testing"
)

// @lc code=start
func isZeroArray(nums []int, queries [][]int) bool {
	diffArray := make([]int, len(nums)+1)
	for _, query := range queries {
		left, right := query[0], query[1]
		diffArray[left]++
		diffArray[right+1]--
	}
	tmpVal := 0
	for i := range len(nums) {
		tmpVal += diffArray[i]
		if tmpVal < nums[i] {
			return false
		}
	}
	return true
}

// @lc code=end

func Test(t *testing.T) {
	testCases := []struct {
		nums    []int
		queries [][]int
		want    bool
	}{
		{[]int{1, 0, 1}, [][]int{{0, 2}}, true},
		{[]int{4, 3, 2, 1}, [][]int{{1, 3}, {0, 2}}, false},
	}

	for i, tc := range testCases {
		got := isZeroArray(tc.nums, tc.queries)
		if got != tc.want {
			t.Errorf("Test case %d: isZeroArray(%v, %v) = %v, want %v", i, tc.nums, tc.queries, got, tc.want)
		}
	}
}

/*
// @lcpr case=start
// [1,0,1]\n[[0,2]]\n
// @lcpr case=end

// @lcpr case=start
// [4,3,2,1]\n[[1,3],[0,2]]\n
// @lcpr case=end

*/
