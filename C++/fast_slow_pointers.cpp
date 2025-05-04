#include <iostream>
using namespace std;

struct ListNode {
    int val; ListNode* next;
    ListNode(int v): val(v), next(nullptr) {}
};

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

// Returns true if cycle exists
bool hasCycle(ListNode* head) {
    ListNode *slow = head, *fast = head;
    while (fast && fast->next) {
        slow = slow->next;           // 1 step
        fast = fast->next->next;     // 2 steps
        if (slow == fast) return true;
    }
    return false;
}

int main() {
    // Build a cycle: a->b->c->b...
    ListNode *a = new ListNode(1),
             *b = new ListNode(2),
             *c = new ListNode(3);
    a->next = b; b->next = c; c->next = b;

    cout << (hasCycle(a) ? "Cycle detected\n" : "No cycle\n");
    return 0;
} 