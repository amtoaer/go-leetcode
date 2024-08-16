package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=433 lang=golang
 *
 * [433] Minimum Genetic Mutation
 */

// @lc code=start
func diffOne(a, b string) bool {
	var res bool
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			if res {
				return false
			}
			res = true
		}
	}
	return res
}

func minMutation(startGene string, endGene string, bank []string) int {
	if startGene == endGene {
		return 0
	}
	var (
		step    int
		visited = make([]bool, len(bank))
		queue   = []string{startGene}
	)
	for len(queue) > 0 {
		step++
		lenQueue := len(queue)
		for _, str := range queue[:lenQueue] {
			for idx, item := range bank {
				if visited[idx] {
					continue
				}
				if diffOne(str, item) {
					if item == endGene {
						return step
					}
					queue = append(queue, item)
					visited[idx] = true
				}
			}
		}
		queue = queue[lenQueue:]
	}
	return -1
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		startGene string
		endGene   string
		bank      []string
		want      int
	}{
		{"AACCGGTT", "AACCGGTA", []string{"AACCGGTA"}, 1},
		{"AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}, 2},
		{"AAAAACCC", "AACCCCCC", []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"}, 3},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, minMutation(tt.startGene, tt.endGene, tt.bank))
	}
}
