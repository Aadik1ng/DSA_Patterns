#include <iostream>
#include <vector>
using namespace std;

struct TreeNode {
    int val; vector<TreeNode*> children;
    TreeNode(int v): val(v) {}
};

/*
╔════════════════════════════════════════════════════════╗
║             Tree Depth‑First Search (DFS)             ║
╠════════════════════════════════════════════════════════╣
║ Description:                                           ║
║ Pre‑order traversal via recursion (or a stack):       ║
║ visit node, then recurse into each child.            ║
╠════════════════════════════════════════════════════════╣
║ Flow Diagram (ASCII):                                  ║
║      1                                               ║
║     / \                                              ║
║    2   3 → 5                                        ║
║                                                      ║
║ Visit order: 1 → 2 → (back) → 3 → 5                  ║
╚════════════════════════════════════════════════════════╝
*/

void dfs(TreeNode* node) {
    if (!node) return;
    cout << node->val << " ";        // visit
    for (auto ch : node->children)
        dfs(ch);                     // recurse
}

int main() {
    // Build tree: 1→{2,3}, 3→{5}
    TreeNode* root = new TreeNode(1);
    root->children = { new TreeNode(2), new TreeNode(3) };
    root->children[1]->children = { new TreeNode(5) };

    dfs(root);
    cout << "\n";
    return 0;
} 