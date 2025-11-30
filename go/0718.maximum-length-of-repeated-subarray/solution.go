// Created by amtoaer at 2025/11/30 16:31
// leetgo: 1.4.15
// https://leetcode.cn/problems/maximum-length-of-repeated-subarray/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func findLength(nums1 []int, nums2 []int) (ans int) {
	compare := func(nums1, nums2 []int) {
		var cnt int
		for i := range min(len(nums1), len(nums2)) {
			if nums1[i] == nums2[i] {
				cnt++
			} else {
				cnt = 0
			}
			ans = max(ans, cnt)
		}
	}
	for i := range nums1 {
		compare(nums1[i:], nums2)
	}
	for i := range nums2 {
		compare(nums1, nums2[i:])
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums1 := Deserialize[[]int](ReadLine(stdin))
	nums2 := Deserialize[[]int](ReadLine(stdin))
	ans := findLength(nums1, nums2)

	fmt.Println("\noutput:", Serialize(ans))
}
