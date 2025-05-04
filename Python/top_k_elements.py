"""
╔════════════════════════════════════════════════════════╗
║               Top K Elements Pattern                  ║
╠════════════════════════════════════════════════════════╣
║ Description:                                           ║
║ Use nth_element to partition so that the first K      ║
║ elements are the top K (unordered).                  ║
╠════════════════════════════════════════════════════════╣
║ Flow Diagram (ASCII):                                  ║
║ arr=[3,2,1,5,6,4], k=2                                 ║
║ nth_element to place 2nd largest at index 2          ║
║ Partitioned: [6,5 | 1,2,3,4]                          ║
║ First 2 = [6,5]                                        ║
╚════════════════════════════════════════════════════════╝
"""

import heapq

def main():
    arr = [3,2,1,5,6,4]
    k = 2
    # Use a min-heap of size k
    minH = []
    for x in arr:
        heapq.heappush(minH, x)
        if len(minH) > k:
            heapq.heappop(minH)
    result = sorted(minH, reverse=True)
    print(f"Top {k}:", *result)

if __name__ == "__main__":
    main() 