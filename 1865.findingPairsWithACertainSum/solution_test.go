/*
 * @lc app=leetcode.cn id=1865 lang=golang
 * @lcpr version=30201
 *
 * [1865] 找出和为指定值的下标对
 */

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// @lc code=start
type FindSumPairs struct {
	nums1, nums2 []int
	cnt          map[int]int
}

func Constructor(nums1 []int, nums2 []int) FindSumPairs {
	res := FindSumPairs{
		nums1, nums2, make(map[int]int),
	}
	for _, num := range nums2 {
		res.cnt[num]++
	}
	return res
}

func (f *FindSumPairs) Add(index int, val int) {
	f.cnt[f.nums2[index]]--
	f.nums2[index] += val
	f.cnt[f.nums2[index]]++
}

func (f *FindSumPairs) Count(tot int) int {
	var res int
	for _, num := range f.nums1 {
		res += f.cnt[tot-num]
	}
	return res
}

/**
 * Your FindSumPairs object will be instantiated and called as such:
 * obj := Constructor(nums1, nums2);
 * obj.Add(index,val);
 * param_2 := obj.Count(tot);
 */
// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		operations []string
		params     [][]int
		expected   []any
	}{
		{
			operations: []string{"count", "add", "count", "count", "add", "add", "count"},
			params:     [][]int{{1, 1, 2, 2, 2, 3}, {1, 4, 5, 2, 5, 4}, {7}, {3, 2}, {8}, {4}, {0, 1}, {1, 1}, {7}},
			expected:   []any{8, nil, 2, 1, nil, nil, 11},
		},
	}

	for _, tt := range tc {
		nums1 := tt.params[0]
		nums2 := tt.params[1]
		findSumPairs := Constructor(nums1, nums2)

		for i := 2; i < len(tt.params); i++ {
			switch tt.operations[i-2] {
			case "add":
				findSumPairs.Add(tt.params[i][0], tt.params[i][1])
			case "count":
				result := findSumPairs.Count(tt.params[i][0])
				t.Logf("Count(%d) = %d", tt.params[i][0], result)
				assert.Equal(t, tt.expected[i-2], result)
			}
		}
	}
}

/*
// @lcpr case=start
// ["FindSumPairs", "count", "add", "count", "count", "add", "add", "count"]\n[[[1, 1, 2, 2, 2, 3], [1, 4, 5, 2, 5, 4]], [7], [3, 2], [8], [4], [0, 1], [1, 1], [7]]\n
// @lcpr case=end

*/
