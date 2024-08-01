package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=68 lang=golang
 *
 * [68] Text Justification
 */

// @lc code=start
func fullJustify(words []string, maxWidth int) []string {
	var sb strings.Builder
	var res []string
	for len(words) != 0 {
		chosenIdx, characterCount := 0, 0
		for chosenIdx < len(words) {
			afterCharacterCount := characterCount + len(words[chosenIdx])
			if afterCharacterCount+chosenIdx > maxWidth {
				break
			}
			characterCount = afterCharacterCount
			chosenIdx++
		}
		// 此时 chosenIdx 是下一个不属于当前行的词的索引，同时也是当前行的词数
		// spaceBlockCount 比词数少 1，可以代表词之间的空格块数，同时也是最后一个词的索引
		spaceBlockCount := chosenIdx - 1
		// 需要的空格个数是指定的宽度减去所有词的长度
		totalSpaceCount := maxWidth - characterCount
		for i, w := range words[:chosenIdx] {
			sb.WriteString(w)
			var spaceCount int
			if chosenIdx == 1 {
				// 如果这行只有一个词，则剩余的空格全部填充
				spaceCount = totalSpaceCount
			} else if len(words) <= chosenIdx { // 如果是最后一行（即没有剩下的词）
				if i == spaceBlockCount {
					// 把所有多余的空格都堆到最后一个词的后面
					spaceCount = totalSpaceCount - spaceBlockCount
				} else {
					// 前面的词之间只需要一个空格
					spaceCount = 1
				}
			} else { // 不是最后一行且不止一个词的普通情况
				if i == spaceBlockCount {
					// 最后一个词后面不需要空格
					spaceCount = 0
				} else {
					// 前面的词需要平均分配空格，如果有多余的空格，从左往右依次分配
					spaceCount = totalSpaceCount / spaceBlockCount
					if i < totalSpaceCount%spaceBlockCount {
						spaceCount++
					}
				}
			}
			for j := 0; j < spaceCount; j++ {
				sb.WriteByte(' ')
			}
		}
		res = append(res, sb.String())
		sb.Reset()
		words = words[chosenIdx:]
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input    []string
		maxWidth int
		output   []string
	}{
		{
			[]string{"This", "is", "an", "example", "of", "text", "justification."},
			16,
			[]string{"This    is    an", "example  of text", "justification.  "},
		},
		{
			[]string{"What", "must", "be", "acknowledgment", "shall", "be"},
			16,
			[]string{"What   must   be", "acknowledgment  ", "shall be        "},
		},
		{
			[]string{
				"Science",
				"is",
				"what",
				"we",
				"understand",
				"well",
				"enough",
				"to",
				"explain",
				"to",
				"a",
				"computer.",
				"Art",
				"is",
				"everything",
				"else",
				"we",
				"do",
			},
			20,
			[]string{
				"Science  is  what we",
				"understand      well",
				"enough to explain to",
				"a  computer.  Art is",
				"everything  else  we",
				"do                  ",
			},
		},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, fullJustify(tt.input, tt.maxWidth))
	}
}
