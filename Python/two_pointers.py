"""
╔════════════════════════════════════════════════════════╗
║            Two‑Pointer Pattern (Sorted Array)          ║
╠════════════════════════════════════════════════════════╣
║ Description:                                           ║
║ Use two indices—one at start, one at end—to find a    ║
║ pair whose sum equals the target in O(N) time.        ║
╠════════════════════════════════════════════════════════╣
║ Flow Diagram (ASCII):                                  ║
║   arr = [1,2,3,4,6], target=6                          ║
║                                                      ║
║     left=0           right=4                           ║
║       ↓                ↑                               ║
║       1  2  3  4  6                                      ║
║       |--------------|   sum=1+6=7>6 → decrement right ║
║                                                      ║
║     left=0         right=3                             ║
║       ↓              ↑                                 ║
║       1  2  3  4  6    sum=1+4=5<6 → increment left    ║
║                                                      ║
║     left=1       right=3                               ║
║        ↓           ↑                                   ║
║       1  2  3  4  6      sum=2+4=6 → found             ║
╚════════════════════════════════════════════════════════╝
"""

def main():
    arr = [1, 2, 3, 4, 6]
    target = 6
    left, right = 0, len(arr) - 1

    # Move inward until pointers meet
    while left < right:
        s = arr[left] + arr[right]
        if s == target:
            print(f"Pair: ({arr[left]},{arr[right]})")
            return
        if s < target:
            left += 1
        else:
            right -= 1
    print("No valid pair")

if __name__ == "__main__":
    main() 