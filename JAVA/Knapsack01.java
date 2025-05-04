import java.util.*;

/*
╔════════════════════════════════════════════════════════╗
║          0/1 Knapsack Pattern (Dynamic Programming)   ║
╠════════════════════════════════════════════════════════╣
║ Description:                                           ║
║ Build a 1D DP array dp[w] = max value for capacity w. ║
║ For each item, iterate w from W→weight to avoid reuse.║
╠════════════════════════════════════════════════════════╣
║ Flow Diagram (ASCII):                                  ║
║ weights=[1,3,4], values=[15,50,60], W=8               ║
║ dp initially [0,0,0,0,0,0,0,0,0]                     ║
║ Item0(w=1,v=15):                                     ║
║  dp[8]=max(dp[8],15+dp[7]) → … → dp[1]=15              ║
║ Item1(w=3,v=50):                                     ║
║  …                                                   ║
║ Item2(w=4,v=60):                                     ║
║  …                                                   ║
║ Result=dp[8]=80 (items 1 and 2)                       ║
╚════════════════════════════════════════════════════════╝
*/

public class Knapsack01 {
    public static int knapsack(int W, int[] wt, int[] val) {
        int n = wt.length;
        int[] dp = new int[W+1];
        for (int i = 0; i < n; ++i) {
            for (int w = W; w >= wt[i]; --w) {
                dp[w] = Math.max(dp[w], val[i] + dp[w - wt[i]]);
            }
        }
        return dp[W];
    }
    public static void main(String[] args) {
        int[] weights = {1,3,4};
        int[] values  = {15,50,60};
        int capacity = 8;
        System.out.println("Max value = " + knapsack(capacity, weights, values));
    }
} 