/*
 * @lc app=leetcode.cn id=252 lang=golang
 * @lcpr version=30202
 *
 * [252] 会议室
 */

package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// @lc code=start
func canAttendMeetings(intervals [][]int) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := range len(intervals) - 1 {
		if intervals[i][1] > intervals[i+1][0] {
			return false
		}
	}
	return true
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		intervals [][]int
		want      bool
	}{
		{[][]int{{0, 30}, {5, 10}, {15, 20}}, false},
		{[][]int{{7, 10}, {2, 4}}, true},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, canAttendMeetings(tt.intervals))
	}
}

/*
// @lcpr case=start
// [[0,30],[5,10],[15,20]]\n
// @lcpr case=end

// @lcpr case=start
// [[7,10],[2,4]]\n
// @lcpr case=end

*/
