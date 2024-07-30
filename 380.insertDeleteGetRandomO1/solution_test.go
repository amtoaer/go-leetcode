package main

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=380 lang=golang
 *
 * [380] Insert Delete GetRandom O(1)
 */

// @lc code=start
type RandomizedSet struct {
	keys []int
	set  map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		keys: nil,
		set:  make(map[int]int),
	}
}

func (r *RandomizedSet) Insert(val int) bool {
	if _, ok := r.set[val]; ok {
		return false
	}
	r.keys = append(r.keys, val)
	r.set[val] = len(r.keys) - 1
	return true
}

func (r *RandomizedSet) Remove(val int) bool {
	idx, ok := r.set[val]
	if !ok {
		return false
	}
	r.keys[idx] = r.keys[len(r.keys)-1]
	r.set[r.keys[idx]] = idx
	r.keys = r.keys[:len(r.keys)-1]
	delete(r.set, val)
	return true
}

func (r *RandomizedSet) GetRandom() int {
	return r.keys[rand.Intn(len(r.keys))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end

func Test(t *testing.T) {
	r := Constructor()
	assert.Equal(t, false, r.Remove(0))
	assert.Equal(t, true, r.Insert(0))
	assert.Equal(t, true, r.Remove(0))
	assert.Equal(t, true, r.Insert(1))
	assert.Equal(t, false, r.Remove(2))
	assert.Equal(t, true, r.Insert(2))
	assert.Equal(t, false, r.Insert(2))
	assert.Equal(t, []int{1, 2}, r.keys)
	assert.Equal(t, map[int]int{1: 0, 2: 1}, r.set)
}
