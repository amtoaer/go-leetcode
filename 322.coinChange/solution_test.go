package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] Coin Change
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	m := make([]int, amount+1)
	m[0] = 0
	for i := 1; i <= amount; i++ {
		cnt := math.MaxInt - 1
		for _, coin := range coins {
			x := i - coin
			if x >= 0 {
				cnt = min(cnt, m[x]+1)
			}
		}
		m[i] = cnt
	}
	cnt := m[amount]
	if cnt == math.MaxInt-1 {
		return -1
	}
	return cnt
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		coins  []int
		amount int
		want   int
	}{
		{[]int{1, 2, 5}, 11, 3},
		{[]int{2}, 3, -1},
		{[]int{1}, 0, 0},
		{[]int{1}, 1, 1},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, coinChange(tt.coins, tt.amount))
	}
}
