package main

import (
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * @lc app=leetcode.cn id=107 lang=golang
 * @lcpr version=20004
 *
 * [107] 二叉树的层序遍历 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		length := len(stack)
		var row []int
		for _, node := range stack[:length] {
			row = append(row, node.Val)
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
		res = append(res, row)
		stack = stack[length:]
	}
	left, right := 0, len(res)-1
	for left < right {
		res[left], res[right] = res[right], res[left]
		left++
		right--
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [3,9,20,null,null,15,7]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		root     *TreeNode
		expected [][]int
	}{
		{
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			expected: [][]int{
				{15, 7},
				{9, 20},
				{3},
			},
		},
		{
			root: &TreeNode{
				Val: 1,
			},
			expected: [][]int{
				{1},
			},
		},
		{
			root:     nil,
			expected: [][]int{},
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 6,
					},
				},
			},
			expected: [][]int{
				{4, 5, 6},
				{2, 3},
				{1},
			},
		},
	}

	for _, test := range tests {
		result := levelOrderBottom(test.root)
		if !equal(result, test.expected) {
			t.Errorf("For root %v, expected %v, but got %v", test.root, test.expected, result)
		}
	}
}

func equal(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
