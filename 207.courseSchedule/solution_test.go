package main

import (
	"testing"
)

/*
 * @lc app=leetcode.cn id=207 lang=golang
 * @lcpr version=20004
 *
 * [207] 课程表
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func canFinish(numCourses int, prerequisites [][]int) bool {
	prerequisitesCnt := make([]int, numCourses)
	nextCourses := make([][]int, numCourses)
	for _, dep := range prerequisites {
		prerequisitesCnt[dep[0]]++
		nextCourses[dep[1]] = append(nextCourses[dep[1]], dep[0])
	}
	var validCnt int
	var queue []int
	for idx, cnt := range prerequisitesCnt {
		if cnt == 0 {
			queue = append(queue, idx)
		}
	}
	for len(queue) > 0 {
		length := len(queue)
		validCnt += length
		for _, course := range queue[:length] {
			for _, nextIdx := range nextCourses[course] {
				prerequisitesCnt[nextIdx]--
				if prerequisitesCnt[nextIdx] == 0 {
					queue = append(queue, nextIdx)
				}
			}
		}
		queue = queue[length:]
	}
	return validCnt == numCourses
}

// @lc code=end

/*
// @lcpr case=start
// 2\n[[1,0]]\n
// @lcpr case=end

// @lcpr case=start
// 2\n[[1,0],[0,1]]\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		numCourses    int
		prerequisites [][]int
		expected      bool
	}{
		{2, [][]int{{0, 1}}, true},
		{2, [][]int{{1, 0}}, true},
		{2, [][]int{{1, 0}, {0, 1}}, false},
		{4, [][]int{{1, 0}, {2, 1}, {3, 2}}, true},
		{3, [][]int{{1, 0}, {2, 1}, {0, 2}}, false},
		{5, [][]int{{1, 0}, {2, 1}, {3, 2}, {4, 3}}, true},
	}

	for _, test := range tests {
		result := canFinish(test.numCourses, test.prerequisites)
		if result != test.expected {
			t.Errorf("canFinish(%d, %v) = %v; expected %v", test.numCourses, test.prerequisites, result, test.expected)
		}
	}
}
