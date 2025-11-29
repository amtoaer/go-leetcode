// Created by amtoaer at 2025/11/29 21:58
// leetgo: 1.4.15
// https://leetcode.cn/problems/basic-calculator-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func calculate(s string) (ans int) {
	var nums []int
	var num int
	sign := byte('+')
	for idx := range s {
		if isNumber(s[idx]) {
			num = num*10 + int(s[idx]-'0')
		}
		if (s[idx] != ' ' && !isNumber(s[idx])) || idx == len(s)-1 {
			switch sign {
			case '+':
				{
					nums = append(nums, num)
				}
			case '-':
				{
					nums = append(nums, -num)
				}
			case '*':
				{
					nums[len(nums)-1] *= num
				}
			case '/':
				{
					nums[len(nums)-1] /= num
				}
			}
			sign = s[idx]
			num = 0
		}
	}
	for _, num := range nums {
		ans += num
	}
	return ans
}

func isNumber(b byte) bool {
	return b >= '0' && b <= '9'
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := calculate(s)

	fmt.Println("\noutput:", Serialize(ans))
}
