package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=2594 lang=golang
 *
 * [2594] Minimum Time to Repair Cars
 */

// @lc code=start
func repairCars(ranks []int, cars int) int64 {
	rankCount := make([]int, 100)
	minRank := math.MaxInt
	for _, rank := range ranks {
		rankCount[rank-1]++
		if rank < minRank {
			minRank = rank
		}
	}
	check := func(time int) bool {
		var res int
		for rank, cnt := range rankCount {
			res += cnt * int(math.Sqrt(float64(time/(rank+1))))
		}
		return res >= cars
	}
	left, right := 0, ranks[0]*cars*cars
	for left < right {
		mid := (left + right) >> 1
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return int64(left)
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(t, int64(16), repairCars([]int{4, 2, 3, 1}, 10))
	assert.Equal(t, int64(16), repairCars([]int{5, 1, 8}, 6))
}
