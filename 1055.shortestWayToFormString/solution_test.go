/*
 * @lc app=leetcode.cn id=1055 lang=golang
 * @lcpr version=30202
 *
 * [1055] 形成字符串的最短路径
 */

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// @lc code=start
func shortestWay(source string, target string) int {
	var (
		round       int
		left, right int
		existed     = make(map[byte]struct{})
	)
	for i := range len(source) {
		existed[source[i]] = struct{}{}
	}
	for i := range len(target) {
		if _, ok := existed[target[i]]; !ok {
			return -1
		}
	}
	for right < len(target) {
		for right < len(target) && left < len(source) {
			if source[left] == target[right] {
				right++
			}
			left++
		}
		if left == len(source) {
			round++
			left = 0
		}
	}
	if left != 0 {
		round += 1
	}
	return round
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		source string
		target string
		want   int
	}{
		{"abc", "abcbc", 2},
		{"abc", "acdbc", -1},
		{"xyz", "xzyxz", 3},
		{"aaaaa", "aaaaaaaaaaaaa", 3},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, shortestWay(tt.source, tt.target))
	}
}

/*
// @lcpr case=start
// "abc"\n"abcbc"\n
// @lcpr case=end

// @lcpr case=start
// "abc"\n"acdbc"\n
// @lcpr case=end

// @lcpr case=start
// "xyz"\n"xzyxz"\n
// @lcpr case=end

*/
