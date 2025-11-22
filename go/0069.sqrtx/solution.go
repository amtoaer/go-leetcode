// Created by amtoaer at 2025/11/22 11:35
// leetgo: 1.4.15
// https://leetcode.cn/problems/sqrtx/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func mySqrt(x int) (ans int) {
	if x <= 1 {
		return x
	}
	left, right := 0, x/2
	for left <= right {
		mid := left + (right-left)>>1
		res := mid * mid
		if res == x {
			return mid
		}
		if res < x {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	x := Deserialize[int](ReadLine(stdin))
	ans := mySqrt(x)

	fmt.Println("\noutput:", Serialize(ans))
}
