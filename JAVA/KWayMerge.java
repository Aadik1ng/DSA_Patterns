import java.util.*;

/*
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
*/

class Item implements Comparable<Item> {
    int val, listIdx, elemIdx;
    Item(int v, int l, int e) { val = v; listIdx = l; elemIdx = e; }
    public int compareTo(Item o) { return Integer.compare(val, o.val); }
}

public class KWayMerge {
    public static void main(String[] args) {
        int[][] lists = {{1,4,7},{2,5,8},{3,6,9}};
        PriorityQueue<Item> minH = new PriorityQueue<>();
        for (int i = 0; i < lists.length; ++i)
            if (lists[i].length > 0)
                minH.add(new Item(lists[i][0], i, 0));
        List<Integer> result = new ArrayList<>();
        while (!minH.isEmpty()) {
            Item cur = minH.poll();
            result.add(cur.val);
            if (cur.elemIdx + 1 < lists[cur.listIdx].length)
                minH.add(new Item(lists[cur.listIdx][cur.elemIdx+1], cur.listIdx, cur.elemIdx+1));
        }
        for (int x : result) System.out.print(x + " ");
        System.out.println();
    }
} 