# 从上到下打印二叉树
## 1. 题目描述
从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。


## 2. 示例
给定二叉树: [3,9,20,null,null,15,7],
```
    3
   / \
  9  20
    /  \
   15   7
```

返回：  
[3,9,20,15,7]

提示：   
节点总数 <= 1000

## 3. 解题
层序遍历
