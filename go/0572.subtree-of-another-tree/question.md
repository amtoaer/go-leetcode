# [572. 另一棵树的子树][link] (Easy)

[link]: https://leetcode.cn/problems/subtree-of-another-tree/

给你两棵二叉树 `root` 和 `subRoot` 。检验 `root` 中是否包含和 `subRoot` 具有相同结构和节点值的子树。
如果存在，返回 `true` ；否则，返回 `false` 。

二叉树 `tree` 的一棵子树包括 `tree` 的某个节点和这个节点的所有后代节点。 `tree` 也可以看做它自身的一
棵子树。

**示例 1：**

![](https://pic.leetcode.cn/1724998676-cATjhe-image.png)

```
输入：root = [3,4,5,1,2], subRoot = [4,1,2]
输出：true
```

**示例 2：**

![](https://pic.leetcode.cn/1724998698-sEJWnq-image.png)

```
输入：root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]
输出：false
```

**提示：**

- `root` 树上的节点数量范围是 `[1, 2000]`
- `subRoot` 树上的节点数量范围是 `[1, 1000]`
- `-10⁴ <= root.val <= 10⁴`
- `-10⁴ <= subRoot.val <= 10⁴`
