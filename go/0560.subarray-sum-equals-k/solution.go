// Created by amtoaer at 2025/11/29 17:18
// leetgo: 1.4.15
// https://leetcode.cn/problems/subarray-sum-equals-k/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func subarraySum(nums []int, k int) (ans int) {
	prefixSum := make([]int, len(nums)+1)
	for i := range nums {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}
	m := make(map[int]int, len(prefixSum))
	for _, s := range prefixSum {
		ans += m[s-k]
		m[s]++
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	ans := subarraySum(nums, k)

	fmt.Println("\noutput:", Serialize(ans))
}
