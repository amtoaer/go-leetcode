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
 * @lc app=leetcode.cn id=236 lang=golang
 * @lcpr version=20004
 *
 * [236] 二叉树的最近公共祖先
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
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left, right := lowestCommonAncestor(root.Left, p, q), lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

// @lc code=end

/*
// @lcpr case=start
// [3,5,1,6,2,0,8,null,null,7,4]\n5\n1\n
// @lcpr case=end

// @lcpr case=start
// [3,5,1,6,2,0,8,null,null,7,4]\n5\n4\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n1\n2\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		root     *TreeNode
		p        *TreeNode
		q        *TreeNode
		expected *TreeNode
	}{
		{
			root: &TreeNode{3,
				&TreeNode{5,
					&TreeNode{6, nil, nil},
					&TreeNode{2,
						&TreeNode{7, nil, nil},
						&TreeNode{4, nil, nil},
					},
				},
				&TreeNode{1,
					&TreeNode{0, nil, nil},
					&TreeNode{8, nil, nil},
				},
			},
			p:        &TreeNode{Val: 5},
			q:        &TreeNode{Val: 1},
			expected: &TreeNode{Val: 3},
		},
		{
			root: &TreeNode{3,
				&TreeNode{5,
					&TreeNode{6, nil, nil},
					&TreeNode{2,
						&TreeNode{7, nil, nil},
						&TreeNode{4, nil, nil},
					},
				},
				&TreeNode{1,
					&TreeNode{0, nil, nil},
					&TreeNode{8, nil, nil},
				},
			},
			p:        &TreeNode{Val: 5},
			q:        &TreeNode{Val: 4},
			expected: &TreeNode{Val: 5},
		},
		{
			root: &TreeNode{1,
				&TreeNode{2, nil, nil},
				nil,
			},
			p:        &TreeNode{Val: 1},
			q:        &TreeNode{Val: 2},
			expected: &TreeNode{Val: 1},
		},
	}

	for _, test := range tests {
		result := lowestCommonAncestor(test.root, test.p, test.q)
		if result.Val != test.expected.Val {
			t.Errorf("expected %v, got %v", test.expected.Val, result.Val)
		}
	}
}
