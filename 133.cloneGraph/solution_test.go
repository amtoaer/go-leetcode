package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Node struct {
	Val       int
	Neighbors []*Node
}

/*
 * @lc app=leetcode.cn id=133 lang=golang
 *
 * [133] Clone Graph
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	preToAfter := make(map[*Node]*Node)
	var cloneNode func(*Node) *Node
	cloneNode = func(n *Node) *Node {
		if n == nil {
			return nil
		}
		if newNode, ok := preToAfter[n]; ok {
			return newNode
		}
		newNode := &Node{
			Val: n.Val,
		}
		preToAfter[n] = newNode
		newNeighbors := make([]*Node, 0, len(n.Neighbors))
		for _, neighbor := range n.Neighbors {
			newNeighbors = append(newNeighbors, cloneNode(neighbor))
		}
		newNode.Neighbors = newNeighbors
		return newNode
	}
	return cloneNode(node)
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  *Node
		output *Node
	}{
		{nil, nil},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, cloneGraph(tt.input))
	}
}
