"""
╔════════════════════════════════════════════════════════╗
║          In‑Place LinkedList Reversal Pattern         ║
╠════════════════════════════════════════════════════════╣
║ Description:                                           ║
║ Reverse the `next` pointers by walking through the    ║
║ list and, for each node, redirecting its next to the  ║
║ previous node.                                        ║
╠════════════════════════════════════════════════════════╣
║ Flow Diagram (ASCII):                                  ║
║   Input: 1 → 2 → 3 → null                              ║
║                                                      ║
║   prev=null, cur=1                                   ║
║      cur->next = prev  → 1→null                       ║
║      prev=1, cur=2                                   ║
║                                                      ║
║   prev=1→null, cur=2                                 ║
║      2→next = 1→null  → 2→1→null                      ║
║      prev=2, cur=3                                   ║
║                                                      ║
║   prev=2→1→null, cur=3                               ║
║      3→next = 2→1→null  → 3→2→1→null                  ║
║   Done: new head = prev (3)                          ║
╚════════════════════════════════════════════════════════╝
"""

class Node:
    def __init__(self, val):
        self.val = val
        self.next = None

def reverse_list(head):
    prev, cur = None, head
    while cur:
        nxt = cur.next
        cur.next = prev
        prev = cur
        cur = nxt
    return prev

def main():
    # Build list: 1->2->3
    h = Node(1)
    h.next = Node(2)
    h.next.next = Node(3)
    r = reverse_list(h)
    while r:
        print(r.val, end=' ')
        r = r.next
    print()

if __name__ == "__main__":
    main() 