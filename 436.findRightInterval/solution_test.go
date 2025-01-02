package main

import (
	"sort"
	"testing"
)

/*
 * @lc app=leetcode.cn id=436 lang=golang
 * @lcpr version=20004
 *
 * [436] 寻找右区间
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func findRightInterval(intervals [][]int) []int {
	for i := range intervals {
		intervals[i] = append(intervals[i], i)
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	binarySearch := func(v int) int {
		target := -1
		left, right := 0, len(intervals)-1
		for left <= right {
			mid := left + (right-left)>>1
			if intervals[mid][0] >= v {
				target = intervals[mid][2]
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		return target
	}
	res := make([]int, len(intervals))
	for _, interval := range intervals {
		res[interval[2]] = binarySearch(interval[1])
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [[1,2]]\n
// @lcpr case=end

// @lcpr case=start
// [[3,4],[2,3],[1,2]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,4],[2,3],[3,4]]\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		intervals [][]int
		expected  []int
	}{
		{[][]int{{1, 2}}, []int{-1}},
		{[][]int{{3, 4}, {2, 3}, {1, 2}}, []int{-1, 0, 1}},
		{[][]int{{1, 4}, {2, 3}, {3, 4}}, []int{-1, 2, -1}},
	}

	for _, test := range tests {
		result := findRightInterval(test.intervals)
		for i, v := range result {
			if v != test.expected[i] {
				t.Errorf("For intervals %v, expected %v but got %v", test.intervals, test.expected, result)
				break
			}
		}
	}
}
