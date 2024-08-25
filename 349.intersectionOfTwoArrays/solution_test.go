package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] Intersection of Two Arrays
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	next := func(arr []int, idx int) int {
		for cur := idx; cur < len(arr); cur++ {
			if arr[cur] != arr[idx] {
				return cur
			}
		}
		return len(arr)
	}
	var lp, rp int
	res := []int{}
	for lp < len(nums1) && rp < len(nums2) {
		lpv, rpv := nums1[lp], nums2[rp]
		if lpv == rpv {
			res = append(res, lpv)
			lp = next(nums1, lp)
			rp = next(nums2, rp)
		} else if lpv < rpv {
			lp = next(nums1, lp)
		} else {
			rp = next(nums2, rp)
		}
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		nums1  []int
		nums2  []int
		expect []int
	}{
		{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2}},
		{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{4, 9}},
	}
	for _, tt := range tc {
		assert.ElementsMatch(t, tt.expect, intersection(tt.nums1, tt.nums2))
	}
}
