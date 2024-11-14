package main

import (
	"testing"
)

/*
 * @lc app=leetcode.cn id=3249 lang=golang
 * @lcpr version=20003
 *
 * [3249] 统计好节点的数目
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func countGoodNodes(edges [][]int) int {
	n := len(edges) + 1
	adjList := make([][]int, n)
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		adjList[x] = append(adjList[x], y)
		adjList[y] = append(adjList[y], x)
	}
	var finalRes int
	var dfs func(cur, parent int) int
	dfs = func(cur, parent int) int {
		var treeSize, childSize int
		isValid := true
		for _, next := range adjList[cur] {
			if next == parent {
				continue
			}
			size := dfs(next, cur)
			if childSize == 0 {
				childSize = size
			} else if size != childSize {
				isValid = false
			}
			treeSize += size
		}
		if isValid {
			finalRes++
		}
		return treeSize + 1
	}
	dfs(0, -1)
	return finalRes
}

// @lc code=end

/*
// @lcpr case=start
// [[0,1],[0,2],[1,3],[1,4],[2,5],[2,6]]\n
// @lcpr case=end

// @lcpr case=start
// [[0,1],[1,2],[2,3],[3,4],[0,5],[1,6],[2,7],[3,8]]\n
// @lcpr case=end

// @lcpr case=start
// [[0,1],[1,2],[1,3],[1,4],[0,5],[5,6],[6,7],[7,8],[0,9],[9,10],[9,12],[10,11]]\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		edges    [][]int
		expected int
	}{
		{
			edges:    [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6}},
			expected: 7,
		},
		{
			edges:    [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {0, 5}, {1, 6}, {2, 7}, {3, 8}},
			expected: 6,
		},
		{
			edges:    [][]int{{0, 1}, {1, 2}, {1, 3}, {1, 4}, {0, 5}, {5, 6}, {6, 7}, {7, 8}, {0, 9}, {9, 10}, {9, 12}, {10, 11}},
			expected: 12,
		},
	}

	for _, test := range tests {
		result := countGoodNodes(test.edges)
		if result != test.expected {
			t.Errorf("For edges %v, expected %d, but got %d", test.edges, test.expected, result)
		}
	}
}
