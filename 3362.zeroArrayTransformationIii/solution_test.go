/*
 * @lc app=leetcode.cn id=3362 lang=golang
 * @lcpr version=30201
 *
 * [3362] 零数组变换 III
 */

package main

import (
	"container/heap"
	"sort"
	"testing"
)

type Heap []int

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() any {
	last := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return last
}

// @lc code=start
func maxRemoval(nums []int, queries [][]int) int {
	sort.Slice(queries, func(i, j int) bool {
		return queries[i][0] < queries[j][0]
	})
	var queryIdx, currentVal int
	var h Heap
	heap.Init(&h)
	diffArray := make([]int, len(nums)+1)
	for i, num := range nums {
		currentVal += diffArray[i]
		for queryIdx < len(queries) && queries[queryIdx][0] == i {
			heap.Push(&h, queries[queryIdx][1])
			queryIdx++
		}
		for h.Len() > 0 && currentVal < num && h[0] >= i {
			right := heap.Pop(&h).(int)
			currentVal++
			diffArray[right+1]--
		}
		if currentVal < num {
			return -1
		}
	}
	return h.Len()
}

// @lc code=end

func Test(t *testing.T) {
	testCases := []struct {
		nums     []int
		queries  [][]int
		expected int
	}{
		{[]int{2, 0, 2}, [][]int{{0, 2}, {0, 2}, {1, 1}}, 1},
		{[]int{1, 1, 1, 1}, [][]int{{1, 3}, {0, 2}, {1, 3}, {1, 2}}, 2},
		{[]int{1, 2, 3, 4}, [][]int{{0, 3}}, -1},
	}

	for i, tc := range testCases {
		result := maxRemoval(tc.nums, tc.queries)
		if result != tc.expected {
			t.Errorf("Test case %d: Expected %d but got %d", i+1, tc.expected, result)
		}
	}
}

/*
// @lcpr case=start
// [2,0,2]\n[[0,2],[0,2],[1,1]]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,1,1]\n[[1,3],[0,2],[1,3],[1,2]]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4]\n[[0,3]]\n
// @lcpr case=end

*/
