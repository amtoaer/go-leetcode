package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=933 lang=golang
 *
 * [933] Number of Recent Calls
 */

// @lc code=start
type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (r *RecentCounter) Ping(t int) int {
	r.queue = append(r.queue, t)
	var x int
	for r.queue[x] < t-3000 {
		x++
	}
	r.queue = r.queue[x:]
	return len(r.queue)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
// @lc code=end

func Test(t *testing.T) {
	obj := Constructor()
	assert.Equal(t, 1, obj.Ping(1))
	assert.Equal(t, 2, obj.Ping(100))
	assert.Equal(t, 3, obj.Ping(3001))
	assert.Equal(t, 3, obj.Ping(3002))
}
