package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=649 lang=golang
 *
 * [649] Dota2 Senate
 */

// @lc code=start
func predictPartyVictory(senate string) string {
	var rQueue, dQueue []int
	for idx, item := range senate {
		if item == 'R' {
			rQueue = append(rQueue, idx)
		} else {
			dQueue = append(dQueue, idx)
		}
	}
	for len(rQueue) > 0 && len(dQueue) > 0 {
		if rQueue[0] < dQueue[0] {
			rQueue = append(rQueue, rQueue[0]+len(senate))
		} else {
			dQueue = append(dQueue, dQueue[0]+len(senate))
		}
		rQueue = rQueue[1:]
		dQueue = dQueue[1:]
	}
	if len(rQueue) > 0 {
		return "Radiant"
	}
	return "Dire"
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  string
		output string
	}{
		{"RD", "Radiant"},
		{"RDD", "Dire"},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, predictPartyVictory(tt.input))
	}
}
