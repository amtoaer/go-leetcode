// Created by amtoaer at 2025/12/12 00:47
// leetgo: 1.4.15
// https://leetcode.cn/problems/single-number/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func singleNumber(nums []int) (ans int) {
	for _, num := range nums {
		ans ^= num
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := singleNumber(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
