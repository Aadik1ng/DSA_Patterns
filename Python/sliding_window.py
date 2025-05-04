"""
╔════════════════════════════════════════════════════════╗
║            Sliding Window Pattern (Size K)            ║
╠════════════════════════════════════════════════════════╣
║ Description:                                           ║
║ Maintain a running sum of a K‑sized window and slide   ║
║ across the array in O(N) time by subtracting the      ║
║ element that "goes out" and adding the new one "in".   ║
╠════════════════════════════════════════════════════════╣
║ Flow Diagram (ASCII):                                  ║
║ Array: [1,3,2,6,-1,4,1,8,2], K=5                       ║
║                                                      ║
║   Initial Window:                                      ║
║   Indexes:   0   1   2   3    4    5    6    7   8     ║
║   Elements: [1] – [3] – [2] – [6] – [ -1 ]   4    1    8  2 ║
║              |--------------------------------|         ║
║                                                      ║
║   Slide Right by 1:                                    ║
║   Remove arr[0], Add arr[5]                            ║
║   [3] – [2] – [6] – [ -1 ] – [4]                        ║
║                 |--------------------------------|       ║
╚════════════════════════════════════════════════════════╝
"""

def main():
    arr = [1,3,2,6,-1,4,1,8,2]
    K = 5  # Window size
    n = len(arr)
    if n < K:
        return

    # 1) Sum of first K elements
    window_sum = sum(arr[:K])
    print(f"Avg [0..{K-1}] = {window_sum / K}")

    # 2) Slide the window one element at a time
    for end in range(K, n):
        window_sum += arr[end] - arr[end - K]
        print(f"Avg [{end-K+1}..{end}] = {window_sum / K}")

if __name__ == "__main__":
    main() 