#include <iostream>
#include <vector>
#include <queue>
using namespace std;

struct TreeNode {
    int val; vector<TreeNode*> children;
    TreeNode(int v): val(v) {}
};

/*
╔════════════════════════════════════════════════════════╗
║               Tree Breadth‑First Search               ║
╠════════════════════════════════════════════════════════╣
║ Description:                                           ║
║ Level‑order traversal using a FIFO queue: push root,  ║
║ then for each node pop→visit→enqueue its children.    ║
╠════════════════════════════════════════════════════════╣
║ Flow Diagram (ASCII):                                  ║
║      1                                               ║
║     / \                                              ║
║    2   3                                             ║
║   /                                               ║
║  4                                                ║
║                                                      ║
║ Queue evolves:                                       ║
║  [1] → visit 1, enqueue [2,3]                        ║
║  [2,3] → visit 2, enqueue [3,4]                      ║
║  [3,4] → visit 3, enqueue [4]                        ║
║  [4] → visit 4, enqueue [] → done                    ║
╚════════════════════════════════════════════════════════╝
*/

void bfs(TreeNode* root) {
    if (!root) return;
    queue<TreeNode*> q;
    q.push(root);

    while (!q.empty()) {
        TreeNode* node = q.front(); q.pop();
        cout << node->val << " ";
        for (auto ch : node->children)
            q.push(ch);
    }
}

int main() {
    // Build tree: 1→{2,3}, 2→{4}
    TreeNode* root = new TreeNode(1);
    root->children = { new TreeNode(2), new TreeNode(3) };
    root->children[0]->children = { new TreeNode(4) };

    bfs(root);
    cout << "\n";
    return 0;
} 