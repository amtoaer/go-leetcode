package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=134 lang=golang
 *
 * [134] Gas Station
 */

// @lc code=start
func canCompleteCircuit(gas []int, cost []int) int {
	var (
		start = 0
		l     = len(gas)
	)
	for start < l {
		var gasSum, costSum, cnt int
		for cnt < l {
			j := (start + cnt) % l
			gasSum += gas[j]
			costSum += cost[j]
			if gasSum < costSum {
				break
			}
			cnt++
		}
		if cnt == l {
			return start
		}
		start += cnt + 1
	}
	return -1
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		gas  []int
		cost []int
		want int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}, 3},
		{[]int{2, 3, 4}, []int{3, 4, 3}, -1},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, canCompleteCircuit(tt.gas, tt.cost))
	}
}
