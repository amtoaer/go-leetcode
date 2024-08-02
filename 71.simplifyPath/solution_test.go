package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=71 lang=golang
 *
 * [71] Simplify Path
 */

// @lc code=start
func simplifyPath(path string) string {
	var start int
	var paths []string
	var sb strings.Builder
	for i := 1; i < len(path); i++ {
		if path[i-1] == '/' && path[i] != '/' {
			start = i
		}
		if (path[i] == '/' && path[i-1] != '/') || (i == len(path)-1 && path[i] != '/') {
			if i == len(path)-1 && path[i] != '/' {
				i++
			}
			p := path[start:i]
			switch p {
			case ".":
				{
					continue
				}
			case "..":
				{
					if len(paths) > 0 {
						paths = paths[:len(paths)-1]
					}
				}
			default:
				paths = append(paths, p)
			}
		}
	}
	for _, p := range paths {
		sb.WriteByte('/')
		sb.WriteString(p)
	}
	res := sb.String()
	if len(res) == 0 {
		return "/"
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		path string
		ans  string
	}{
		{"/home/", "/home"},
		{"/a/./b/../../c/", "/c"},
		{"/a/../../b/../c//.//", "/c"},
		{"/a/../b/../c//.//", "/c"},
		{"/home/user/Documents/../Pictures", "/home/user/Pictures"},
		{"/../", "/"},
	}
	for _, tc := range tc {
		assert.Equal(t, tc.ans, simplifyPath(tc.path))
	}
}
