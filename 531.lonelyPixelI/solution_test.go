/*
 * @lc app=leetcode.cn id=531 lang=golang
 * @lcpr version=30202
 *
 * [531] 孤独像素 I
 */

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// @lc code=start
func findLonelyPixel(picture [][]byte) int {
	rowCnt := make([]int, len(picture))
	colCnt := make([]int, len(picture[0]))
	for i, row := range picture {
		for j, val := range row {
			if val == 'B' {
				rowCnt[i]++
				colCnt[j]++
			}
		}
	}
	var res int
	for i, rCnt := range rowCnt {
		for j, cCnt := range colCnt {
			if rCnt == 1 && cCnt == 1 && picture[i][j] == 'B' {
				res++
			}
		}
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		picture [][]byte
		want    int
	}{
		{[][]byte{{'W', 'W', 'B'}, {'W', 'B', 'W'}, {'B', 'W', 'W'}}, 3},
		{[][]byte{{'B', 'B', 'B'}, {'B', 'B', 'W'}, {'B', 'B', 'B'}}, 0},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, findLonelyPixel(tt.picture))
	}
}

/*
// @lcpr case=start
// [["W","W","B"],["W","B","W"],["B","W","W"]]\n
// @lcpr case=end

// @lcpr case=start
// [["B","B","B"],["B","B","W"],["B","B","B"]]\n
// @lcpr case=end

*/
