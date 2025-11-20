// Created by amtoaer at 2025/11/20 22:22
// leetgo: 1.4.15
// https://leetcode.cn/problems/maximum-subarray/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maxSubArray(nums []int) (ans int) {
	var tmp int
	ans = math.MinInt
	for _, num := range nums {
		tmp = max(tmp+num, num)
		ans = max(ans, tmp)
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := maxSubArray(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
