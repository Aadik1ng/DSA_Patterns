# Data Structure and Algorithm Patterns

This repository contains implementations of common algorithm patterns used in coding interviews, implemented in multiple programming languages (C++, Java, Go, and Python).

## Table of Contents

1. [Sliding Window](#sliding-window)
2. [Two Pointers](#two-pointers)
3. [Fast & Slow Pointers](#fast--slow-pointers)
4. [Merge Intervals](#merge-intervals)
5. [Cyclic Sort](#cyclic-sort)
6. [In-place Reversal of a LinkedList](#in-place-reversal-of-a-linkedlist)
7. [Tree Breadth First Search](#tree-breadth-first-search)
8. [Tree Depth First Search](#tree-depth-first-search)
9. [Two Heaps](#two-heaps)
10. [Subsets](#subsets)
11. [Modified Binary Search](#modified-binary-search)
12. [Bitwise XOR](#bitwise-xor)
13. [Top 'K' Elements](#top-k-elements)
14. [K-way Merge](#k-way-merge)
15. [0/1 Knapsack](#01-knapsack)
16. [Topological Sort](#topological-sort)

## Pattern Descriptions

### Sliding Window

The Sliding Window pattern is used to perform a required operation on a specific window size of a given array or linked list, such as finding the longest subarray containing all 1s. Sliding Window starts from the 1st element and keeps shifting right by one element and adjusts the length of the window according to the problem that you are solving. It's useful for keeping track of a subset of data in an array/string.

**Time Complexity**: O(N) where N is the size of the input array.

### Two Pointers

The Two Pointers pattern uses two pointers to iterate through the data structure in tandem until one or both of the pointers hit a certain condition. Two Pointers is often useful when searching pairs in a sorted array or linked list; for example, when you need to find a target value from a sorted array.

**Time Complexity**: O(N) where N is the size of the input array.

### Fast & Slow Pointers

The Fast & Slow pointer approach, also known as the Hare & Tortoise algorithm, is a pointer algorithm that uses two pointers which move through the array (or sequence/LinkedList) at different speeds. This approach is quite useful when dealing with cyclic LinkedLists or arrays.

By moving at different speeds, the algorithm proves that the two pointers are bound to meet. The fast pointer should catch the slow pointer once both the pointers are in a cyclic loop.

**Common problems**: Finding cycles in a linked list, finding the middle of a linked list in one pass.

### Merge Intervals

The Merge Intervals pattern is an efficient technique to deal with overlapping intervals. In many interval-related problems, we either need to find overlapping intervals or merge intervals if they overlap.

Given two intervals (a and b), there will be six distinct ways the two intervals can relate to each other:

1. a and b do not overlap
2. a and b overlap, b ends after a
3. a completely overlaps b
4. a and b overlap, a ends after b
5. b completely overlaps a
6. a and b do not overlap

**Time Complexity**: O(N*logN) where N is the total number of intervals.

### Cyclic Sort

The Cyclic Sort pattern is used to deal with problems involving arrays containing numbers in a given range. This pattern places each number at its correct position in the array, and keeps swapping numbers until they are in their correct positions. It's especially useful when the input array contains numbers from 1 to n (with possible duplicates or missing numbers).

**Time Complexity**: O(N) where N is the size of the input array.

### In-place Reversal of a LinkedList

In many problems, we are asked to reverse the links between a set of nodes of a LinkedList. Often, the constraint is that we need to do this in-place, i.e., using the existing node objects and without using extra memory.

This pattern reverses one node at a time, starting with the current node and moving backward through the list, reconnecting the next pointers to point in the opposite direction.

**Time Complexity**: O(N) where N is the total number of nodes in the LinkedList.

### Tree Breadth First Search

The Tree BFS pattern is based on the Breadth First Search (BFS) technique to traverse a tree. It uses a queue to keep track of all the nodes of a level before jumping to the next level. This is also called level order traversal.

**Time Complexity**: O(N) where N is the total number of nodes in the tree.
**Space Complexity**: O(W) where W is the maximum width of the tree.

### Tree Depth First Search

The Tree DFS pattern is based on the Depth First Search (DFS) technique to traverse a tree. We use recursion (or an explicit stack for an iterative approach) to keep track of all previous (parent) nodes while traversing. This pattern is useful for problems involving tree paths or tree structure analysis.

**Time Complexity**: O(N) where N is the total number of nodes in the tree.
**Space Complexity**: O(H) where H is the height of the tree, for recursion stack.

### Two Heaps

The Two Heaps pattern uses two heaps to efficiently solve problems where you need to find the median of a set of numbers, or need to find the smallest element in one part and the biggest element in another part. It uses a Min Heap to find the smallest element and a Max Heap to find the biggest element.

**Time Complexity**: O(log N) for heap operations where N is the number of elements in the heaps.

### Subsets

The Subsets pattern deals with problems that involve finding permutations and combinations of a given set of elements. It uses an efficient Breadth First Search (BFS) approach to generate all possible subsets, combinations, or permutations.

**Time Complexity**: O(2^N) for generating all subsets where N is the number of elements.

### Modified Binary Search

The Modified Binary Search pattern is a variation of the traditional binary search. It's used when dealing with sorted arrays or lists, but with some modifications such as finding an element in a rotated sorted array, or finding a peak in a bitonic array.

**Time Complexity**: O(log N) where N is the size of the array.

### Bitwise XOR

The Bitwise XOR pattern uses XOR operations to solve various mathematical and bit manipulation problems efficiently. XOR is particularly useful for solving problems involving finding missing numbers, finding duplicate numbers, detecting a single different element, etc.

**Time Complexity**: O(N) where N is the size of the input.

### Top 'K' Elements

The Top 'K' Elements pattern is used to find the top (or bottom) K elements from a dataset. The best data structure to use for finding Top 'K' Elements is a Heap. This pattern is useful for solving problems like "Find the K largest elements in an array" or "Find K closest points to the origin."

**Time Complexity**: O(N log K) for building a heap where N is the total input elements.

### K-way Merge

The K-way Merge pattern helps solve problems involving a set of sorted arrays. Whenever we're given K sorted arrays, we can use a heap to efficiently perform a sorted traversal of all elements. We push the smallest element of each array into a Min Heap and then extract the minimum element from the heap and add the next element from the same array.

**Time Complexity**: O(N log K) where K is the number of arrays and N is the total number of elements in all arrays.

### 0/1 Knapsack

The 0/1 Knapsack pattern is a dynamic programming pattern based on the knapsack problem. It's useful when you need to optimize (maximize or minimize) a certain parameter, given some constraints. The pattern is about filling a knapsack of fixed capacity with items of different values and weights to maximize the value.

**Time Complexity**: O(N*C) where N is the number of items and C is the knapsack capacity.

### Topological Sort

The Topological Sort pattern is used to find a linear ordering of elements that have dependencies on each other. It's particularly useful for scheduling problems where some tasks need to be completed before others.

**Time Complexity**: O(V+E) where V is the number of vertices and E is the number of edges in the graph.

## Repository Structure

This repository contains implementations of all the above patterns in four programming languages:

- C++
- Java
- Go
- Python

Each implementation includes:
- Core algorithm implementation
- Examples of common problems solved using the pattern
- Test cases to verify correctness

## Usage

To run any specific pattern implementation, navigate to the respective language directory and run the appropriate file.

For example, to run the Sliding Window pattern in Python:

```bash
python Python/sliding_window.py
```

## Contributions

Contributions are welcome! Please feel free to submit a Pull Request. 