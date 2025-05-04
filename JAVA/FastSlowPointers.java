/*
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
*/

class ListNode {
    int val;
    ListNode next;
    ListNode(int v) { val = v; }
}

public class FastSlowPointers {
    // Returns true if cycle exists
    public static boolean hasCycle(ListNode head) {
        ListNode slow = head, fast = head;
        while (fast != null && fast.next != null) {
            slow = slow.next;           // 1 step
            fast = fast.next.next;      // 2 steps
            if (slow == fast) return true;
        }
        return false;
    }

    public static void main(String[] args) {
        // Build a cycle: a->b->c->b...
        ListNode a = new ListNode(1);
        ListNode b = new ListNode(2);
        ListNode c = new ListNode(3);
        a.next = b; b.next = c; c.next = b;

        System.out.println(hasCycle(a) ? "Cycle detected" : "No cycle");
    }
} 