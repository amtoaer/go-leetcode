package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=150 lang=golang
 *
 * [150] Evaluate Reverse Polish Notation
 */

// @lc code=start
func evalRPN(tokens []string) int {
	var tmp []int
	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			{

				x, y := tmp[len(tmp)-2], tmp[len(tmp)-1]
				tmp = tmp[:len(tmp)-2]
				var res int
				switch token {
				case "+":
					{
						res = x + y
					}
				case "-":
					{
						res = x - y
					}
				case "*":
					{
						res = x * y
					}
				case "/":
					{
						res = x / y
					}
				}
				tmp = append(tmp, res)
			}
		default:
			v, _ := strconv.Atoi(token)
			tmp = append(tmp, v)
		}
	}
	return tmp[0]
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		tokens []string
		expect int
	}{
		{[]string{"2", "1", "+", "3", "*"}, 9},
		{[]string{"4", "13", "5", "/", "+"}, 6},
		{[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22},
	}
	for _, tc := range tc {
		assert.Equal(t, tc.expect, evalRPN(tc.tokens))
	}
}
