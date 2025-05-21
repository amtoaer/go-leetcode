/*
 * @lc app=leetcode.cn id=3356 lang=golang
 * @lcpr version=30201
 *
 * [3356] 零数组变换 II
 */

package main

import (
	"testing"
)

// @lc code=start
func minZeroArray(nums []int, queries [][]int) int {
	diffArray := make([]int, len(nums)+1)
	var tmpVal, queryIdx int
	for i, num := range nums {
		tmpVal += diffArray[i]
		for tmpVal < num && queryIdx < len(queries) {
			query := queries[queryIdx]
			left, right, value := query[0], query[1], query[2]
			diffArray[left] += value
			diffArray[right+1] -= value
			if i >= left && i <= right {
				tmpVal += value
			}
			queryIdx++
		}
		if tmpVal < num {
			return -1
		}
	}
	return queryIdx
}

// @lc code=end

func Test(t *testing.T) {
	testCases := []struct {
		nums     []int
		queries  [][]int
		expected int
	}{
		{
			nums:     []int{2, 0, 2},
			queries:  [][]int{{0, 2, 1}, {0, 2, 1}, {1, 1, 3}},
			expected: 2,
		},
		{
			nums:     []int{4, 3, 2, 1},
			queries:  [][]int{{1, 3, 2}, {0, 2, 1}},
			expected: -1,
		},
	}

	for i, tc := range testCases {
		result := minZeroArray(tc.nums, tc.queries)
		if result != tc.expected {
			t.Errorf("Test case %d: expected %d, got %d", i, tc.expected, result)
		}
	}
}

/*
// @lcpr case=start
// [2,0,2]\n[[0,2,1],[0,2,1],[1,1,3]]\n
// @lcpr case=end

// @lcpr case=start
// [4,3,2,1]\n[[1,3,2],[0,2,1]]\n
// @lcpr case=end

*/
