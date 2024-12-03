package main

import (
	"testing"
)

/*
 * @lc app=leetcode.cn id=547 lang=golang
 * @lcpr version=20004
 *
 * [547] 省份数量
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func findCircleNum(isConnected [][]int) int {
	var dfs func(int)
	visited := make([]bool, len(isConnected))
	dfs = func(i int) {
		visited[i] = true
		for nextIdx, next := range isConnected[i] {
			if next == 1 && !visited[nextIdx] {
				dfs(nextIdx)
			}
		}
	}
	var res int
	for idx := range isConnected {
		if !visited[idx] {
			res++
			dfs(idx)
		}
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [[1,1,0],[1,1,0],[0,0,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,0,0],[0,1,0],[0,0,1]]\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		isConnected [][]int
		expected    int
	}{
		{[][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}, 2},
		{[][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}, 3},
		{[][]int{{1, 1, 0, 0}, {1, 1, 0, 0}, {0, 0, 1, 1}, {0, 0, 1, 1}}, 2},
		{[][]int{{1, 0, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 1}, {1, 0, 1, 1}}, 1},
	}

	for _, test := range tests {
		result := findCircleNum(test.isConnected)
		if result != test.expected {
			t.Errorf("For isConnected = %v, expected %d, but got %d", test.isConnected, test.expected, result)
		}
	}
}
