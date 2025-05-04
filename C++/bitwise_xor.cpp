#include <iostream>
#include <vector>
using namespace std;

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

int singleNumber(const vector<int>& nums) {
    int result = 0;
    for (int x : nums)
        result ^= x;           // xor accumulate
    return result;
}

int main() {
    vector<int> arr = {4,1,2,1,2};
    cout << "Single number = " << singleNumber(arr) << "\n";
    return 0;
} 