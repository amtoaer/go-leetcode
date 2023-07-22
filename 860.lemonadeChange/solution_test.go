package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=860 lang=golang
 *
 * [860] Lemonade Change
 */

// @lc code=start
func lemonadeChange(bills []int) bool {
	var five, ten int
	for _, item := range bills {
		switch item {
		case 5:
			five += 1
		case 10:
			if five == 0 {
				return false
			}
			five--
			ten++
		case 20:
			if ten > 0 && five > 0 {
				ten--
				five--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(t, true, lemonadeChange([]int{5, 5, 5, 10, 20}))
	assert.Equal(t, true, lemonadeChange([]int{5, 5, 10}))
	assert.Equal(t, false, lemonadeChange([]int{10, 10}))
	assert.Equal(t, false, lemonadeChange([]int{5, 5, 10, 10, 20}))
}
