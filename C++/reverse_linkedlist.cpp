#include <iostream>
using namespace std;

struct Node {
    int val; Node* next;
    Node(int v): val(v), next(nullptr) {}
};

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

Node* reverseList(Node* head) {
    Node *prev = nullptr, *cur = head;

    while (cur) {
        Node* nxt = cur->next;     // save next
        cur->next = prev;          // reverse link
        prev = cur;                // advance prev
        cur = nxt;                 // advance cur
    }
    return prev;                   // new head
}

int main() {
    // Build list: 1->2->3
    Node* h = new Node(1);
    h->next = new Node(2);
    h->next->next = new Node(3);

    Node* rev = reverseList(h);
    while (rev) {
        cout << rev->val << " ";
        rev = rev->next;
    }
    cout << "\n";
    return 0;
} 