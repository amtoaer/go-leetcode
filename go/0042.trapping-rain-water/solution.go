// Created by amtoaer at 2025/11/21 16:48
// leetgo: 1.4.15
// https://leetcode.cn/problems/trapping-rain-water/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func trap(height []int) (ans int) {
	var stack []int
	for idx, h := range height {
		for len(stack) > 0 && h > height[stack[len(stack)-1]] {
			bottom := height[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				leftIdx := stack[len(stack)-1]
				left := height[leftIdx]
				ans += (min(h, left) - bottom) * (idx - leftIdx - 1)
			}
		}
		stack = append(stack, idx)
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	height := Deserialize[[]int](ReadLine(stdin))
	ans := trap(height)

	fmt.Println("\noutput:", Serialize(ans))
}
