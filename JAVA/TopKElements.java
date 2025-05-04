import java.util.*;

/*
╔════════════════════════════════════════════════════════╗
║               Top K Elements Pattern                  ║
╠════════════════════════════════════════════════════════╣
║ Description:                                           ║
║ Use nth_element to partition so that the first K      ║
║ elements are the top K (unordered).                  ║
╠════════════════════════════════════════════════════════╣
║ Flow Diagram (ASCII):                                  ║
║ arr=[3,2,1,5,6,4], k=2                                 ║
║ nth_element to place 2nd largest at index 2          ║
║ Partitioned: [6,5 | 1,2,3,4]                          ║
║ First 2 = [6,5]                                        ║
╚════════════════════════════════════════════════════════╝
*/

public class TopKElements {
    public static void main(String[] args) {
        int[] arr = {3,2,1,5,6,4};
        int k = 2;
        // Use a min-heap of size k
        PriorityQueue<Integer> minH = new PriorityQueue<>();
        for (int x : arr) {
            minH.add(x);
            if (minH.size() > k) minH.poll();
        }
        System.out.print("Top " + k + ":");
        List<Integer> result = new ArrayList<>(minH);
        Collections.sort(result, Collections.reverseOrder());
        for (int x : result)
            System.out.print(" " + x);
        System.out.println();
    }
} 