// Created by amtoaer at 2025/11/22 13:12
// leetgo: 1.4.15
// https://leetcode.cn/problems/coin-change/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	for _, c := range coins {
		if c <= amount {
			dp[c] = 1
		}
	}
	for i := 1; i <= amount; i++ {
		if dp[i] == 1 {
			continue
		}
		tmp := math.MaxInt
		for _, c := range coins {
			if i-c >= 0 && dp[i-c] != 0 {
				tmp = min(tmp, dp[i-c]+1)
			}
		}
		if tmp != math.MaxInt {
			dp[i] = tmp
		}
	}
	if dp[amount] == 0 {
		return -1
	}
	return dp[amount]
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	coins := Deserialize[[]int](ReadLine(stdin))
	amount := Deserialize[int](ReadLine(stdin))
	ans := coinChange(coins, amount)

	fmt.Println("\noutput:", Serialize(ans))
}
