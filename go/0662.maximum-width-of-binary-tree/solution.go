// Created by amtoaer at 2025/11/29 02:02
// leetgo: 1.4.15
// https://leetcode.cn/problems/maximum-width-of-binary-tree/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

type Pair struct {
	node *TreeNode
	idx  int
}

func widthOfBinaryTree(root *TreeNode) (ans int) {
	queue := []Pair{{root, 1}}
	for len(queue) > 0 {
		ans = max(ans, queue[len(queue)-1].idx-queue[0].idx+1)
		lenQueue := len(queue)
		for _, pair := range queue[:lenQueue] {
			if pair.node.Left != nil {
				queue = append(queue, Pair{pair.node.Left, pair.idx * 2})
			}
			if pair.node.Right != nil {
				queue = append(queue, Pair{pair.node.Right, pair.idx*2 + 1})
			}
		}
		queue = queue[lenQueue:]
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	root := Deserialize[*TreeNode](ReadLine(stdin))
	ans := widthOfBinaryTree(root)

	fmt.Println("\noutput:", Serialize(ans))
}
