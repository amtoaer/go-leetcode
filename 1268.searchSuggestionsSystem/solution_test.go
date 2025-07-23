/*
 * @lc app=leetcode.cn id=1268 lang=golang
 * @lcpr version=30202
 *
 * [1268] 搜索推荐系统
 */

package main

import (
	"strings"
	"testing"

	"github.com/emirpasic/gods/v2/queues/priorityqueue"
	"github.com/stretchr/testify/assert"
)

// @lc code=start
type PrefixTree struct {
	next  [26]*PrefixTree
	words *priorityqueue.Queue[string]
}

func NewPrefixTree() *PrefixTree {
	return &PrefixTree{
		words: priorityqueue.NewWith(func(x, y string) int {
			if x > y {
				return -1
			}
			return 1
		}),
	}
}

func (p *PrefixTree) Insert(word string) {
	cur := p
	for _, r := range word {
		idx := r - 'a'
		if cur.next[idx] == nil {
			cur.next[idx] = NewPrefixTree()
		}
		cur = cur.next[idx]
		cur.words.Enqueue(word)
		for cur.words.Size() > 3 {
			cur.words.Dequeue()
		}
	}
}

func (p *PrefixTree) Search(word string) []string {
	cur := p
	for _, r := range word {
		idx := r - 'a'
		if cur.next[idx] == nil {
			return []string(nil)
		}
		cur = cur.next[idx]
	}
	res := cur.words.Values()
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func suggestedProducts(products []string, searchWord string) [][]string {
	tree := NewPrefixTree()
	for _, product := range products {
		tree.Insert(product)
	}
	res := make([][]string, 0, len(searchWord))
	var sb strings.Builder
	for _, r := range searchWord {
		sb.WriteRune(r)
		res = append(res, tree.Search(sb.String()))
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		products   []string
		searchWord string
		want       [][]string
	}{
		{
			[]string{"mobile", "mouse", "moneypot", "monitor", "mousepad"},
			"mouse",
			[][]string{
				{"mobile", "moneypot", "monitor"},
				{"mobile", "moneypot", "monitor"},
				{"mouse", "mousepad"},
				{"mouse", "mousepad"},
				{"mouse", "mousepad"},
			},
		},
		{
			[]string{"havana"},
			"havana",
			[][]string{
				{"havana"},
				{"havana"},
				{"havana"},
				{"havana"},
				{"havana"},
				{"havana"},
			},
		},
		{
			[]string{"bags", "baggage", "banner", "box", "cloths"},
			"bags",
			[][]string{
				{"baggage", "bags", "banner"},
				{"baggage", "bags", "banner"},
				{"baggage", "bags"},
				{"bags"},
			},
		},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, suggestedProducts(tt.products, tt.searchWord))
	}
}

/*
// @lcpr case=start
// ["mobile","mouse","moneypot","monitor","mousepad"]\n"mouse"\n
// @lcpr case=end

// @lcpr case=start
// ["havana"]\n"havana"\n
// @lcpr case=end

// @lcpr case=start
// ["bags","baggage","banner","box","cloths"]\n"bags"\n
// @lcpr case=end

// @lcpr case=start
// ["havana"]\n"tatiana"\n
// @lcpr case=end

*/
