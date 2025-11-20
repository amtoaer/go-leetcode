# [143. 重排链表][link] (Medium)

[link]: https://leetcode.cn/problems/reorder-list/

给定一个单链表 `L` 的头节点 `head` ，单链表 `L` 表示为：

```
L₀ → L₁ → … → Lₙ ₋ ₁ → Lₙ
```

请将其重新排列后变为：

```
L₀ → Lₙ → L₁ → Lₙ ₋ ₁ → L₂ → Lₙ ₋ ₂ → …
```

不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

**示例 1：**

![](https://pic.leetcode-cn.com/1626420311-PkUiGI-image.png)

```
输入：head = [1,2,3,4]
输出：[1,4,2,3]
```

**示例 2：**

![](https://pic.leetcode-cn.com/1626420320-YUiulT-image.png)

```
输入：head = [1,2,3,4,5]
输出：[1,5,2,4,3]
```

**提示：**

- 链表的长度范围为 `[1, 5 * 10⁴]`
- `1 <= node.val <= 1000`
