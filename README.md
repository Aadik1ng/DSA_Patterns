# 🎯 Data Structure and Algorithm Patterns

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Languages](https://img.shields.io/badge/Languages-4-blue.svg)](https://github.com/yourusername/dsa-patterns)
[![Patterns](https://img.shields.io/badge/Patterns-16-green.svg)](https://github.com/yourusername/dsa-patterns)

> A comprehensive collection of algorithm patterns implemented in multiple languages to help you ace coding interviews! 🚀

## 📚 Table of Contents

| # | Pattern | Description |
|---|---------|-------------|
| 1 | [Sliding Window](#sliding-window) | Efficient subarray operations |
| 2 | [Two Pointers](#two-pointers) | Array manipulation with dual pointers |
| 3 | [Fast & Slow Pointers](#fast--slow-pointers) | Cycle detection in linked lists |
| 4 | [Merge Intervals](#merge-intervals) | Handling overlapping intervals |
| 5 | [Cyclic Sort](#cyclic-sort) | Sorting numbers in a given range |
| 6 | [In-place Reversal](#in-place-reversal-of-a-linkedlist) | Reversing linked lists in-place |
| 7 | [Tree BFS](#tree-breadth-first-search) | Level-order tree traversal |
| 8 | [Tree DFS](#tree-depth-first-search) | Depth-first tree traversal |
| 9 | [Two Heaps](#two-heaps) | Median finding and more |
| 10 | [Subsets](#subsets) | Generating all possible subsets |
| 11 | [Modified Binary Search](#modified-binary-search) | Advanced binary search variations |
| 12 | [Bitwise XOR](#bitwise-xor) | Efficient bit manipulation |
| 13 | [Top 'K' Elements](#top-k-elements) | Finding top K elements |
| 14 | [K-way Merge](#k-way-merge) | Merging multiple sorted arrays |
| 15 | [0/1 Knapsack](#01-knapsack) | Dynamic programming optimization |
| 16 | [Topological Sort](#topological-sort) | Dependency resolution |

## 🧩 Pattern Descriptions

### 🔄 Sliding Window
**Time Complexity**: O(N) | **Space Complexity**: O(1)

The Sliding Window pattern is used to perform operations on a specific window size of a given array or linked list. It's particularly useful for:
- Finding the longest subarray containing all 1s
- Calculating averages of subarrays
- Finding maximum/minimum in a window

```python
# Example: Maximum sum of subarray of size K
def max_subarray_sum(arr, k):
    window_sum = sum(arr[:k])
    max_sum = window_sum
    
    for i in range(len(arr) - k):
        window_sum = window_sum - arr[i] + arr[i + k]
        max_sum = max(max_sum, window_sum)
    
    return max_sum
```

### 🎯 Two Pointers
**Time Complexity**: O(N) | **Space Complexity**: O(1)

The Two Pointers pattern uses two pointers to iterate through data structures in tandem. Common applications:
- Finding pairs in a sorted array
- Removing duplicates
- Finding subarrays with given sum

```python
# Example: Finding pairs with target sum
def find_pair(arr, target):
    left, right = 0, len(arr) - 1
    while left < right:
        current_sum = arr[left] + arr[right]
        if current_sum == target:
            return [left, right]
        elif current_sum < target:
            left += 1
        else:
            right -= 1
    return [-1, -1]
```

### 🐇 Fast & Slow Pointers
**Time Complexity**: O(N) | **Space Complexity**: O(1)

The Fast & Slow pointer approach is essential for:
- Cycle detection in linked lists
- Finding the middle of a linked list
- Finding the start of a cycle

```python
# Example: Cycle detection
def has_cycle(head):
    slow = fast = head
    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next
        if slow == fast:
            return True
    return False
```

### 📊 Merge Intervals
**Time Complexity**: O(N*logN) | **Space Complexity**: O(N)

The Merge Intervals pattern handles overlapping intervals efficiently:
- Merging overlapping intervals
- Finding conflicting appointments
- Scheduling problems

```python
# Example: Merging overlapping intervals
def merge_intervals(intervals):
    intervals.sort(key=lambda x: x[0])
    merged = []
    for interval in intervals:
        if not merged or merged[-1][1] < interval[0]:
            merged.append(interval)
        else:
            merged[-1][1] = max(merged[-1][1], interval[1])
    return merged
```

### 🔄 Cyclic Sort
**Time Complexity**: O(N) | **Space Complexity**: O(1)

The Cyclic Sort pattern is perfect for:
- Sorting numbers in a given range
- Finding missing numbers
- Finding duplicate numbers

```python
# Example: Cyclic sort
def cyclic_sort(nums):
    i = 0
    while i < len(nums):
        j = nums[i] - 1
        if nums[i] != nums[j]:
            nums[i], nums[j] = nums[j], nums[i]
        else:
            i += 1
    return nums
```

### 🔗 In-place Reversal of a LinkedList
**Time Complexity**: O(N) | **Space Complexity**: O(1)

This pattern is essential for:
- Reversing linked lists
- Reversing sublists
- Palindrome checking

```python
# Example: Reversing a linked list
def reverse_linked_list(head):
    prev = None
    current = head
    while current:
        next_node = current.next
        current.next = prev
        prev = current
        current = next_node
    return prev
```

### 🌳 Tree Breadth First Search
**Time Complexity**: O(N) | **Space Complexity**: O(W)

The Tree BFS pattern is used for:
- Level order traversal
- Finding minimum depth
- Finding right view

```python
# Example: Level order traversal
def level_order_traversal(root):
    if not root:
        return []
    result = []
    queue = [root]
    while queue:
        level_size = len(queue)
        current_level = []
        for _ in range(level_size):
            node = queue.pop(0)
            current_level.append(node.val)
            if node.left:
                queue.append(node.left)
            if node.right:
                queue.append(node.right)
        result.append(current_level)
    return result
```

### 🌲 Tree Depth First Search
**Time Complexity**: O(N) | **Space Complexity**: O(H)

The Tree DFS pattern is useful for:
- Path sum problems
- Tree serialization
- Finding diameter

```python
# Example: Path sum
def has_path_sum(root, target_sum):
    if not root:
        return False
    if not root.left and not root.right:
        return target_sum == root.val
    return (has_path_sum(root.left, target_sum - root.val) or
            has_path_sum(root.right, target_sum - root.val))
```

### ⚖️ Two Heaps
**Time Complexity**: O(logN) | **Space Complexity**: O(N)

The Two Heaps pattern is perfect for:
- Finding median of a stream
- Sliding window median
- Interval problems

```python
# Example: Median of a stream
class MedianFinder:
    def __init__(self):
        self.max_heap = []  # smaller half
        self.min_heap = []  # larger half

    def add_num(self, num):
        heapq.heappush(self.max_heap, -num)
        heapq.heappush(self.min_heap, -heapq.heappop(self.max_heap))
        if len(self.min_heap) > len(self.max_heap):
            heapq.heappush(self.max_heap, -heapq.heappop(self.min_heap))
```

### 🔢 Subsets
**Time Complexity**: O(2^N) | **Space Complexity**: O(2^N)

The Subsets pattern is used for:
- Generating all subsets
- Permutations
- Combinations

```python
# Example: Generating subsets
def generate_subsets(nums):
    subsets = [[]]
    for num in nums:
        n = len(subsets)
        for i in range(n):
            subsets.append(subsets[i] + [num])
    return subsets
```

### 🔍 Modified Binary Search
**Time Complexity**: O(logN) | **Space Complexity**: O(1)

The Modified Binary Search pattern is used for:
- Finding elements in rotated arrays
- Finding peak elements
- Bitonic arrays

```python
# Example: Finding element in rotated array
def search_rotated_array(nums, target):
    left, right = 0, len(nums) - 1
    while left <= right:
        mid = (left + right) // 2
        if nums[mid] == target:
            return mid
        if nums[left] <= nums[mid]:
            if nums[left] <= target < nums[mid]:
                right = mid - 1
            else:
                left = mid + 1
        else:
            if nums[mid] < target <= nums[right]:
                left = mid + 1
            else:
                right = mid - 1
    return -1
```

### 🧮 Bitwise XOR
**Time Complexity**: O(N) | **Space Complexity**: O(1)

The Bitwise XOR pattern is used for:
- Finding missing numbers
- Finding single numbers
- Bit manipulation

```python
# Example: Finding single number
def single_number(nums):
    result = 0
    for num in nums:
        result ^= num
    return result
```

### 🔝 Top 'K' Elements
**Time Complexity**: O(N logK) | **Space Complexity**: O(K)

The Top 'K' Elements pattern is used for:
- Finding K largest elements
- Finding K closest points
- Frequency problems

```python
# Example: Finding K largest elements
def find_k_largest(nums, k):
    return heapq.nlargest(k, nums)
```

### 🔄 K-way Merge
**Time Complexity**: O(N logK) | **Space Complexity**: O(K)

The K-way Merge pattern is used for:
- Merging K sorted lists
- Finding smallest range
- External sorting

```python
# Example: Merging K sorted lists
def merge_k_sorted_lists(lists):
    min_heap = []
    for i, lst in enumerate(lists):
        if lst:
            heapq.heappush(min_heap, (lst[0], i, 0))
    
    result = []
    while min_heap:
        val, list_idx, element_idx = heapq.heappop(min_heap)
        result.append(val)
        if element_idx + 1 < len(lists[list_idx]):
            heapq.heappush(min_heap, (lists[list_idx][element_idx + 1], list_idx, element_idx + 1))
    return result
```

### 🎒 0/1 Knapsack
**Time Complexity**: O(N*C) | **Space Complexity**: O(C)

The 0/1 Knapsack pattern is used for:
- Resource allocation
- Optimization problems
- Dynamic programming

```python
# Example: 0/1 Knapsack
def knapsack(weights, values, capacity):
    n = len(weights)
    dp = [0] * (capacity + 1)
    for i in range(n):
        for w in range(capacity, weights[i] - 1, -1):
            dp[w] = max(dp[w], dp[w - weights[i]] + values[i])
    return dp[capacity]
```

### 📊 Topological Sort
**Time Complexity**: O(V+E) | **Space Complexity**: O(V)

The Topological Sort pattern is used for:
- Task scheduling
- Course prerequisites
- Dependency resolution

```python
# Example: Topological sort
def topological_sort(vertices, edges):
    graph = {i: [] for i in range(vertices)}
    in_degree = {i: 0 for i in range(vertices)}
    
    for parent, child in edges:
        graph[parent].append(child)
        in_degree[child] += 1
    
    sources = deque([i for i in range(vertices) if in_degree[i] == 0])
    sorted_order = []
    
    while sources:
        vertex = sources.popleft()
        sorted_order.append(vertex)
        for child in graph[vertex]:
            in_degree[child] -= 1
            if in_degree[child] == 0:
                sources.append(child)
    
    return sorted_order if len(sorted_order) == vertices else []
```

## 🏗️ Repository Structure

```
.
├── C++/
│   ├── sliding_window.cpp
│   ├── two_pointers.cpp
│   └── ...
├── Java/
│   ├── SlidingWindow.java
│   ├── TwoPointers.java
│   └── ...
├── Go/
│   ├── sliding_window.go
│   ├── two_pointers.go
│   └── ...
└── Python/
    ├── sliding_window.py
    ├── two_pointers.py
    └── ...
```

## 🚀 Usage

Each pattern is implemented in multiple languages. To run a specific implementation:

```bash
# Python
python Python/sliding_window.py

# Java
javac Java/SlidingWindow.java
java SlidingWindow

# C++
g++ C++/sliding_window.cpp -o sliding_window
./sliding_window

# Go
go run Go/sliding_window.go
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Inspired by various coding interview preparation resources
- Special thanks to the open-source community 