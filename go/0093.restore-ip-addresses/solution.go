// Created by amtoaer at 2025/11/21 17:32
// leetgo: 1.4.15
// https://leetcode.cn/problems/restore-ip-addresses/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func restoreIpAddresses(s string) (ans []string) {
	bytes := make([][]byte, 0, 4)
	sByte := []byte(s)
	var sb strings.Builder
	var backtrack func(idx int)
	backtrack = func(idx int) {
		if len(bytes) == 4 {
			if idx == len(s) {
				for idx, bytesItem := range bytes {
					sb.Write(bytesItem)
					if idx != 3 {
						sb.WriteByte('.')
					}
				}
				ans = append(ans, sb.String())
				sb.Reset()
			}
			return
		}
		if idx >= len(sByte) {
			return
		}
		if sByte[idx] == '0' {
			// 0 只能直接使用，后面不能跟其它数字
			bytes = append(bytes, sByte[idx:idx+1])
			backtrack(idx + 1)
			bytes = bytes[:len(bytes)-1]
		} else {
			var num int
			for cur := idx; cur < len(sByte); cur++ {
				num = num*10 + int(sByte[cur]-'0')
				if num >= 0 && num <= 255 {
					bytes = append(bytes, sByte[idx:cur+1])
					backtrack(cur + 1)
					bytes = bytes[:len(bytes)-1]
				}
			}
		}
	}
	backtrack(0)
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := restoreIpAddresses(s)

	fmt.Println("\noutput:", Serialize(ans))
}
