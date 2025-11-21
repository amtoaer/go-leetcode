// Created by amtoaer at 2025/11/21 17:20
// leetgo: 1.4.15
// https://leetcode.cn/problems/binary-tree-maximum-path-sum/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maxPathSum(root *TreeNode) (ans int) {
	ans = math.MinInt
	var helper func(node *TreeNode) int
	helper = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left, right := helper(node.Left), helper(node.Right)
		ans = max(ans, left+right+node.Val)
		return max(max(left, right)+node.Val, 0)
	}
	helper(root)
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	root := Deserialize[*TreeNode](ReadLine(stdin))
	ans := maxPathSum(root)

	fmt.Println("\noutput:", Serialize(ans))
}
