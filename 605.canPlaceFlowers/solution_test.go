package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=605 lang=golang
 *
 * [605] Can Place Flowers
 */

// @lc code=start
func canPlaceFlowers(flowerbed []int, n int) bool {
	var c int
	if c >= n {
		return true
	}
	for idx, val := range flowerbed {
		if (idx == 0 || flowerbed[idx-1] == 0) && val == 0 && (idx == len(flowerbed)-1 || flowerbed[idx+1] == 0) {
			c += 1
			flowerbed[idx] = 1
			if c >= n {
				return true
			}
		}
	}
	return false
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		flowerbed []int
		n         int
		want      bool
	}{
		{[]int{1, 0, 0, 0, 1}, 1, true},
		{[]int{1, 0, 0, 0, 1}, 2, false},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, canPlaceFlowers(tt.flowerbed, tt.n))
	}
}
