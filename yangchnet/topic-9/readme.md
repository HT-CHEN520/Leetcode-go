# 回文数

## 1. 题目描述
 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。

回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，121 是回文，而 123 不是。

## 2. 示例

示例 1：
```
输入：x = 121
输出：true
```

示例 2：
```
输入：x = -121
输出：false
解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
```

示例 3：
```
输入：x = 10
输出：false
解释：从右向左读, 为 01 。因此它不是一个回文数。
```

示例 4：
```
输入：x = -101
输出：false
```

提示：
$$-2^{31} \le x \le 2^{31} - 1$$

## 3. 解题
特殊情况： 若为负数直接返回false， 0-9直接返回true
1. 将整数不断模10获取其每一位，再依次比较
2. 将整数转化为字符串再依次比较
3. 将数字的后半部分翻转，再比较数字的后半部分和前半部分
