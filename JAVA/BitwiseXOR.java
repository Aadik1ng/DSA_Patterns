import java.util.*;

/*
╔════════════════════════════════════════════════════════╗
║               Bitwise XOR Pattern (Single Number)     ║
╠════════════════════════════════════════════════════════╣
║ Description:                                           ║
║ XOR all numbers; since x⊕x=0 and x⊕0=x, pairs cancel ║
║ leaving the unique number.                            ║
╠════════════════════════════════════════════════════════╣
║ Flow Diagram (ASCII):                                  ║
║ arr=[4,1,2,1,2]                                       ║
║  res=0⊕4=4;                                           ║
║  res=4⊕1=5;                                           ║
║  res=5⊕2=7;                                           ║
║  res=7⊕1=6;                                           ║
║  res=6⊕2=4;                                           ║
║  → result=4                                           ║
╚════════════════════════════════════════════════════════╝
*/

public class BitwiseXOR {
    public static int singleNumber(int[] nums) {
        int result = 0;
        for (int x : nums)
            result ^= x; // xor accumulate
        return result;
    }
    public static void main(String[] args) {
        int[] arr = {4,1,2,1,2};
        System.out.println("Single number = " + singleNumber(arr));
    }
} 