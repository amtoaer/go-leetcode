这是一个用于存储 leetcode 题解的 Golang 代码仓库，包含 testify 测试框架。

当生成单元测试时，请采用结构体列表存储测试样例的输入与期望结果。

严格使用你对该题目的已知测试样例，不要自己生成额外的新样例。

在断言结果一致时，请尽可能采用 testify 框架中提供的工具包而非自己编写对应实现。

如果上下文中已经存在 func Test(t \*testing.T) {}，请直接生成函数体，不要额外重复包含方法签名。

以下是一个例子：

```go
func Test(t *testing.T) {
	tc := []struct {
		nums1 []int
		nums2 []int
		want  int
	}{
		{[]int{1, 4, 2}, []int{1, 2, 4}, 2},
		{[]int{2, 5, 1, 2, 5}, []int{10, 5, 2, 1, 5, 2}, 3},
		{[]int{1, 3, 7, 1, 7, 5}, []int{1, 9, 2, 5, 1}, 2},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, maxUncrossedLines(tt.nums1, tt.nums2))
	}
}
```
