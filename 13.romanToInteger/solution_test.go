package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
type pair struct {
	key   string
	value int
}

func romanToInt(s string) int {
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
	var res int
	for _, p := range pairs {
		if s == "" {
			break
		}
		for strings.HasPrefix(s, p.key) {
			res += p.value
			s = s[len(p.key):]
		}
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(t, 3, romanToInt("III"))
	assert.Equal(t, 4, romanToInt("IV"))
	assert.Equal(t, 9, romanToInt("IX"))
	assert.Equal(t, 58, romanToInt("LVIII"))
	assert.Equal(t, 1994, romanToInt("MCMXCIV"))
}
