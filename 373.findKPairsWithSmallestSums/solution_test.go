package main

import (
	"testing"
)

/*
 * @lc app=leetcode.cn id=373 lang=golang
 * @lcpr version=20004
 *
 * [373] 查找和最小的 K 对数字
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func siftDown(heap [][2]int, start, end int, greater func(int, int) bool) {
	cur, next := start, start*2+1
	for next <= end {
		if next+1 <= end && greater(next, next+1) {
			next = next + 1
		}
		if greater(next, cur) {
			break
		}
		heap[cur], heap[next] = heap[next], heap[cur]
		cur, next = next, next*2+1
	}
}

func siftUp(heap [][2]int, cur int, greater func(int, int) bool) {
	for cur != 0 {
		next := (cur - 1) / 2
		if greater(cur, next) {
			break
		}
		heap[next], heap[cur] = heap[cur], heap[next]
		cur = next
	}
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) (res [][]int) {
	heap := [][2]int{{0, 0}}
	visited := map[[2]int]struct{}{
		{0, 0}: {},
	}
	greater := func(i, j int) bool {
		heapI, heapJ := heap[i], heap[j]
		return nums1[heapI[0]]+nums2[heapI[1]] > nums1[heapJ[0]]+nums2[heapJ[1]]
	}
	for len(res) < k && len(heap) > 0 {
		// 将最小值加入结果集中
		minPair := heap[0]
		res = append(res, []int{nums1[minPair[0]], nums2[minPair[1]]})
		// 移除最小值并调整堆
		heap[0], heap[len(heap)-1] = heap[len(heap)-1], heap[0]
		heap = heap[:len(heap)-1]
		siftDown(heap, 0, len(heap)-1, greater)
		// 尝试拓展当前的堆
		next1, next2 := [2]int{minPair[0] + 1, minPair[1]}, [2]int{minPair[0], minPair[1] + 1}
		if _, ok := visited[next1]; !ok && next1[0] < len(nums1) && next1[1] < len(nums2) {
			heap = append(heap, next1)
			siftUp(heap, len(heap)-1, greater)
			visited[next1] = struct{}{}
		}
		if _, ok := visited[next2]; !ok && next2[0] < len(nums1) && next2[1] < len(nums2) {
			heap = append(heap, next2)
			siftUp(heap, len(heap)-1, greater)
			visited[next2] = struct{}{}
		}
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [1,7,11]\n[2,4,6]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1,1,2]\n[1,2,3]\n2\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		k     int
		want  [][]int
	}{
		{[]int{1, 7, 11}, []int{2, 4, 6}, 3, [][]int{{1, 2}, {1, 4}, {1, 6}}},
		{[]int{1, 1, 2}, []int{1, 2, 3}, 2, [][]int{{1, 1}, {1, 1}}},
		{[]int{1, 2}, []int{3}, 3, [][]int{{1, 3}, {2, 3}}},
		{[]int{1, 1, 2}, []int{1, 2, 3}, 5, [][]int{{1, 1}, {1, 1}, {2, 1}, {1, 2}, {1, 2}}},
		{[]int{1, 2, 3}, []int{1, 2, 3}, 4, [][]int{{1, 1}, {1, 2}, {2, 1}, {1, 3}}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := kSmallestPairs(tt.nums1, tt.nums2, tt.k)
			if !equal(got, tt.want) {
				t.Errorf("kSmallestPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equal(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) || a[i][0]+a[i][1] != b[i][0]+b[i][1] {
			return false
		}
	}
	return true
}
