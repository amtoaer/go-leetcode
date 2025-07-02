/*
 * @lc app=leetcode.cn id=3304 lang=golang
 * @lcpr version=30201
 *
 * [3304] 找出第 K 个字符 I
 */

package main

import (
	"math/bits"
	"testing"
)

// @lc code=start
func kthCharacter(k int) byte {
	var offset int
	for k != 1 {
		t := bits.Len(uint(k)) - 1
		if k == 1<<t {
			t--
		}
		k -= 1 << t
		offset++
	}
	return byte('a' + offset%26)
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		k    int
		want byte
	}{
		{5, 'b'},
		{10, 'c'},
	}
	for _, tt := range tc {
		if got := kthCharacter(tt.k); got != tt.want {
			t.Errorf("kthCharacter(%v) = %c, want %c", tt.k, got, tt.want)
		}
	}
}

/*
// @lcpr case=start
// 5\n
// @lcpr case=end

// @lcpr case=start
// 10\n
// @lcpr case=end

*/
