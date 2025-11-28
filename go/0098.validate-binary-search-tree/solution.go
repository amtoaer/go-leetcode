// Created by amtoaer at 2025/11/28 21:07
// leetgo: 1.4.15
// https://leetcode.cn/problems/validate-binary-search-tree/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(node *TreeNode, lower, upper int) bool {
	if node == nil {
		return true
	}
	if node.Val <= lower || node.Val >= upper {
		return false
	}
	return helper(node.Left, lower, node.Val) && helper(node.Right, node.Val, upper)
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	root := Deserialize[*TreeNode](ReadLine(stdin))
	ans := isValidBST(root)

	fmt.Println("\noutput:", Serialize(ans))
}
