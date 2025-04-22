/*
 * @lc app=leetcode.cn id=1399 lang=golang
 * @lcpr version=30104
 *
 * [1399] 统计最大组的数目
 */

package main

import (
	"maps"
	"testing"
)

// @lc code=start
func countLargestGroup(n int) int {
	var (
		m      = make(map[int]int)
		maxVal int
	)
	for i := range n {
		var (
			key int
			x   = i + 1
		)
		for x > 0 {
			key += x % 10
			x /= 10
		}
		m[key]++
		maxVal = max(maxVal, m[key])
	}
	var res int
	for val := range maps.Values(m) {
		if val == maxVal {
			res++
		}
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	if countLargestGroup(13) != 4 {
		t.Errorf("Expected 4, but got %d", countLargestGroup(13))
	}
	if countLargestGroup(2) != 2 {
		t.Errorf("Expected 2, but got %d", countLargestGroup(2))
	}
	if countLargestGroup(15) != 6 {
		t.Errorf("Expected 6, but got %d", countLargestGroup(15))
	}
}

/*
// @lcpr case=start
// 13\n
// @lcpr case=end

// @lcpr case=start
// 2\n
// @lcpr case=end

// @lcpr case=start
// 15\n
// @lcpr case=end

// @lcpr case=start
// 24\n
// @lcpr case=end

*/
