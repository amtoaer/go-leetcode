package main

import (
	"fmt"
	"testing"
)

/*
 * @lc app=leetcode.cn id=3222 lang=golang
 * @lcpr version=20002
 *
 * [3222] 求出硬币游戏的赢家
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func losingPlayer(x int, y int) string {
	times := min(x, y/4)
	if times&1 == 1 {
		return "Alice"
	}
	return "Bob"
}

// @lc code=end

/*
// @lcpr case=start
// 2\n7\n
// @lcpr case=end

// @lcpr case=start
// 4\n11\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		x, y int
		want string
	}{
		{2, 7, "Alice"},
		{4, 11, "Bob"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d,%d", tt.x, tt.y), func(t *testing.T) {
			if got := losingPlayer(tt.x, tt.y); got != tt.want {
				t.Errorf("losingPlayer(%d, %d) = %v; want %v", tt.x, tt.y, got, tt.want)
			}
		})
	}
}
