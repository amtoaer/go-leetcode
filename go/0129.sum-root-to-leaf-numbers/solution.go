// Created by amtoaer at 2025/11/27 18:03
// leetgo: 1.4.15
// https://leetcode.cn/problems/sum-root-to-leaf-numbers/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func sumNumbers(root *TreeNode) (ans int) {
	var dfs func(node *TreeNode, val int)
	dfs = func(node *TreeNode, val int) {
		if node == nil {
			return
		}
		val = val*10 + node.Val
		if node.Left == nil && node.Right == nil {
			ans = ans + val
			return
		}
		dfs(node.Left, val)
		dfs(node.Right, val)
	}
	dfs(root, 0)
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	root := Deserialize[*TreeNode](ReadLine(stdin))
	ans := sumNumbers(root)

	fmt.Println("\noutput:", Serialize(ans))
}
