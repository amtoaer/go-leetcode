package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=735 lang=golang
 *
 * [735] Asteroid Collision
 */

// @lc code=start
func abs(v int) int {
	if v > 0 {
		return v
	}
	return -v
}

func hasSameSign(a, b int) bool {
	return a^b >= 0
}

func asteroidCollision(asteroids []int) []int {
	var stack []int
	for _, asteroid := range asteroids {
		if len(stack) == 0 || hasSameSign(stack[len(stack)-1], asteroid) {
			stack = append(stack, asteroid)
		} else {
			needAppend := true
			for len(stack) > 0 {
				top := stack[len(stack)-1]
				if hasSameSign(top, asteroid) || (top < 0 && asteroid > 0) {
					break
				} else if abs(top) == abs(asteroid) {
					needAppend = false
					stack = stack[:len(stack)-1]
					break
				} else if abs(top) > abs(asteroid) {
					needAppend = false
					break
				} else {
					stack = stack[:len(stack)-1]
				}
			}
			if needAppend {
				stack = append(stack, asteroid)
			}
		}
	}
	return stack
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		asteroids []int
		want      []int
	}{
		{[]int{5, 10, -5}, []int{5, 10}},
		{[]int{8, -8}, []int{}},
		{[]int{10, 2, -5}, []int{10}},
		{[]int{-2, -1, 1, 2}, []int{-2, -1, 1, 2}},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, asteroidCollision(tt.asteroids))
	}
}
