// Created by amtoaer at 2025/12/18 00:35
// leetgo: 1.4.15
// https://leetcode.cn/problems/daily-temperatures/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func dailyTemperatures(temperatures []int) (ans []int) {
	ans = make([]int, len(temperatures))
	var stack []int
	for idx, temperature := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperature {
			ans[stack[len(stack)-1]] = idx - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, idx)
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	temperatures := Deserialize[[]int](ReadLine(stdin))
	ans := dailyTemperatures(temperatures)

	fmt.Println("\noutput:", Serialize(ans))
}
