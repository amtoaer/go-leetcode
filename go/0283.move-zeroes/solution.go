// Created by amtoaer at 2025/11/30 16:07
// leetgo: 1.4.15
// https://leetcode.cn/problems/move-zeroes/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func moveZeroes(nums []int) {
	var slow int
	for fast := range nums {
		if nums[fast] != 0 {
			nums[fast], nums[slow] = nums[slow], nums[fast]
			slow++
		}
	}
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	moveZeroes(nums)
	ans := nums

	fmt.Println("\noutput:", Serialize(ans))
}
