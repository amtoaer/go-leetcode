// Created by amtoaer at 2025/11/24 21:18
// leetgo: 1.4.15
// https://leetcode.cn/problems/next-permutation/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func nextPermutation(nums []int) {
	var smallIdx, bigIdx int
	for smallIdx = len(nums) - 2; smallIdx >= 0; smallIdx-- {
		if nums[smallIdx] < nums[smallIdx+1] {
			break
		}
	}
	if smallIdx != -1 {
		for bigIdx = len(nums) - 1; bigIdx > smallIdx; bigIdx-- {
			if nums[bigIdx] > nums[smallIdx] {
				break
			}
		}
		nums[bigIdx], nums[smallIdx] = nums[smallIdx], nums[bigIdx]
	}
	for i, j := smallIdx+1, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	nextPermutation(nums)
	ans := nums

	fmt.Println("\noutput:", Serialize(ans))
}
