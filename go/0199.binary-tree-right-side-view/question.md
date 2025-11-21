# [199. 二叉树的右视图][link] (Medium)

[link]: https://leetcode.cn/problems/binary-tree-right-side-view/

给定一个二叉树的 **根节点** `root`，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到
的节点值。

**示例 1：**

**输入：** root = \[1,2,3,null,5,null,4\]

**输出：**\[1,3,4\]

**解释：**

![](https://assets.leetcode.com/uploads/2024/11/24/tmpd5jn43fs-1.png)

**示例 2：**

**输入：** root = \[1,2,3,4,null,null,null,5\]

**输出：**\[1,3,4,5\]

**解释：**

![](https://assets.leetcode.com/uploads/2024/11/24/tmpkpe40xeh-1.png)

**示例 3：**

**输入：** root = \[1,null,3\]

**输出：**\[1,3\]

**示例 4：**

**输入：** root = \[\]

**输出：**\[\]

**提示:**

- 二叉树的节点个数的范围是 `[0,100]`
- `-100 <= Node.val <= 100`
