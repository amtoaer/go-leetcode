/*
 * @lc app=leetcode.cn id=1056 lang=golang
 * @lcpr version=30202
 *
 * [1056] 易混淆数
 */

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// @lc code=start
func confusingNumber(n int) bool {
	var (
		old = n
		new int
		m   = map[int]int{
			0: 0,
			1: 1,
			6: 9,
			8: 8,
			9: 6,
		}
	)
	for n > 0 {
		num := n % 10
		if mapped, ok := m[num]; ok {
			new = new*10 + mapped
		} else {
			return false
		}
		n /= 10
	}
	return old != new
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		n    int
		want bool
	}{
		{6, true},
		{89, true},
		{11, false},
		{25, false},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, confusingNumber(tt.n))
	}
}

/*
// @lcpr case=start
// 6\n
// @lcpr case=end

// @lcpr case=start
// 89\n
// @lcpr case=end

// @lcpr case=start
// 11\n
// @lcpr case=end

// @lcpr case=start
// 25\n
// @lcpr case=end

*/
