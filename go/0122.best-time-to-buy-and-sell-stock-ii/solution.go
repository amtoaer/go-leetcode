// Created by amtoaer at 2025/11/28 11:56
// leetgo: 1.4.15
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maxProfit(prices []int) (ans int) {
	for i := range len(prices) - 1 {
		ans += max(prices[i+1]-prices[i], 0)
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	prices := Deserialize[[]int](ReadLine(stdin))
	ans := maxProfit(prices)

	fmt.Println("\noutput:", Serialize(ans))
}
