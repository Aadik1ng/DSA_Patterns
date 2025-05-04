/*
╔════════════════════════════════════════════════════════╗
║         Modified Binary Search (Rotated Array)        ║
╠════════════════════════════════════════════════════════╣
║ Description:                                           ║
║ Search in a rotated sorted array by comparing mid     ║
║ with start/end to determine which half is sorted and  ║
║ if target lies within that half.                      ║
╠════════════════════════════════════════════════════════╣
║ Flow Diagram (ASCII):                                  ║
║   arr = [4,5,6,7,0,1,2], target=0                     ║
║                                                      ║
║   start=0, end=6, mid=3                              ║
║   4 < 7 → left half sorted                           ║
║   0 not in [4,7] → search right half                 ║
║                                                      ║
║   start=4, end=6, mid=5                              ║
║   0 < 2 → right half sorted                          ║
║   0 in [0,2] → search left half                      ║
║                                                      ║
║   start=4, end=4 → found at index 4                  ║
╚════════════════════════════════════════════════════════╝
*/

package main

import "fmt"

func search(nums []int, target int) int {
    start, end := 0, len(nums)-1

    for start <= end {
        mid := start + (end-start)/2

        if nums[mid] == target {
            return mid
        }

        // Check if left half is sorted
        if nums[start] <= nums[mid] {
            // Check if target is in left half
            if target >= nums[start] && target < nums[mid] {
                end = mid - 1
            } else {
                start = mid + 1
            }
        } else {
            // Check if target is in right half
            if target > nums[mid] && target <= nums[end] {
                start = mid + 1
            } else {
                end = mid - 1
            }
        }
    }

    return -1
}

func main() {
    nums := []int{4, 5, 6, 7, 0, 1, 2}
    target := 0
    
    index := search(nums, target)
    if index != -1 {
        fmt.Printf("Found %d at index %d\n", target, index)
    } else {
        fmt.Printf("%d not found in array\n", target)
    }
} 