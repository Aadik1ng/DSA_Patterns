"""
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
"""

def main():
    intervals = [[1,3],[2,6],[8,10],[15,18]]
    intervals.sort()
    merged = []
    for in_ in intervals:
        if not merged or merged[-1][1] < in_[0]:
            merged.append(in_)
        else:
            merged[-1][1] = max(merged[-1][1], in_[1])
    for in_ in merged:
        print(f"[{in_[0]},{in_[1]}]", end=' ')
    print()

if __name__ == "__main__":
    main() 