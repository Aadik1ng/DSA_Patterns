"""
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
"""

def main():
    nums = [1,2,3]
    subsets = [[]]
    for x in nums:
        sz = len(subsets)
        for i in range(sz):
            curr = list(subsets[i])
            curr.append(x)
            subsets.append(curr)
    # Print all subsets
    for s in subsets:
        print("[" + "".join(str(v) for v in s) + "]", end=' ')
    print()

if __name__ == "__main__":
    main() 