// Created by amtoaer at 2025/11/18 15:33
// leetgo: 1.4.15
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maxProfit(prices []int) (ans int) {
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		ans = max(ans, prices[i]-minPrice)
		minPrice = min(minPrice, prices[i])
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
