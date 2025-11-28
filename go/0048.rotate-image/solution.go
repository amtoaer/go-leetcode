// Created by amtoaer at 2025/11/28 11:27
// leetgo: 1.4.15
// https://leetcode.cn/problems/rotate-image/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func rotate(matrix [][]int) {
	for i := range matrix {
		for j := range i {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for i := range matrix {
		for j, k := 0, len(matrix[i])-1; j < k; j, k = j+1, k-1 {
			matrix[i][j], matrix[i][k] = matrix[i][k], matrix[i][j]
		}
	}
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	matrix := Deserialize[[][]int](ReadLine(stdin))
	rotate(matrix)
	ans := matrix

	fmt.Println("\noutput:", Serialize(ans))
}
