package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 */

// @lc code=start
func siftDown(nums []int, start, end int) {
	cur, next := start, start*2+1
	for next <= end {
		if next+1 <= end && nums[next+1] > nums[next] {
			next++
		}
		if nums[next] <= nums[cur] {
			return
		}
		nums[cur], nums[next] = nums[next], nums[cur]
		cur, next = next, next*2+1
	}
}

func topKFrequent(nums []int, k int) []int {
	cnt := make(map[int]int)
	for _, num := range nums {
		cnt[num]++
	}
	visited := make(map[int]struct{})
	var cnts []int
	cntToNum := make(map[int][]int)
	res := make([]int, 0, k)
	for k, v := range cnt {
		if _, ok := visited[v]; !ok {
			cnts = append(cnts, v)
			visited[v] = struct{}{}
		}
		cntToNum[v] = append(cntToNum[v], k)
	}
	for i := len(cnts)/2 - 1; i >= 0; i-- {
		siftDown(cnts, i, len(cnts)-1)
	}
	i := 0
	for len(res) < k {
		res = append(res, cntToNum[cnts[0]]...)
		cnts[0], cnts[len(cnts)-1-i] = cnts[len(cnts)-1-i], cnts[0]
		siftDown(cnts, 0, len(cnts)-2-i)
		i++
	}
	return res[:k]
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{[]int{1}, 1, []int{1}},
		{[]int{1, 2}, 2, []int{1, 2}},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6}, 10, []int{2, 1, 5, 6, 7, 3, 8, 4, 10, 11}},
	}
	for _, tt := range tc {
		res := topKFrequent(tt.nums, tt.k)
		assert.ElementsMatch(t, tt.expected, res)
	}
}
