package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=2332 lang=golang
 *
 * [2332] The Latest Time to Catch a Bus
 */

// @lc code=start
func latestTimeCatchTheBus(buses []int, passengers []int, capacity int) int {
	sort.Ints(buses)
	sort.Ints(passengers)
	var pos, space int
	for _, busArrived := range buses {
		space = capacity
		for space > 0 && pos < len(passengers) && passengers[pos] <= busArrived {
			space--
			pos++
		}
	}
	// pos -- 后，pos 指向最后一个上车的乘客
	pos--
	// 有空位，那么最后的上车时间可以假设为最后一辆车开来的时间
	res := buses[len(buses)-1]
	// 没有空位，最后的上车时间假设为最后一个赶上车的乘客的时间
	if space <= 0 {
		res = passengers[pos]
	}
	// 接下来就要从后往前找，找到最后一个不和现有乘客重合的时间，我们在这个时间上车可以把最后一个乘客挤掉
	for pos >= 0 && passengers[pos] == res {
		pos--
		res--
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		buses      []int
		passengers []int
		capacity   int
		expected   int
	}{
		{[]int{10, 20}, []int{2, 17, 18, 19}, 2, 16},
		{[]int{20, 30, 10}, []int{19, 13, 26, 4, 25, 11, 21}, 2, 20},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.expected, latestTimeCatchTheBus(tt.buses, tt.passengers, tt.capacity))
	}
}
