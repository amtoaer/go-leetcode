package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

/*
 * @lc app=leetcode.cn id=427 lang=golang
 *
 * [427] Construct Quad Tree
 */

// @lc code=start
/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */

func construct(grid [][]int) *Node {
	size := len(grid)
	var buildNode func(left, right, up, down int) *Node
	buildNode = func(left, right, up, down int) *Node {
		value := grid[up][left]
		for i := up; i < down; i++ {
			for j := left; j < right; j++ {
				if grid[i][j] != value {
					return &Node{
						IsLeaf:      false,
						TopLeft:     buildNode(left, (left+right)/2, up, (up+down)/2),
						TopRight:    buildNode((left+right)/2, right, up, (up+down)/2),
						BottomLeft:  buildNode(left, (left+right)/2, (up+down)/2, down),
						BottomRight: buildNode((left+right)/2, right, (up+down)/2, down),
					}
				}
			}
		}
		return &Node{
			IsLeaf: true,
			Val:    value == 1,
		}
	}
	return buildNode(0, size, 0, size)
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		grid [][]int
		node *Node
	}{
		{
			grid: [][]int{
				{0, 1},
				{1, 0},
			},
			node: &Node{
				IsLeaf: false,
				TopLeft: &Node{
					IsLeaf: true,
					Val:    false,
				},
				TopRight: &Node{
					IsLeaf: true,
					Val:    true,
				},
				BottomLeft: &Node{
					IsLeaf: true,
					Val:    true,
				},
				BottomRight: &Node{
					IsLeaf: true,
					Val:    false,
				},
			},
		},
	}
	for _, tt := range tc {
		assert.True(t, reflect.DeepEqual(construct(tt.grid), tt.node))
	}
}
