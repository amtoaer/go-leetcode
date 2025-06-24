/*
 * @lc app=leetcode.cn id=2200 lang=golang
 * @lcpr version=30201
 *
 * [2200] 找出数组中的所有 K 近邻下标
 */

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// @lc code=start
func findKDistantIndices(nums []int, key int, k int) []int {
	flag := make([]bool, len(nums))
	for idx, num := range nums {
		if num == key {
			for i := max(idx-k, 0); i <= min(idx+k, len(nums)-1); i++ {
				flag[i] = true
			}
		}
	}
	var res []int
	for idx, label := range flag {
		if label {
			res = append(res, idx)
		}
	}
	return res
}

// @lc code=end
func Test(t *testing.T) {
	testCases := []struct {
		nums     []int
		key      int
		k        int
		expected []int
	}{
		{[]int{3, 4, 9, 1, 3, 9, 5}, 9, 1, []int{1, 2, 3, 4, 5, 6}},
		{[]int{2, 2, 2, 2, 2}, 2, 2, []int{0, 1, 2, 3, 4}},
		// Additional test cases
		{[]int{1, 2, 3, 4, 5}, 3, 2, []int{0, 1, 2, 3, 4}},
		{[]int{1, 2, 3, 4, 5}, 6, 1, []int{}},               // key not in array
		{[]int{1, 2, 3, 1, 2}, 1, 0, []int{0, 3}},           // k=0 means only exact indices
		{[]int{5, 5, 5, 5, 5}, 5, 10, []int{0, 1, 2, 3, 4}}, // k larger than array length
		{[]int{1}, 1, 3, []int{0}},                          // single element array
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 1, []int{3, 4, 5}},
	}

	for _, tc := range testCases {
		result := findKDistantIndices(tc.nums, tc.key, tc.k)
		assert.ElementsMatch(t, result, tc.expected)
	}
}

/*
// @lcpr case=start
// [3,4,9,1,3,9,5]\n9\n1\n
// @lcpr case=end

// @lcpr case=start
// [2,2,2,2,2]\n2\n2\n
// @lcpr case=end

*/
