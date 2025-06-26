/*
 * @lc app=leetcode.cn id=2311 lang=golang
 * @lcpr version=30201
 *
 * [2311] 小于等于 K 的最长二进制子序列
 */

package main

import (
	"math/bits"
	"testing"

	"github.com/stretchr/testify/assert"
)

// @lc code=start
func longestSubsequence(s string, k int) int {
	var cnt, sum int
	l := bits.Len(uint(k))
	for i := range s {
		pos := len(s) - 1 - i
		if s[pos] == '0' {
			cnt++
		} else {
			// i <= l 可以剪枝并防止溢出
			if i <= l && sum+1<<i <= k {
				sum += (1 << i)
				cnt++
			}
		}
	}
	return cnt
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		s    string
		k    int
		want int
	}{
		{"1001010", 5, 5},
		{"00101001", 1, 6},
		{"111100010000011101001110001111000000001011101111111110111000011111011000010101110100110110001111001001011001010011010000011111101001101000000101101001110110000111101011000101", 11713332, 96},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, longestSubsequence(tt.s, tt.k))
	}
}

/*
// @lcpr case=start
// "1001010"\n5\n
// @lcpr case=end

// @lcpr case=start
// "00101001"\n1\n
// @lcpr case=end

*/
