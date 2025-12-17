// Created by amtoaer at 2025/12/17 23:12
// leetgo: 1.4.15
// https://leetcode.cn/problems/check-completeness-of-a-binary-tree/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func isCompleteTree(root *TreeNode) bool {
	stack := []*TreeNode{root}
	requiredLen := 1
	for len(stack) > 0 {
		length := len(stack)
		lastLevelSatisfied := requiredLen == length
		hasNil := false
		for _, node := range stack[:length] {
			if !lastLevelSatisfied && (node.Left != nil || node.Right != nil) {
				return false
			}
			if node.Left != nil {
				if hasNil {
					return false
				}
				stack = append(stack, node.Left)
			}
			hasNil = hasNil || node.Left == nil
			if node.Right != nil {
				if hasNil {
					return false
				}
				stack = append(stack, node.Right)
			}
			hasNil = hasNil || node.Right == nil
		}
		stack = stack[length:]
		requiredLen *= 2
	}
	return true
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	root := Deserialize[*TreeNode](ReadLine(stdin))
	ans := isCompleteTree(root)

	fmt.Println("\noutput:", Serialize(ans))
}
