package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=165 lang=golang
 *
 * [165] Compare Version Numbers
 */

// @lc code=start
func compareVersion(version1 string, version2 string) int {
	next := func(version string, idx int) (int, int) {
		for i := idx; i <= len(version); i++ {
			if i == len(version) || version[i] == '.' {
				v, _ := strconv.Atoi(version[idx:i])
				return v, i + 1
			}
		}
		return 0, idx
	}

	var i, j, verLeft, verRight int
	for i < len(version1) || j < len(version2) {
		verLeft, i = next(version1, i)
		verRight, j = next(version2, j)
		if verLeft > verRight {
			return 1
		} else if verLeft < verRight {
			return -1
		}
	}
	return 0
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		version1 string
		version2 string
		want     int
	}{
		{"1.2", "1.10", -1},
		{"1.01", "1.001", 0},
		{"1.0", "1.0.0.0", 0},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, compareVersion(tt.version1, tt.version2))
	}
}
