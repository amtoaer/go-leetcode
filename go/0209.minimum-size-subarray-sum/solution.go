// Created by amtoaer at 2025/11/29 17:35
// leetgo: 1.4.15
// https://leetcode.cn/problems/minimum-size-subarray-sum/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func minSubArrayLen(target int, nums []int) (ans int) {
	ans = math.MaxInt64
	var left, cur int
	for right := range nums {
		cur += nums[right]
		for cur >= target {
			ans = min(ans, right-left+1)
			cur -= nums[left]
			left++
		}
	}
	if ans == math.MaxInt64 {
		return 0
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	target := Deserialize[int](ReadLine(stdin))
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := minSubArrayLen(target, nums)

	fmt.Println("\noutput:", Serialize(ans))
}
