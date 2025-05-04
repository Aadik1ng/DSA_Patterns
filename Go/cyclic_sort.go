/*
╔════════════════════════════════════════════════════════╗
║                  Cyclic Sort Pattern                  ║
╠════════════════════════════════════════════════════════╣
║ Description:                                           ║
║ Numbers in [1..n] appear in an unsorted array with    ║
║ duplicates. Place each number at index = value-1 by   ║
║ swapping until all are either correct or duplicates.  ║
╠════════════════════════════════════════════════════════╣
║ Flow Diagram (ASCII):                                  ║
║   arr = [3,1,5,4,2,3]                                  ║
║     i=0: arr[0]=3, correct=2 → swap arr[0]<->arr[2]   ║
║   arr-> [5,1,3,4,2,3]                                  ║
║     i still=0: arr[0]=5, correct=4 → swap [0]<->[4]   ║
║   arr-> [2,1,3,4,5,3]                                  ║
║     i still=0: arr[0]=2, correct=1 → swap [0]<->[1]   ║
║   arr-> [1,2,3,4,5,3], i++ → …                          ║
║   final: positions 0–4 correct, idx5 holds duplicate 3 ║
║   missing: none, duplicate: 3                          ║
╚════════════════════════════════════════════════════════╝
*/

package main

import "fmt"

func main() {
    arr := []int{3, 1, 5, 4, 2, 3}
    i, n := 0, len(arr)

    // 1) Place each number in its correct index
    for i < n {
        correct := arr[i] - 1 // target index
        if arr[i] >= 1 && arr[i] <= n && arr[i] != arr[correct] {
            // Swap to put number in correct position
            arr[i], arr[correct] = arr[correct], arr[i]
        } else {
            i++ // Move on
        }
    }

    // 2) Find missing numbers
    fmt.Print("Missing numbers:")
    for j := 0; j < n; j++ {
        if arr[j] != j+1 {
            fmt.Printf(" %d", j+1)
        }
    }
    fmt.Println()
} 