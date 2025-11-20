// Created by amtoaer at 2025/11/20 14:04
// leetgo: 1.4.15
// https://leetcode.cn/problems/longest-increasing-subsequence/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func lengthOfLIS(nums []int) (ans int) {
	dp := make([]int, len(nums))
	for idx := range dp {
		dp[idx] = 1
	}
	for i := range nums {
		for j := range i {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ans = max(ans, dp[i])
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := lengthOfLIS(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
