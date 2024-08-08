package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=3132 lang=golang
 *
 * [3132] Find the Integer Added to Array II
 */

// @lc code=start
func minimumAddedInteger(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	choices := []int{nums2[0] - nums1[2], nums2[0] - nums1[1], nums2[0] - nums1[0]}
label:
	for _, choice := range choices {
		left, right := 0, 0
		count := 0
		for left < len(nums1) && right < len(nums2) {
			if nums2[right]-nums1[left] != choice {
				count++
				if count > 2 {
					continue label
				}
			} else {
				right++
			}
			left++
		}
		return choice
	}
	return -1
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  []int
		input2 []int
		output int
	}{
		{[]int{4, 20, 16, 12, 8}, []int{14, 18, 10}, -2},
		{[]int{3, 5, 5, 3}, []int{7, 7}, 2},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, minimumAddedInteger(tt.input, tt.input2))
	}

}
