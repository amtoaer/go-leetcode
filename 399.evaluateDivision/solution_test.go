package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=399 lang=golang
 *
 * [399] Evaluate Division
 */

// @lc code=start
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	id := make(map[string]int)
	for _, equation := range equations {
		left, right := equation[0], equation[1]
		if _, ok := id[left]; !ok {
			id[left] = len(id)
		}
		if _, ok := id[right]; !ok {
			id[right] = len(id)
		}
	}
	graph := make([][]float64, len(id))
	for i := range graph {
		graph[i] = make([]float64, len(id))
	}
	for idx, equation := range equations {
		left, right := id[equation[0]], id[equation[1]]
		graph[left][right] = values[idx]
		graph[right][left] = 1 / values[idx]
	}
	for j := range graph {
		for i := range graph {
			for k := range graph {
				if graph[i][j] > 1e-6 && graph[j][k] > 1e-6 {
					graph[i][k] = graph[i][j] * graph[j][k]
				}
			}
		}
	}
	res := make([]float64, 0, len(queries))
	for _, query := range queries {
		left, okLeft := id[query[0]]
		right, okRight := id[query[1]]
		if !okLeft || !okRight || graph[left][right] == 0 {
			res = append(res, -1.0)
			continue
		}
		res = append(res, graph[left][right])
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		equations [][]string
		values    []float64
		queries   [][]string
		expected  []float64
	}{
		{
			[][]string{{"a", "c"}, {"b", "e"}, {"c", "d"}, {"e", "d"}},
			[]float64{2.0, 3.0, 0.5, 5.0},
			[][]string{{"a", "b"}},
			[]float64{0.06667},
		},
	}
	for _, tt := range tc {
		assert.InDeltaSlice(t, tt.expected, calcEquation(tt.equations, tt.values, tt.queries), 1e-5)
	}
}
