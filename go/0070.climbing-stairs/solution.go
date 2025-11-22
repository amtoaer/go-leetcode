// Created by amtoaer at 2025/11/22 13:07
// leetgo: 1.4.15
// https://leetcode.cn/problems/climbing-stairs/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	a, b := 1, 2
	for n > 2 {
		a, b = b, a+b
		n--
	}
	return b
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	n := Deserialize[int](ReadLine(stdin))
	ans := climbStairs(n)

	fmt.Println("\noutput:", Serialize(ans))
}
