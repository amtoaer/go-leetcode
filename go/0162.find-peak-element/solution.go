// Created by amtoaer at 2025/11/29 00:50
// leetgo: 1.4.15
// https://leetcode.cn/problems/find-peak-element/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func findPeakElement(nums []int) (ans int) {
	get := func(idx int) int {
		if idx >= 0 && idx < len(nums) {
			return nums[idx]
		}
		return math.MinInt64
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if get(mid) > get(mid-1) && get(mid) > get(mid+1) {
			return mid
		}
		if get(mid) < get(mid+1) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := findPeakElement(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
