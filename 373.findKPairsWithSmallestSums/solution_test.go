package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=373 lang=golang
 *
 * [373] Find K Pairs with Smallest Sums
 */

// @lc code=start
type Pair struct {
	i, j int
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	getCompareValue := func(p Pair) int {
		return nums1[p.i] + nums2[p.j]
	}

	siftDown := func(pairs []Pair, start, end int) {
		cur, next := start, start*2+1
		for next <= end {
			if next+1 <= end && getCompareValue(pairs[next+1]) < getCompareValue(pairs[next]) {
				next++
			}
			if getCompareValue(pairs[next]) >= getCompareValue(pairs[cur]) {
				return
			}
			pairs[cur], pairs[next] = pairs[next], pairs[cur]
			cur, next = next, next*2+1
		}
	}

	siftUp := func(pairs []Pair, idx int) {
		for idx > 0 {
			parent := (idx - 1) / 2
			if getCompareValue(pairs[idx]) >= getCompareValue(pairs[parent]) {
				break
			}
			pairs[idx], pairs[parent] = pairs[parent], pairs[idx]
			idx = parent
		}
	}

	visited := make(map[Pair]struct{})
	heap := []Pair{{0, 0}}
	res := make([][]int, 0, k)
	for len(heap) > 0 && len(res) < k {
		poped := heap[0]
		res = append(res, []int{nums1[poped.i], nums2[poped.j]})
		heap[0], heap[len(heap)-1] = heap[len(heap)-1], heap[0]
		heap = heap[:len(heap)-1]
		siftDown(heap, 0, len(heap)-1)
		leftAdd, rightAdd := poped, poped
		leftAdd.i++
		rightAdd.j++
		if _, ok := visited[leftAdd]; !ok && leftAdd.i < len(nums1) {
			visited[leftAdd] = struct{}{}
			heap = append(heap, leftAdd)
			siftUp(heap, len(heap)-1)
		}
		if _, ok := visited[rightAdd]; !ok && rightAdd.j < len(nums2) {
			visited[rightAdd] = struct{}{}
			heap = append(heap, rightAdd)
			siftUp(heap, len(heap)-1)
		}
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		nums1 []int
		nums2 []int
		k     int
		want  [][]int
	}{
		{[]int{1, 7, 11}, []int{2, 4, 6}, 3, [][]int{{1, 2}, {1, 4}, {1, 6}}},
		{[]int{1, 1, 2}, []int{1, 2, 3}, 2, [][]int{{1, 1}, {1, 1}}},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, kSmallestPairs(tt.nums1, tt.nums2, tt.k))
	}
}
