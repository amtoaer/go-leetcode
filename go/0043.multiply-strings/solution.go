// Created by amtoaer at 2025/11/22 13:57
// leetgo: 1.4.15
// https://leetcode.cn/problems/multiply-strings/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func multiply(num1 string, num2 string) (ans string) {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	for i := len(num2) - 1; i >= 0; i-- {
		res := multiplySingle(num1, num2[i], len(num2)-i-1)
		if len(ans) > 0 {
			ans = sum(res, ans)
		} else {
			ans = res
		}
	}
	return
}

func multiplySingle(num1 string, num2 byte, offset int) string {
	carry := 0
	var bytes []byte
	for range offset {
		bytes = append(bytes, '0')
	}
	num2Num := int(num2 - '0')
	for idx := len(num1) - 1; idx >= 0 || carry != 0; idx-- {
		if idx >= 0 {
			carry = num2Num*int(num1[idx]-'0') + carry
		}
		bytes = append(bytes, byte(carry%10+'0'))
		carry /= 10
	}
	reverse(bytes)
	return string(bytes)
}

func sum(num1, num2 string) string {
	carry := 0
	var bytes []byte
	i, j := len(num1)-1, len(num2)-1
	for i >= 0 || j >= 0 || carry > 0 {
		if i >= 0 {
			carry += int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			carry += int(num2[j] - '0')
			j--
		}
		bytes = append(bytes, byte(carry%10+'0'))
		carry /= 10
	}
	reverse(bytes)
	return string(bytes)
}

func reverse(bytes []byte) {
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	num1 := Deserialize[string](ReadLine(stdin))
	num2 := Deserialize[string](ReadLine(stdin))
	ans := multiply(num1, num2)

	fmt.Println("\noutput:", Serialize(ans))
}
