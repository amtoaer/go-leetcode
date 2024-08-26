package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=692 lang=golang
 *
 * [692] Top K Frequent Words
 */

// @lc code=start

type Pair struct {
	word string
	cnt  int
}

func Greater(left, right *Pair) bool {
	if left.cnt != right.cnt {
		return left.cnt > right.cnt
	}
	return left.word < right.word
}

func siftDown(pairs []Pair, start, end int) {
	cur, next := start, start*2+1
	for next <= end {
		if next+1 <= end && Greater(&pairs[next+1], &pairs[next]) {
			next++
		}
		if Greater(&pairs[cur], &pairs[next]) {
			return
		}
		pairs[cur], pairs[next] = pairs[next], pairs[cur]
		cur, next = next, next*2+1
	}
}

func topKFrequent(words []string, k int) []string {
	m := make(map[string]int)
	for _, word := range words {
		m[word]++
	}
	pairs := make([]Pair, 0, len(m))
	for k, v := range m {
		pairs = append(pairs, Pair{k, v})
	}
	for i := len(pairs)/2 - 1; i >= 0; i-- {
		siftDown(pairs, i, len(pairs)-1)
	}
	for i := 0; i < k; i++ {
		end := len(pairs) - 1 - i
		pairs[0], pairs[end] = pairs[end], pairs[0]
		siftDown(pairs, 0, end-1)
	}
	res := make([]string, 0, k)
	for i := len(pairs) - 1; i >= len(pairs)-k; i-- {
		res = append(res, pairs[i].word)
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  []string
		k      int
		output []string
	}{
		{[]string{"i", "love", "leetcode", "i", "love", "coding"}, 2, []string{"i", "love"}},
		{[]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4, []string{"the", "is", "sunny", "day"}},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, topKFrequent(tt.input, tt.k))
	}
}
