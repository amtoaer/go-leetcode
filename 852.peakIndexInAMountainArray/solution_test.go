package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=852 lang=golang
 *
 * [852] Peak Index in a Mountain Array
 */

// @lc code=start
type Direction int

const (
	LEFT = iota
	RIGHT
)

func peakIndexInMountainArray(arr []int) int {
	greater := func(idx int, d Direction) bool {
		return (d == LEFT && (idx == 0 || arr[idx] > arr[idx-1])) ||
			(d == RIGHT && (idx == len(arr)-1 || arr[idx] > arr[idx+1]))
	}

	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if greater(mid, LEFT) && greater(mid, RIGHT) {
			return mid
		} else if greater(mid, LEFT) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return 0
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		arr []int
		ans int
	}{
		{[]int{0, 1, 0}, 1},
		{[]int{0, 2, 1, 0}, 1},
		{[]int{0, 10, 5, 2}, 1},
		{[]int{3, 4, 5, 1}, 2},
	}
	for _, tt := range tc {
		assert.Equal(t, peakIndexInMountainArray(tt.arr), tt.ans)
	}
}
