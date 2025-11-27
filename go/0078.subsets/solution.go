// Created by amtoaer at 2025/11/27 16:38
// leetgo: 1.4.15
// https://leetcode.cn/problems/subsets/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func subsets(nums []int) (ans [][]int) {
	var backtrack func(i int)
	tmp := make([]int, 0, len(nums))
	backtrack = func(i int) {
		if i >= len(nums) {
			tmpRes := make([]int, len(tmp))
			copy(tmpRes, tmp)
			ans = append(ans, tmpRes)
			return
		}
		tmp = append(tmp, nums[i])
		backtrack(i + 1)
		tmp = tmp[:len(tmp)-1]
		backtrack(i + 1)

	}
	backtrack(0)
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := subsets(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
