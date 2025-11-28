// Created by amtoaer at 2025/11/29 01:51
// leetgo: 1.4.15
// https://leetcode.cn/problems/path-sum-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func pathSum(root *TreeNode, targetSum int) (ans [][]int) {
	var tmp []int
	var helper func(node *TreeNode, target int)
	helper = func(node *TreeNode, target int) {
		if node == nil {
			return
		}
		target = target - node.Val
		tmp = append(tmp, node.Val)
		if target == 0 && node.Left == nil && node.Right == nil {
			if len(tmp) != 0 {
				tmpRes := make([]int, len(tmp))
				copy(tmpRes, tmp)
				ans = append(ans, tmpRes)
			}
		} else {
			helper(node.Left, target)
			helper(node.Right, target)
		}
		tmp = tmp[:len(tmp)-1]
	}
	helper(root, targetSum)
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	root := Deserialize[*TreeNode](ReadLine(stdin))
	targetSum := Deserialize[int](ReadLine(stdin))
	ans := pathSum(root, targetSum)

	fmt.Println("\noutput:", Serialize(ans))
}
