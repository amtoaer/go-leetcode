// Created by amtoaer at 2025/11/27 20:18
// leetgo: 1.4.15
// https://leetcode.cn/problems/combination-sum/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func combinationSum(candidates []int, target int) (ans [][]int) {
	var tmp []int
	var backtrack func(idx, target int)
	backtrack = func(idx, target int) {
		if target < 0 {
			return
		}
		if target == 0 && len(tmp) > 0 {
			tmpRes := make([]int, len(tmp))
			copy(tmpRes, tmp)
			ans = append(ans, tmpRes)
			return
		}
		for cur := idx; cur < len(candidates); cur++ {
			if target-candidates[cur] >= 0 {
				tmp = append(tmp, candidates[cur])
				backtrack(cur, target-candidates[cur])
				tmp = tmp[:len(tmp)-1]
			}
		}
	}
	backtrack(0, target)
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	candidates := Deserialize[[]int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := combinationSum(candidates, target)

	fmt.Println("\noutput:", Serialize(ans))
}
