"""
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
"""

def knapsack(W, wt, val):
    n = len(wt)
    dp = [0] * (W+1)
    for i in range(n):
        for w in range(W, wt[i]-1, -1):
            dp[w] = max(dp[w], val[i] + dp[w - wt[i]])
    return dp[W]

def main():
    weights = [1,3,4]
    values  = [15,50,60]
    capacity = 8
    print(f"Max value = {knapsack(capacity, weights, values)}")

if __name__ == "__main__":
    main() 