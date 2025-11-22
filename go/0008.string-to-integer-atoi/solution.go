// Created by amtoaer at 2025/11/22 12:40
// leetgo: 1.4.15
// https://leetcode.cn/problems/string-to-integer-atoi/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func myAtoi(s string) (ans int) {
	leadingZero, leadingSpace, leadingLabel := true, true, true
	flag := 1
	var num int
	for _, r := range s {
		if r == ' ' {
			if leadingSpace {
				continue
			}
			break
		}
		leadingSpace = false
		if r == '-' || r == '+' {
			if leadingLabel {
				if r == '-' {
					flag = -1
				}
				leadingLabel = false
				continue
			}
			break
		}
		if r < '0' || r > '9' {
			break
		}
		leadingLabel = false
		if r == '0' && leadingZero {
			continue
		}
		leadingZero = false
		runeNum := r - '0'
		num = num*10 + int(runeNum)
		if flag == 1 && num > math.MaxInt32 {
			return math.MaxInt32
		}
		if flag == -1 && num > -math.MinInt32 {
			return math.MinInt32
		}
	}
	return num * flag
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := myAtoi(s)

	fmt.Println("\noutput:", Serialize(ans))
}
