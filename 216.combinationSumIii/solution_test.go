/*
 * @lc app=leetcode.cn id=216 lang=golang
 * @lcpr version=30200
 *
 * [216] 组合总和 III
 */

package main

import (
	"slices"
	"testing"
)

// @lc code=start
func combinationSum3(k int, n int) [][]int {
	var (
		res       [][]int
		tmpSum    int
		tmp       = make([]int, 0, k)
		backtrack func(start int)
	)
	backtrack = func(start int) {
		if len(tmp) == k && tmpSum == n {
			res = append(res, slices.Clone(tmp))
			return
		}
		if len(tmp) >= k || tmpSum >= n {
			return
		}
		for i := start; i <= 9; i++ {
			tmp = append(tmp, i)
			tmpSum += i
			backtrack(i + 1)
			tmpSum -= i
			tmp = tmp[:len(tmp)-1]
		}
	}
	backtrack(1)
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tests := []struct {
		k      int
		n      int
		output [][]int
	}{
		{k: 3, n: 7, output: [][]int{{1, 2, 4}}},
		{k: 3, n: 9, output: [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}}},
		{k: 4, n: 1, output: [][]int{}},
		{k: 9, n: 45, output: [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9}}},
	}

	customSliceEqual := func(a, b [][]int) bool {
		if len(a) != len(b) {
			return false
		}
		for i := range a {
			if !slices.Equal(a[i], b[i]) {
				return false
			}
		}
		return true
	}

	for _, test := range tests {
		result := combinationSum3(test.k, test.n)
		if !customSliceEqual(result, test.output) {
			t.Errorf("For k=%d, n=%d, expected %v, but got %v", test.k, test.n, test.output, result)
		}
	}
}

/*
// @lcpr case=start
// 3\n7\n
// @lcpr case=end

// @lcpr case=start
// 3\n9\n
// @lcpr case=end

// @lcpr case=start
// 4\n1\n
// @lcpr case=end

*/
