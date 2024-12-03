package main

import (
	"testing"
)

type Node struct {
	Val       int
	Neighbors []*Node
}

/*
 * @lc app=leetcode.cn id=133 lang=golang
 * @lcpr version=20004
 *
 * [133] 克隆图
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	nodeMap := make(map[*Node]*Node)
	var clonedNode func(*Node) *Node
	clonedNode = func(node *Node) *Node {
		if newNode, ok := nodeMap[node]; ok {
			return newNode
		}
		newNode := &Node{
			Val:       node.Val,
			Neighbors: make([]*Node, 0, len(node.Neighbors)),
		}
		nodeMap[node] = newNode
		for _, neighbor := range node.Neighbors {
			newNode.Neighbors = append(newNode.Neighbors, clonedNode(neighbor))
		}
		return newNode
	}
	return clonedNode(node)
}

// @lc code=end

/*
// @lcpr case=start
// [[2,4],[1,3],[2,4],[1,3]]\n
// @lcpr case=end

// @lcpr case=start
// [[]]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		input    *Node
		expected *Node
	}{
		{
			input: &Node{
				Val: 1,
				Neighbors: []*Node{
					{Val: 2},
					{Val: 4},
				},
			},
			expected: &Node{
				Val: 1,
				Neighbors: []*Node{
					{Val: 2},
					{Val: 4},
				},
			},
		},
		{
			input: &Node{
				Val: 1,
				Neighbors: []*Node{
					{Val: 2, Neighbors: []*Node{{Val: 1}}},
					{Val: 4, Neighbors: []*Node{{Val: 1}}},
				},
			},
			expected: &Node{
				Val: 1,
				Neighbors: []*Node{
					{Val: 2, Neighbors: []*Node{{Val: 1}}},
					{Val: 4, Neighbors: []*Node{{Val: 1}}},
				},
			},
		},
		{
			input: &Node{
				Val: 1,
				Neighbors: []*Node{
					{Val: 2, Neighbors: []*Node{{Val: 1}, {Val: 3}}},
					{Val: 4, Neighbors: []*Node{{Val: 1}, {Val: 3}}},
				},
			},
			expected: &Node{
				Val: 1,
				Neighbors: []*Node{
					{Val: 2, Neighbors: []*Node{{Val: 1}, {Val: 3}}},
					{Val: 4, Neighbors: []*Node{{Val: 1}, {Val: 3}}},
				},
			},
		},
	}

	for _, test := range tests {
		cloned := cloneGraph(test.input)
		if !isEqual(cloned, test.expected) {
			t.Errorf("expected %v, but got %v", test.expected, cloned)
		}
	}
}

func isEqual(node1, node2 *Node) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}
	if node1.Val != node2.Val {
		return false
	}
	if len(node1.Neighbors) != len(node2.Neighbors) {
		return false
	}
	for i := range node1.Neighbors {
		if !isEqual(node1.Neighbors[i], node2.Neighbors[i]) {
			return false
		}
	}
	return true
}
