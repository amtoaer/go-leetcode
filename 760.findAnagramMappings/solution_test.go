/*
 * @lc app=leetcode.cn id=760 lang=golang
 * @lcpr version=30202
 *
 * [760] 找出变位映射
 */

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// @lc code=start
type helper struct {
	offset int
	items  []int
}

func newHelper(item int) *helper {
	return &helper{
		items: []int{item},
	}
}

func anagramMappings(nums1 []int, nums2 []int) []int {
	m := make(map[int]*helper)
	for idx, num := range nums2 {
		if h, ok := m[num]; ok {
			h.items = append(h.items, idx)
		} else {
			m[num] = newHelper(idx)
		}
	}
	res := make([]int, len(nums1))
	for idx, num := range nums1 {
		helper := m[num]
		res[idx] = helper.items[helper.offset]
		helper.offset++
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		{[]int{12, 28, 46, 32, 50}, []int{50, 12, 32, 46, 28}, []int{1, 4, 3, 2, 0}},
		{[]int{84, 46}, []int{84, 46}, []int{0, 1}},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, anagramMappings(tt.nums1, tt.nums2))
	}
}

/*
// @lcpr case=start
// [12,28,46,32,50]\n[50,12,32,46,28]\n
// @lcpr case=end

// @lcpr case=start
// [84,46]\n[84,46]\n
// @lcpr case=end

*/
