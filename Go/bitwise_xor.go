/*
╔════════════════════════════════════════════════════════╗
║              Bitwise XOR (Single Number)              ║
╠════════════════════════════════════════════════════════╣
║ Description:                                           ║
║ Use XOR to find the single number in an array where   ║
║ all other numbers appear twice. XOR cancels out       ║
║ duplicates: a^a=0, a^0=a.                            ║
╠════════════════════════════════════════════════════════╣
║ Flow Diagram (ASCII):                                  ║
║   nums = [4,1,2,1,2]                                  ║
║                                                      ║
║   result = 0                                          ║
║   0 ^ 4 = 4                                           ║
║   4 ^ 1 = 5                                           ║
║   5 ^ 2 = 7                                           ║
║   7 ^ 1 = 6                                           ║
║   6 ^ 2 = 4 → single number found                     ║
╚════════════════════════════════════════════════════════╝
*/

package main

import "fmt"

func singleNumber(nums []int) int {
    result := 0
    for _, num := range nums {
        result ^= num
    }
    return result
}

func main() {
    nums := []int{4, 1, 2, 1, 2}
    single := singleNumber(nums)
    fmt.Printf("Single number: %d\n", single)
} 