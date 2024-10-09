package main

import (
	"testing"
)

/*
 * @lc app=leetcode.cn id=841 lang=golang
 * @lcpr version=20000
 *
 * [841] 钥匙和房间
 */

// @lc code=start
func canVisitAllRooms(rooms [][]int) bool {
	visited := make(map[int]struct{})
	cur := []int{0}
	for len(cur) > 0 {
		var next []int
		for _, room := range cur {
			if _, ok := visited[room]; ok {
				continue
			}
			visited[room] = struct{}{}
			next = append(next, rooms[room]...)
		}
		cur = next
	}
	return len(visited) == len(rooms)
}

// @lc code=end

/*
// @lcpr case=start
// [[1],[2],[3],[]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,3],[3,0,1],[2],[0]]\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		rooms [][]int
		want  bool
	}{
		{[][]int{{1}, {2}, {3}, {}}, true},
		{[][]int{{1, 3}, {3, 0, 1}, {2}, {0}}, false},
		{[][]int{{1, 2}, {2, 3}, {3}, {}}, true},
		{[][]int{{1}, {2}, {0, 3}, {1}}, true},
		{[][]int{{1, 2, 3}, {}, {}, {}}, true},
		{[][]int{{1}, {2}, {3}, {0}}, true},
		{[][]int{{1, 3}, {3, 0, 1}, {2}, {}}, false},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := canVisitAllRooms(tt.rooms); got != tt.want {
				t.Errorf("canVisitAllRooms() = %v, want %v", got, tt.want)
			}
		})
	}
}
