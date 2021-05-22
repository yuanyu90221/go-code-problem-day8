# go-code-problem-day8

## 題目描述

給定 Binary Search Tree (BST) 的根節點 root, 找出在此 BST 中第 k 小的數.

時間複雜度要求為 O(n)
Constraints:

此 BST 至少存在一個 node
1 <= k <= number of total nodes in the BST
Example:

Input: k = 4, binary tree = 

    Level 0         5 (root)
                   / \
    Level 1       3   9
                 /   / \
    Level 2    -1   6   10
                     \
    Level 3           7

Output: 6
Explanation: 此 BST 的 node 由小至大排序為 [-1, 3, 5, 6, 7, 9, 10]，所以第 4 小的 node 為 6

## 題目分析

已知 BST tree 特性

在節點的左子樹 節點的值比節點大
再節點的右子樹 節點的值比節點小

所以只要從 root節點順著 作 BFT就可逐步把每個數值大小做排序排好 因此 此過程為 O(n)
然後再來找出第 k-1個值即可（0-base）

