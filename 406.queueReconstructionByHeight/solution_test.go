package main

import (
	"sort"
	"testing"
)

/*
 * @lc app=leetcode.cn id=406 lang=golang
 * @lcpr version=20004
 *
 * [406] 根据身高重建队列
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func reconstructQueue(people [][]int) [][]int {
	res := make([][]int, len(people))
	sort.Slice(people, func(i, j int) bool {
		return people[i][0] < people[j][0]
	})
	for _, p := range people {
		loc := p[1]
		for idx, pos := range res {
			if loc == 0 {
				if pos == nil {
					res[idx] = p
					break
				}
				continue
			}
			if pos == nil || pos[0] == p[0] {
				loc--
			}
		}
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]\n
// @lcpr case=end

// @lcpr case=start
// [[6,0],[5,0],[4,0],[3,2],[2,2],[1,4]]\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected [][]int
	}{
		{
			input:    [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}},
			expected: [][]int{{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 4}, {7, 1}},
		},
		{
			input:    [][]int{{6, 0}, {5, 0}, {4, 0}, {3, 2}, {2, 2}, {1, 4}},
			expected: [][]int{{4, 0}, {5, 0}, {2, 2}, {3, 2}, {1, 4}, {6, 0}},
		},
	}

	for _, test := range tests {
		result := reconstructQueue(test.input)
		for i := range result {
			if result[i][0] != test.expected[i][0] || result[i][1] != test.expected[i][1] {
				t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, result)
			}
		}
	}
}
