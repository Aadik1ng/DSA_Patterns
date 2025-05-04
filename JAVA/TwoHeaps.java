import java.util.*;

/*
╔════════════════════════════════════════════════════════╗
║              Two Heaps Pattern (Median)               ║
╠════════════════════════════════════════════════════════╣
║ Description:                                           ║
║ Maintain a max‑heap for the lower half and a min‑heap ║
║ for the upper half of the stream to get median in O(log N) per insert. ║
╠════════════════════════════════════════════════════════╣
║ Flow Diagram (ASCII):                                  ║
║ Stream: 5,15,1,3                                      ║
║                                                      ║
║ Insert 5 → maxH:[5], minH:[] → median=5               ║
║ Insert 15→ maxH:[5], minH:[15] → median=(5+15)/2=10   ║
║ Insert 1 → maxH:[5,1], minH:[15] → median=5           ║
║ Insert 3 → maxH:[3,1], minH:[5,15] → median=(3+5)/2=4 ║
╚════════════════════════════════════════════════════════╝
*/

public class TwoHeaps {
    public static void main(String[] args) {
        PriorityQueue<Integer> maxH = new PriorityQueue<>(Collections.reverseOrder()); // lower half
        PriorityQueue<Integer> minH = new PriorityQueue<>(); // upper half
        int[] nums = {5, 15, 1, 3};
        for (int x : nums) {
            // 1) Add new number to one of the heaps
            if (maxH.isEmpty() || x < maxH.peek())
                maxH.add(x);
            else
                minH.add(x);
            // 2) Rebalance to ensure size difference ≤1
            if (maxH.size() > minH.size() + 1) {
                minH.add(maxH.poll());
            } else if (minH.size() > maxH.size()) {
                maxH.add(minH.poll());
            }
            // 3) Compute median
            if (maxH.size() == minH.size())
                System.out.println("Median: " + ((maxH.peek() + minH.peek())/2.0));
            else
                System.out.println("Median: " + maxH.peek());
        }
    }
} 