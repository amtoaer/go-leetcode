package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=1103 lang=golang
 *
 * [1103] Distribute Candies to People
 */

// @lc code=start
func distributeCandies(candies int, num_people int) []int {
	n := num_people
	p := int(math.Sqrt(float64(2*candies)+0.25) - 0.5)
	remaining := candies - int(float64((p+1)*p)*0.5)
	rows, cols := p/n, p%n

	d := make([]int, n)
	for i := 0; i < n; i++ {
		d[i] = (i+1)*rows + int(float64(rows*(rows-1)*n)*0.5)
		if i < cols {
			d[i] += i + 1 + rows*n
		}
	}
	d[cols] += remaining
	return d
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 1}, distributeCandies(7, 4))
	assert.Equal(t, []int{5, 2, 3}, distributeCandies(10, 3))
}
