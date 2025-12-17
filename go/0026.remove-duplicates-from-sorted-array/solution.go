// Created by amtoaer at 2025/12/17 23:46
// leetgo: 1.4.15
// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func removeDuplicates(nums []int) int {
	var left int
	for right := range nums {
		if right == 0 || nums[right] != nums[right-1] {
			nums[left] = nums[right]
			left++
		}
	}
	return left
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := removeDuplicates(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
