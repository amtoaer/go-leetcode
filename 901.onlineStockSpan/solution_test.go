/*
 * @lc app=leetcode.cn id=901 lang=golang
 * @lcpr version=30202
 *
 * [901] 股票价格跨度
 */

package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// @lc code=start
type StockSpanner struct {
	stack []struct {
		idx, price int
	}
	idx int
}

func Constructor() StockSpanner {
	return StockSpanner{
		stack: []struct {
			idx, price int
		}{
			{-1, math.MaxInt},
		},
		idx: -1,
	}
}

func (s *StockSpanner) Next(price int) int {
	s.idx++
	for s.stack[len(s.stack)-1].price <= price {
		s.stack = s.stack[:len(s.stack)-1]
	}
	res := s.idx - s.stack[len(s.stack)-1].idx
	s.stack = append(s.stack, struct {
		idx   int
		price int
	}{
		idx:   s.idx,
		price: price,
	})
	return res
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
// @lc code=end

func Test(t *testing.T) {
	stockSpanner := Constructor()
	tc := []struct {
		price int
		want  int
	}{
		{100, 1},
		{80, 1},
		{60, 1},
		{70, 2},
		{60, 1},
		{75, 4},
		{85, 6},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, stockSpanner.Next(tt.price))
	}
}

/*
// @lcpr case=start
// ["StockSpanner", "next", "next", "next", "next", "next", "next", "next"]\n[[], [100], [80], [60], [70], [60], [75], [85]]\n
// @lcpr case=end

*/
