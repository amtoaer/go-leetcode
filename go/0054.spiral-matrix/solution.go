// Created by amtoaer at 2025/11/20 14:22
// leetgo: 1.4.15
// https://leetcode.cn/problems/spiral-matrix/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func spiralOrder(matrix [][]int) (ans []int) {
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	for top <= bottom && left <= right {
		ans = append(ans, matrix[top][left:right+1]...)
		for i := top + 1; i <= bottom; i++ {
			ans = append(ans, matrix[i][right])
		}
		if top < bottom {
			for j := right - 1; j >= left; j-- {
				ans = append(ans, matrix[bottom][j])
			}
		}
		if left < right {
			for i := bottom - 1; i > top; i-- {
				ans = append(ans, matrix[i][left])
			}
		}
		top, bottom, left, right = top+1, bottom-1, left+1, right-1
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	matrix := Deserialize[[][]int](ReadLine(stdin))
	ans := spiralOrder(matrix)

	fmt.Println("\noutput:", Serialize(ans))
}
