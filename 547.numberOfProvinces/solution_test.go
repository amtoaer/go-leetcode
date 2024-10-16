package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=547 lang=golang
 * @lcpr version=20001
 *
 * [547] 省份数量
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func findCircleNum(isConnected [][]int) int {
	visited := make([]bool, len(isConnected))
	var res int
	var dfs func(int)
	dfs = func(idx int) {
		visited[idx] = true
		for i := 0; i < len(isConnected); i++ {
			if isConnected[idx][i] == 1 && !visited[i] {
				dfs(i)
			}
		}
	}

	for i := 0; i < len(isConnected); i++ {
		if !visited[i] {
			res++
			dfs(i)
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
	type Case struct {
		isConnected [][]int
		res         int
	}
	cases := []Case{
		{isConnected: [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}, res: 2},
		{isConnected: [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}, res: 3},
	}
	for _, c := range cases {
		assert.Equal(t, c.res, findCircleNum(c.isConnected))
	}
}
