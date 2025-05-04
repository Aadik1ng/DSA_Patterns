"""
╔════════════════════════════════════════════════════════╗
║                 K‑way Merge Pattern                  ║
╠════════════════════════════════════════════════════════╣
║ Description:                                           ║
║ Push the first element of each of the K sorted lists ║
║ into a min‑heap. Pop the smallest, then push its      ║
║ successor from the same list, repeating until empty. ║
╠════════════════════════════════════════════════════════╣
║ Flow Diagram (ASCII):                                  ║
║ lists=[[1,4,7],[2,5,8],[3,6,9]]                       ║
║                                                       ║
║ Heap init: [1@L0, 2@L1, 3@L2]                         ║
║ Pop 1@L0 → output 1, push 4@L0                         ║
║ Heap: [2@L1,3@L2,4@L0]                                 ║
║ Pop 2@L1 → output 2, push 5@L1                         ║
║ …                                                     ║
║ Final output: 1,2,3,4,5,6,7,8,9                       ║
╚════════════════════════════════════════════════════════╝
"""

import heapq

def main():
    lists = [[1,4,7],[2,5,8],[3,6,9]]
    minH = []
    for i, lst in enumerate(lists):
        if lst:
            heapq.heappush(minH, (lst[0], i, 0))
    result = []
    while minH:
        val, list_idx, elem_idx = heapq.heappop(minH)
        result.append(val)
        if elem_idx + 1 < len(lists[list_idx]):
            heapq.heappush(minH, (lists[list_idx][elem_idx+1], list_idx, elem_idx+1))
    print(*result)

if __name__ == "__main__":
    main() 