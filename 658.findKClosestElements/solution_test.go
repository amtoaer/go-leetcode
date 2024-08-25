package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=658 lang=golang
 *
 * [658] Find K Closest Elements
 */

// @lc code=start
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func findClosestElements(arr []int, k int, x int) []int {
	better := func(a, b int) bool {
		aDiff, bDiff := abs(a-x), abs(b-x)
		return aDiff < bDiff || (aDiff == bDiff && a < b)
	}
	siftDown := func(start, end int) {
		cur, next := start, start*2+1
		for next <= end {
			if next+1 <= end && better(arr[next+1], arr[next]) {
				next++
			}
			if better(arr[cur], arr[next]) {
				return
			}
			arr[cur], arr[next] = arr[next], arr[cur]
			cur, next = next, next*2+1
		}
	}
	for i := len(arr)/2 - 1; i >= 0; i-- {
		siftDown(i, len(arr)-1)
	}
	for i := 0; i < k; i++ {
		arr[0], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[0]
		siftDown(0, len(arr)-2-i)
	}
	sort.Ints(arr[len(arr)-k:])
	return arr[len(arr)-k:]
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		arr  []int
		k    int
		x    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, 4, 3, []int{1, 2, 3, 4}},
		{[]int{1, 2, 3, 4, 5}, 4, -1, []int{1, 2, 3, 4}},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, findClosestElements(tt.arr, tt.k, tt.x))
	}
}
