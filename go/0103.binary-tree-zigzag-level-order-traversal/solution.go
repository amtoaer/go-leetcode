// Created by amtoaer at 2025/11/20 13:33
// leetgo: 1.4.15
// https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func zigzagLevelOrder(root *TreeNode) (ans [][]int) {
	stack := []*TreeNode{root}
	reverse := true
	for len(stack) > 0 {
		length := len(stack)
		var tmpRes []int
		for _, node := range stack[:length] {
			if node == nil {
				continue
			}
			tmpRes = append(tmpRes, node.Val)
			stack = append(stack, node.Left, node.Right)
		}
		reverse = !reverse
		if len(tmpRes) > 0 {
			if reverse {
				for i, j := 0, len(tmpRes)-1; i < j; i, j = i+1, j-1 {
					tmpRes[i], tmpRes[j] = tmpRes[j], tmpRes[i]
				}
			}
			ans = append(ans, tmpRes)
		}
		stack = stack[length:]
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	root := Deserialize[*TreeNode](ReadLine(stdin))
	ans := zigzagLevelOrder(root)

	fmt.Println("\noutput:", Serialize(ans))
}
