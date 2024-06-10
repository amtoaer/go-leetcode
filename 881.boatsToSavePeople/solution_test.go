package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=881 lang=golang
 *
 * [881] Boats to Save People
 */

// @lc code=start
func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	left, right := 0, len(people)-1
	res := 0
	for left <= right {
		if people[left]+people[right] > limit {
			right--
		} else {
			left++
			right--
		}
		res++
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	testcases := []struct {
		people []int
		limit  int
		want   int
	}{
		{people: []int{1, 2}, limit: 3, want: 1},
		{people: []int{3, 2, 2, 1}, limit: 3, want: 3},
		{people: []int{3, 5, 3, 4}, limit: 5, want: 4},
	}
	for _, tc := range testcases {
		assert.Equal(t, tc.want, numRescueBoats(tc.people, tc.limit))
	}
}
