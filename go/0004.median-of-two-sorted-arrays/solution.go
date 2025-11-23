// Created by amtoaer at 2025/11/23 15:48
// leetgo: 1.4.15
// https://leetcode.cn/problems/median-of-two-sorted-arrays/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength&1 == 1 {
		return float64(findKthElement(nums1, nums2, totalLength/2+1))
	} else {
		return float64(findKthElement(nums1, nums2, totalLength/2)+findKthElement(nums1, nums2, totalLength/2+1)) / 2
	}
}

func findKthElement(nums1, nums2 []int, k int) int {
	for {
		if len(nums1) > len(nums2) {
			nums1, nums2 = nums2, nums1
		}
		// len(nums1) <= len(nums2)
		if len(nums1) == 0 {
			return nums2[k-1]
		}
		if k == 1 {
			return min(nums1[0], nums2[0])
		}
		p1, p2 := min(k/2, len(nums1)), min(k/2, len(nums2))
		if nums1[p1-1] <= nums2[p2-1] {
			k -= p1
			nums1 = nums1[p1:]
		} else {
			k -= p2
			nums2 = nums2[p2:]
		}
	}
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums1 := Deserialize[[]int](ReadLine(stdin))
	nums2 := Deserialize[[]int](ReadLine(stdin))
	ans := findMedianSortedArrays(nums1, nums2)

	fmt.Println("\noutput:", Serialize(ans))
}
