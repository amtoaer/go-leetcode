// Created by amtoaer at 2025/11/17 21:17
// leetgo: 1.4.15
// https://leetcode.cn/problems/two-sum/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func twoSum(nums []int, target int) (ans []int) {
	m := make(map[int]int)
	for idx, num := range nums {
		if prevIdx, ok := m[target-num]; ok {
			return []int{prevIdx, idx}
		}
		m[num] = idx
	}
	return nil
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := twoSum(nums, target)

	fmt.Println("\noutput:", Serialize(ans))
}
