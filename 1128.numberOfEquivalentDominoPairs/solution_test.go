/*
 * @lc app=leetcode.cn id=1128 lang=golang
 * @lcpr version=30200
 *
 * [1128] 等价多米诺骨牌对的数量
 */

package main

import (
	"maps"
	"slices"
	"testing"
)

// @lc code=start
func hash(item []int) int {
	a, b := item[0], item[1]
	if a < b {
		a, b = b, a
	}
	return a*10 + b
}

func numEquivDominoPairs(dominoes [][]int) int {
	var (
		res int
		m   = make(map[int]int)
	)

	for item := range slices.Values(dominoes) {
		m[hash(item)]++
	}
	for val := range maps.Values(m) {
		if val > 1 {
			res += val * (val - 1) / 2
		}
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {

	tests := []struct {
		dominoes [][]int
		expected int
	}{
		{
			dominoes: [][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}},
			expected: 1,
		},
		{
			dominoes: [][]int{{1, 2}, {1, 2}, {1, 1}, {1, 2}, {2, 2}},
			expected: 3,
		},
	}

	for _, test := range tests {
		result := numEquivDominoPairs(test.dominoes)
		if result != test.expected {
			t.Errorf("For dominoes %v, expected %d but got %d", test.dominoes, test.expected, result)
		}
	}
}

/*
// @lcpr case=start
// [[1,2],[2,1],[3,4],[5,6]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2],[1,2],[1,1],[1,2],[2,2]]\n
// @lcpr case=end

*/
