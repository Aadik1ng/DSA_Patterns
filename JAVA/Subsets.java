import java.util.*;

/*
╔════════════════════════════════════════════════════════╗
║                  Subsets Pattern (BFS)                ║
╠════════════════════════════════════════════════════════╣
║ Description:                                           ║
║ Start with empty subset, then for each number, take   ║
║ all existing subsets, append the number, and add to  ║
║ the list (BFS‑style build).                           ║
╠════════════════════════════════════════════════════════╣
║ Flow Diagram (ASCII):                                  ║
║ nums=[1,2,3]                                          ║
║ subsets=[[]]                                          ║
║ → for 1: new=[[1]]      → subsets=[[],[1]]            ║
║ → for 2: new=[[2],[1,2]]→ subsets=[[],[1],[2],[1,2]]  ║
║ → for 3: new=[[3],[1,3],[2,3],[1,2,3]] → final       ║
╚════════════════════════════════════════════════════════╝
*/

public class Subsets {
    public static void main(String[] args) {
        int[] nums = {1,2,3};
        List<List<Integer>> subsets = new ArrayList<>();
        subsets.add(new ArrayList<>());
        for (int x : nums) {
            int sz = subsets.size();
            for (int i = 0; i < sz; ++i) {
                List<Integer> curr = new ArrayList<>(subsets.get(i));
                curr.add(x);
                subsets.add(curr);
            }
        }
        // Print all subsets
        for (List<Integer> s : subsets) {
            System.out.print("[");
            for (int v : s) System.out.print(v);
            System.out.print("] ");
        }
        System.out.println();
    }
} 