"""
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
"""

def single_number(nums):
    result = 0
    for x in nums:
        result ^= x
    return result

def main():
    arr = [4,1,2,1,2]
    print(f"Single number = {single_number(arr)}")

if __name__ == "__main__":
    main() 