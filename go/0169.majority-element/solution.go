// Created by amtoaer at 2025/11/29 22:50
// leetgo: 1.4.15
// https://leetcode.cn/problems/majority-element/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func majorityElement(nums []int) (ans int) {
	var cnt int
	for _, num := range nums {
		if num == ans {
			cnt++
		} else if cnt > 0 {
			cnt--
		} else {
			ans = num
			cnt = 1
		}
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := majorityElement(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
