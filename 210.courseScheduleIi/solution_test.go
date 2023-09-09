package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=210 lang=golang
 *
 * [210] Course Schedule II
 */

// @lc code=start
func findOrder(numCourses int, prerequisites [][]int) []int {
	outEdges := make([][]int, numCourses)
	inEdgeCounts := make([]int, numCourses)
	for _, outEdge := range prerequisites {
		outEdges[outEdge[1]] = append(outEdges[outEdge[1]], outEdge[0])
		inEdgeCounts[outEdge[0]]++
	}
	var nodeQueue []int
	res := make([]int, 0, numCourses)
	for node, inEdgeCount := range inEdgeCounts {
		if inEdgeCount == 0 {
			nodeQueue = append(nodeQueue, node)
		}
	}
	for len(nodeQueue) != 0 {
		res = append(res, nodeQueue...)
		length := len(nodeQueue)
		for _, node := range nodeQueue[:length] {
			for _, targetNode := range outEdges[node] {
				inEdgeCounts[targetNode]--
				if inEdgeCounts[targetNode] == 0 {
					nodeQueue = append(nodeQueue, targetNode)
				}
			}
		}
		nodeQueue = nodeQueue[length:]
	}
	if len(res) != numCourses {
		return []int{}
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(t, []int{0, 1}, findOrder(2, [][]int{{1, 0}}))
	assert.Equal(t, []int{0, 1, 2, 3}, findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
	assert.Equal(t, []int{}, findOrder(2, [][]int{{1, 0}, {0, 1}}))
}
