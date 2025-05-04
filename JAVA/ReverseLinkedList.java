/*
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
*/

class Node {
    int val; Node next;
    Node(int v) { val = v; }
}

public class ReverseLinkedList {
    // Reverse the linked list in-place
    public static Node reverseList(Node head) {
        Node prev = null, cur = head;
        while (cur != null) {
            Node nxt = cur.next;
            cur.next = prev;
            prev = cur;
            cur = nxt;
        }
        return prev;
    }

    public static void main(String[] args) {
        // Build list: 1->2->3
        Node h = new Node(1);
        h.next = new Node(2);
        h.next.next = new Node(3);
        Node r = reverseList(h);
        while (r != null) {
            System.out.print(r.val + " ");
            r = r.next;
        }
        System.out.println();
    }
} 