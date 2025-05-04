"""
╔════════════════════════════════════════════════════════╗
║              Two Heaps Pattern (Median)               ║
╠════════════════════════════════════════════════════════╣
║ Description:                                           ║
║ Maintain a max‑heap for the lower half and a min‑heap ║
║ for the upper half of the stream to get median in O(log N) per insert. ║
╠════════════════════════════════════════════════════════╣
║ Flow Diagram (ASCII):                                  ║
║ Stream: 5,15,1,3                                      ║
║                                                      ║
║ Insert 5 → maxH:[5], minH:[] → median=5               ║
║ Insert 15→ maxH:[5], minH:[15] → median=(5+15)/2=10   ║
║ Insert 1 → maxH:[5,1], minH:[15] → median=5           ║
║ Insert 3 → maxH:[3,1], minH:[5,15] → median=(3+5)/2=4 ║
╚════════════════════════════════════════════════════════╝
"""

import heapq

def main():
    maxH = []  # max-heap (invert sign)
    minH = []  # min-heap
    nums = [5, 15, 1, 3]
    for x in nums:
        if not maxH or x < -maxH[0]:
            heapq.heappush(maxH, -x)
        else:
            heapq.heappush(minH, x)
        # Rebalance
        if len(maxH) > len(minH) + 1:
            heapq.heappush(minH, -heapq.heappop(maxH))
        elif len(minH) > len(maxH):
            heapq.heappush(maxH, -heapq.heappop(minH))
        # Median
        if len(maxH) == len(minH):
            print(f"Median: {(-maxH[0] + minH[0])/2.0}")
        else:
            print(f"Median: {-maxH[0]}")

if __name__ == "__main__":
    main() 