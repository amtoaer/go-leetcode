package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=682 lang=golang
 *
 * [682] Baseball Game
 */

// @lc code=start
func calPoints(operations []string) int {
	stack := make([]int, 0, len(operations))
	for _, val := range operations {
		l := len(stack)
		switch val {
		case "+":
			stack = append(stack, stack[l-1]+stack[l-2])
		case "D":
			stack = append(stack, 2*stack[l-1])
		case "C":
			stack = stack[:l-1]
		default:
			numVal, _ := strconv.Atoi(val)
			stack = append(stack, numVal)
		}
	}
	sum := 0
	for _, val := range stack {
		sum += val
	}
	return sum
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  []string
		output int
	}{
		{[]string{"5", "2", "C", "D", "+"}, 30},
		{[]string{"5", "-2", "4", "C", "D", "9", "+", "+"}, 27},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, calPoints(tt.input))
	}
}
