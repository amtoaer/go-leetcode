// Created by amtoaer at 2025/11/24 21:31
// leetgo: 1.4.15
// https://leetcode.cn/problems/first-missing-positive/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func firstMissingPositive(nums []int) (ans int) {
	for idx := range nums {
		for {
			val := nums[idx]
			if val > 0 && val <= len(nums) && val-1 != idx && nums[val-1] != nums[idx] {
				nums[val-1], nums[idx] = nums[idx], nums[val-1]
			} else {
				break
			}
		}
	}
	for idx, val := range nums {
		if val != idx+1 {
			return idx + 1
		}
	}
	return len(nums) + 1
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := firstMissingPositive(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
