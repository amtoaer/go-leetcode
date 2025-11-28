# [144. 二叉树的前序遍历][link] (Easy)

[link]: https://leetcode.cn/problems/binary-tree-preorder-traversal/

给你二叉树的根节点 `root` ，返回它节点值的 **前序** 遍历。

**示例 1：**

**输入：** root = \[1,null,2,3\]

**输出：**\[1,2,3\]

**解释：**

![](https://assets.leetcode.com/uploads/2024/08/29/screenshot-2024-08-29-202743.png)

**示例 2：**

**输入：** root = \[1,2,3,4,5,null,8,null,null,6,7,9\]

**输出：**\[1,2,4,5,6,7,3,8,9\]

**解释：**

![](https://assets.leetcode.com/uploads/2024/08/29/tree_2.png)

**示例 3：**

**输入：** root = \[\]

**输出：**\[\]

**示例 4：**

**输入：** root = \[1\]

**输出：**\[1\]

**提示：**

- 树中节点数目在范围 `[0, 100]` 内
- `-100 <= Node.val <= 100`

**进阶：** 递归算法很简单，你可以通过迭代算法完成吗？
