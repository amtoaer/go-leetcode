package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=12 lang=golang
 *
 * [12] Integer to Roman
 */

// @lc code=start
type pair struct {
	key   string
	value int
}

func intToRoman(num int) string {
	pairs := []pair{
		{"M", 1000},
		{"CM", 900},
		{"D", 500},
		{"CD", 400},
		{"C", 100},
		{"XC", 90},
		{"L", 50},
		{"XL", 40},
		{"X", 10},
		{"IX", 9},
		{"V", 5},
		{"IV", 4},
		{"I", 1},
	}
	var sb strings.Builder
	for _, p := range pairs {
		if num == 0 {
			break
		}
		for i := 0; i < num/p.value; i++ {
			sb.WriteString(p.key)
		}
		num %= p.value
	}
	return sb.String()
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(t, "III", intToRoman(3))
	assert.Equal(t, "IV", intToRoman(4))
	assert.Equal(t, "IX", intToRoman(9))
	assert.Equal(t, "LVIII", intToRoman(58))
	assert.Equal(t, "MCMXCIV", intToRoman(1994))
}
