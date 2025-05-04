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

package main

import (
    "fmt"
    "sort"
)

type Interval struct {
    Start, End int
}

func main() {
    intervals := []Interval{
        {1, 3},
        {2, 6},
        {8, 10},
        {15, 18},
    }

    // 1) Sort by start time
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i].Start < intervals[j].Start
    })

    var merged []Interval
    for _, interval := range intervals {
        if len(merged) == 0 || merged[len(merged)-1].End < interval.Start {
            // No overlap
            merged = append(merged, interval)
        } else {
            // Overlap → extend end
            if merged[len(merged)-1].End < interval.End {
                merged[len(merged)-1].End = interval.End
            }
        }
    }

    // 2) Print result
    for _, interval := range merged {
        fmt.Printf("[%d,%d] ", interval.Start, interval.End)
    }
    fmt.Println()
} 