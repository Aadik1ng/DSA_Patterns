import java.util.*;

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

class TreeNode {
    int val; List<TreeNode> children = new ArrayList<>();
    TreeNode(int v) { val = v; }
}

public class TreeBFS {
    // Level-order traversal using queue
    public static void bfs(TreeNode root) {
        if (root == null) return;
        Queue<TreeNode> q = new LinkedList<>();
        q.add(root);
        while (!q.isEmpty()) {
            TreeNode node = q.poll();
            System.out.print(node.val + " ");
            for (TreeNode ch : node.children)
                q.add(ch);
        }
    }
    public static void main(String[] args) {
        // Build tree: 1→{2,3}, 2→{4}
        TreeNode root = new TreeNode(1);
        root.children.add(new TreeNode(2));
        root.children.add(new TreeNode(3));
        root.children.get(0).children.add(new TreeNode(4));
        bfs(root);
        System.out.println();
    }
} 