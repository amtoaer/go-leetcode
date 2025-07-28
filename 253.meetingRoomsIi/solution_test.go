/*
 * @lc app=leetcode.cn id=253 lang=golang
 * @lcpr version=30202
 *
 * [253] 会议室 II
 */

package main

import (
	"container/heap"
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// @lc code=start
type Heap []int

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Pop() any {
	last := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return last
}

func (h *Heap) Push(item any) {
	*h = append(*h, item.(int))
}

func minMeetingRooms(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	h := Heap([]int{math.MinInt})
	heap.Init(&h)
	for _, interval := range intervals {
		if interval[0] >= h[0] {
			heap.Pop(&h)
		}
		heap.Push(&h, interval[1])
	}
	return h.Len()
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		intervals [][]int
		want      int
	}{
		{[][]int{{0, 30}, {5, 10}, {15, 20}}, 2},
		{[][]int{{7, 10}, {2, 4}}, 1},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, minMeetingRooms(tt.intervals))
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
