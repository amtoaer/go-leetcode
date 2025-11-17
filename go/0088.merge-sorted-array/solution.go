// Created by amtoaer at 2025/11/18 02:43
// leetgo: 1.4.15
// https://leetcode.cn/problems/merge-sorted-array/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func merge(nums1 []int, m int, nums2 []int, n int) {
	target := m + n - 1
	m, n = m-1, n-1
	for m >= 0 && n >= 0 {
		if nums1[m] > nums2[n] {
			nums1[target] = nums1[m]
			m--
		} else {
			nums1[target] = nums2[n]
			n--
		}
		target--
	}
	for n >= 0 {
		nums1[n] = nums2[n]
		n--
	}
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums1 := Deserialize[[]int](ReadLine(stdin))
	m := Deserialize[int](ReadLine(stdin))
	nums2 := Deserialize[[]int](ReadLine(stdin))
	n := Deserialize[int](ReadLine(stdin))
	merge(nums1, m, nums2, n)
	ans := nums1

	fmt.Println("\noutput:", Serialize(ans))
}
