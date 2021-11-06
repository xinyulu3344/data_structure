# 堆

堆的一个重要性质：任意节点的值总是>=或<=子节点的值

 大于等于：最大堆、大根堆、大顶堆
 小于等于：最小堆、小根堆、小顶堆

堆中元素比较具备可比较性

## 常见的堆实现

二叉堆 （Binary Heap Heap，完全二叉堆 ）
多叉堆 （D-heap 、D-ary Heap Heap）
索引堆 （Index Heap Heap）
二项堆 （BinomialBinomialHeap ）
斐波那契堆 （Fibonacci Heap Heap）
左倾堆 （Leftist Heap Heap，左式堆 ）
斜堆 （Skew Heap Heap）

## 堆的复杂度

1. 获取最大值：O(1)
2. 删除最大值：O(logn)
3. 添加元素：O(logn)

## 堆的使用场景

### Top K问题

从海量数据中找出前K个数据

## 二叉堆

