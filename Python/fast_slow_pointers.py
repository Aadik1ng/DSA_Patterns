"""
╔════════════════════════════════════════════════════════╗
║         Fast & Slow Pointers (Cycle Detection)        ║
╠════════════════════════════════════════════════════════╣
║ Description:                                           ║
║ Use two pointers moving at different speeds—slow      ║
║ moves 1 step, fast moves 2 steps. In a cycle, they   ║
║ must eventually meet.                                 ║
╠════════════════════════════════════════════════════════╣
║ Flow Diagram (ASCII):                                  ║
║   a -> b -> c                                         ║
║         ↑    ↓                                        ║
║         |____| ← fast loops back                        ║
║   slow: a→b→c→b→c→…                                    ║
║   fast: a→c→b→c→b→…                                    ║
║   => slow == fast                                       ║
╚════════════════════════════════════════════════════════╝
"""

class ListNode:
    def __init__(self, val):
        self.val = val
        self.next = None

def has_cycle(head):
    slow = fast = head
    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next
        if slow == fast:
            return True
    return False

def main():
    # Build a cycle: a->b->c->b...
    a = ListNode(1)
    b = ListNode(2)
    c = ListNode(3)
    a.next = b; b.next = c; c.next = b
    print("Cycle detected" if has_cycle(a) else "No cycle")

if __name__ == "__main__":
    main() 