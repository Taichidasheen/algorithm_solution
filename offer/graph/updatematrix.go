package graph


/*
剑指 Offer II 107. 矩阵中的距离
https://leetcode.cn/problems/2bCMpM/
解法：
bfs
需要注意的点：
bfs不要从矩阵中的每个元素开始，如果是这样的话，需要一个矩阵记录每个元素到0的最短距离。
bfs可以从矩阵中所有的零开始，这样只需要做一次bfs搜索即可
https://leetcode-cn.com/problems/01-matrix/solution/01ju-zhen-by-leetcode-solution/
 */
