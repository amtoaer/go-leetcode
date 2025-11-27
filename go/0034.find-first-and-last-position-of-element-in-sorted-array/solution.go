// Created by amtoaer at 2025/11/27 18:29
// leetgo: 1.4.15
// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func searchRange(nums []int, target int) []int {
	left := binarySearch(nums, target)
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}
	right := binarySearch(nums, target+1) - 1
	return []int{left, right}
}

func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := searchRange(nums, target)

	fmt.Println("\noutput:", Serialize(ans))
}
