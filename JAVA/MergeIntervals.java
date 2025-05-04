import java.util.*;

/*
╔════════════════════════════════════════════════════════╗
║               Merge Intervals Pattern                 ║
╠════════════════════════════════════════════════════════╣
║ Description:                                           ║
║ Given a list of intervals, sort by start time and     ║
║ merge overlapping ones by comparing the current end   ║
║ with the next start.                                   ║
╠════════════════════════════════════════════════════════╣
║ Flow Diagram (ASCII):                                  ║
║   Input: [[1,3],[2,6],[8,10],[15,18]]                 ║
║                                                      ║
║   After Sort: [[1,3],[2,6],[8,10],[15,18]]            ║
║     merged = [1,3]                                    ║
║     next = [2,6]  → overlap? 3 >= 2 → merge end=max(6,3)=6 ║
║     merged = [1,6]                                    ║
║     next = [8,10] → 6 < 8 → push [8,10]                ║
║     …                                                  ║
║   Output: [[1,6],[8,10],[15,18]]                      ║
╚════════════════════════════════════════════════════════╝
*/

public class MergeIntervals {
    public static void main(String[] args) {
        int[][] intervals = {{1,3},{2,6},{8,10},{15,18}};
        Arrays.sort(intervals, Comparator.comparingInt(a -> a[0]));
        List<int[]> merged = new ArrayList<>();
        for (int[] in : intervals) {
            if (merged.isEmpty() || merged.get(merged.size()-1)[1] < in[0]) {
                merged.add(in);
            } else {
                merged.get(merged.size()-1)[1] = Math.max(merged.get(merged.size()-1)[1], in[1]);
            }
        }
        for (int[] in : merged)
            System.out.print("[" + in[0] + "," + in[1] + "] ");
        System.out.println();
    }
} 