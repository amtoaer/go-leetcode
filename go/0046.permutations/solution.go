// Created by amtoaer at 2025/11/20 22:06
// leetgo: 1.4.15
// https://leetcode.cn/problems/permutations/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func permute(nums []int) (ans [][]int) {
	visited := make([]bool, len(nums))
	tmpAns := make([]int, 0, len(nums))
	var backtrack func(int)
	backtrack = func(depth int) {
		if depth == len(nums) {
			ans = append(ans, append([]int{}, tmpAns...))
			return
		}
		for i, num := range nums {
			if visited[i] {
				continue
			}
			visited[i] = true
			tmpAns = append(tmpAns, num)
			backtrack(depth + 1)
			visited[i] = false
			tmpAns = tmpAns[:len(tmpAns)-1]
		}
	}
	backtrack(0)
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := permute(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
