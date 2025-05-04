import java.util.*;

/*
╔════════════════════════════════════════════════════════╗
║         Modified Binary Search (Rotated Array)        ║
╠════════════════════════════════════════════════════════╣
║ Description:                                           ║
║ On a rotated sorted array, decide which half is       ║
║ normally ordered each step, then narrow search to     ║
║ the half that could contain the target.               ║
╠════════════════════════════════════════════════════════╣
║ Flow Diagram (ASCII):                                  ║
║ arr=[4,5,6,7,0,1,2], target=0                          ║
║ lo=0, hi=6 → mid=3 (7)                                 ║
║  [4,5,6,7] is sorted, target∉[4..7] → search right     ║
║ lo=4, hi=6 → mid=5 (1)                                 ║
║  [0,1,2] is sorted, target∈[0..2] → hi=mid-1=4         ║
║ lo=4, hi=4 → mid=4 (0) → found                         ║
╚════════════════════════════════════════════════════════╝
*/

public class ModifiedBinarySearch {
    public static int search(int[] nums, int target) {
        int lo = 0, hi = nums.length - 1;
        while (lo <= hi) {
            int mid = lo + (hi - lo) / 2;
            if (nums[mid] == target) return mid;
            // Determine which half is sorted
            if (nums[lo] <= nums[mid]) {
                // left half sorted
                if (target >= nums[lo] && target < nums[mid])
                    hi = mid - 1;
                else
                    lo = mid + 1;
            } else {
                // right half sorted
                if (target > nums[mid] && target <= nums[hi])
                    lo = mid + 1;
                else
                    hi = mid - 1;
            }
        }
        return -1; // not found
    }
    public static void main(String[] args) {
        int[] arr = {4,5,6,7,0,1,2};
        System.out.println("Index of 0 is " + search(arr, 0));
    }
} 