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
 * @lc app=leetcode.cn id=94 lang=golang
 * @lcpr version=20003
 *
 * [94] 二叉树的中序遍历
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
func inorderTraversal(root *TreeNode) []int {
	var res []int
	var inorder func(*TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		res = append(res, root.Val)
		inorder(root.Right)
	}
	inorder(root)
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [1,null,2,3]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		root     *TreeNode
		expected []int
	}{
		{&TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}, []int{1, 3, 2}},
		{nil, []int{}},
		{&TreeNode{1, nil, nil}, []int{1}},
		{&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}, []int{2, 1, 3}},
		{&TreeNode{1, &TreeNode{2, &TreeNode{3, nil, nil}, nil}, nil}, []int{3, 2, 1}},
	}

	for _, test := range tests {
		result := inorderTraversal(test.root)
		if !equal(result, test.expected) {
			t.Errorf("inorderTraversal(%v) = %v; expected %v", test.root, result, test.expected)
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
