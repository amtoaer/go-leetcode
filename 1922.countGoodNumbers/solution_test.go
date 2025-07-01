/*
 * @lc app=leetcode.cn id=1922 lang=golang
 * @lcpr version=30201
 *
 * [1922] 统计好数字的数目
 */

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// @lc code=start
func countGoodNumbers(n int64) int {
	mod := int64(1e9 + 7)
	quickPow := func(a, b int64) int64 {
		var res int64 = 1
		for b > 0 {
			if b&1 == 1 {
				res = (res * a) % mod
			}
			a = a * a % mod
			b >>= 1
		}
		return res
	}
	return int(quickPow(5, (n+1)/2) * quickPow(4, n/2) % mod)
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		n    int64
		want int
	}{
		{1, 5},
		{4, 400},
		{50, 564908303},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, countGoodNumbers(tt.n))
	}
}

/*
// @lcpr case=start
// 1\n
// @lcpr case=end

// @lcpr case=start
// 4\n
// @lcpr case=end

// @lcpr case=start
// 50\n
// @lcpr case=end

*/
