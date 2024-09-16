package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=1184 lang=golang
 *
 * [1184] Distance Between Bus Stops
 */

// @lc code=start
func distanceBetweenBusStops(distance []int, start int, destination int) int {
	if start > destination {
		start, destination = destination, start
	}
	var sum, totalSum int
	for i, d := range distance {
		if i >= start && i < destination {
			sum += d
		}
		totalSum += d
	}
	return min(sum, totalSum-sum)
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		distance    []int
		start       int
		destination int
		want        int
	}{
		{[]int{1, 2, 3, 4}, 0, 1, 1},
		{[]int{1, 2, 3, 4}, 0, 2, 3},
		{[]int{1, 2, 3, 4}, 0, 3, 4},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, distanceBetweenBusStops(tt.distance, tt.start, tt.destination))
	}
}
