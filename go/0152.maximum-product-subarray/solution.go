// Created by amtoaer at 2025/11/29 15:27
// leetgo: 1.4.15
// https://leetcode.cn/problems/maximum-product-subarray/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maxProduct(nums []int) (ans int) {
	ans = math.MinInt64
	dpMax := make([]int, len(nums))
	dpMin := make([]int, len(nums))
	for i := range nums {
		if i == 0 {
			dpMax[i] = nums[i]
			dpMin[i] = nums[i]
		} else {
			dpMax[i] = max(nums[i], dpMax[i-1]*nums[i], dpMin[i-1]*nums[i])
			dpMin[i] = min(nums[i], dpMax[i-1]*nums[i], dpMin[i-1]*nums[i])
		}
		ans = max(ans, dpMax[i])
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := maxProduct(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
