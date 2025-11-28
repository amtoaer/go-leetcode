// Created by amtoaer at 2025/11/28 11:16
// leetgo: 1.4.15
// https://leetcode.cn/problems/longest-consecutive-sequence/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func longestConsecutive(nums []int) (ans int) {
	m := make(map[int]any)
	for _, num := range nums {
		m[num] = struct{}{}
	}
	for num := range m {
		if _, ok := m[num-1]; ok {
			continue
		}
		var cnt int
		for {
			if _, ok := m[num]; !ok {
				break
			}
			cnt++
			num++
		}
		ans = max(ans, cnt)
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := longestConsecutive(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
