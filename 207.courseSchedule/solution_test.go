package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] Course Schedule
 */

// @lc code=start
func canFinish(numCourses int, prerequisites [][]int) bool {
	var validCount int
	outEdges := make([][]int, numCourses)
	inEdgeCounts := make([]int, numCourses)
	for _, outEdge := range prerequisites {
		outEdges[outEdge[0]] = append(outEdges[outEdge[0]], outEdge[1])
		inEdgeCounts[outEdge[1]]++
	}
	var nodeQueue []int
	for node, inEdgeCount := range inEdgeCounts {
		if inEdgeCount == 0 {
			nodeQueue = append(nodeQueue, node)
		}
	}
	for len(nodeQueue) != 0 {
		length := len(nodeQueue)
		validCount += length
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
	return validCount == numCourses
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(t, true, canFinish(2, [][]int{{1, 0}}))
	assert.Equal(t, false, canFinish(2, [][]int{{1, 0}, {0, 1}}))
	assert.Equal(t, true, canFinish(3, [][]int{{1, 0}, {2, 0}}))
}
