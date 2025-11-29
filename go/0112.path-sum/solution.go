// Created by amtoaer at 2025/11/29 17:23
// leetgo: 1.4.15
// https://leetcode.cn/problems/path-sum/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Val == targetSum && root.Left == nil && root.Right == nil {
		return true
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	root := Deserialize[*TreeNode](ReadLine(stdin))
	targetSum := Deserialize[int](ReadLine(stdin))
	ans := hasPathSum(root, targetSum)

	fmt.Println("\noutput:", Serialize(ans))
}
