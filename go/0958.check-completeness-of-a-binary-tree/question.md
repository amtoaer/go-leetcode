# [958. 二叉树的完全性检验][link] (Medium)

[link]: https://leetcode.cn/problems/check-completeness-of-a-binary-tree/

给你一棵二叉树的根节点 `root` ，请你判断这棵树是否是一棵 **完全二叉树** 。

在一棵 **[完全二叉树](https://baike.baidu.com/item/完全二叉树/7773232?fr=aladdin)** 中，除了最后一层
外，所有层都被完全填满，并且最后一层中的所有节点都尽可能靠左。最后一层（第 `h` 层）中可以包含 `1` 到
`2ʰ` 个节点。

**示例 1：**

![](https://assets.leetcode.cn/aliyun-lc-upload/uploads/2018/12/15/complete-binary-tree-1.png)

```
输入：root = [1,2,3,4,5,6]
输出：true
解释：最后一层前的每一层都是满的（即，节点值为 {1} 和 {2,3} 的两层），且最后一层中的所有节点（{4,5,6
}）尽可能靠左。
```

**示例 2：**

**![](https://assets.leetcode.cn/aliyun-lc-upload/uploads/2018/12/15/complete-binary-tree-2.png)**

```
输入：root = [1,2,3,4,5,null,7]
输出：false
解释：值为 7 的节点不满足条件「节点尽可能靠左」。
```

**提示：**

- 树中节点数目在范围 `[1, 100]` 内
- `1 <= Node.val <= 1000`
