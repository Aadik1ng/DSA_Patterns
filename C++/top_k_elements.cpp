#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

/*
╔════════════════════════════════════════════════════════╗
║               Top K Elements Pattern                  ║
╠════════════════════════════════════════════════════════╣
║ Description:                                           ║
║ Use nth_element to partition so that the first K      ║
║ elements are the top K (unordered).                  ║
╠════════════════════════════════════════════════════════╣
║ Flow Diagram (ASCII):                                  ║
║ arr=[3,2,1,5,6,4], k=2                                 ║
║ nth_element to place 2nd largest at index 2          ║
║ Partitioned: [6,5 | 1,2,3,4]                          ║
║ First 2 = [6,5]                                        ║
╚════════════════════════════════════════════════════════╝
*/

int main() {
    vector<int> arr = {3,2,1,5,6,4};
    int k = 2;

    // rearrange so that arr[0..k-1] are the k largest
    nth_element(arr.begin(), arr.begin()+k, arr.end(),
                greater<int>());

    cout << "Top " << k << ":";
    for (int i = 0; i < k; ++i)
        cout << " " << arr[i];
    cout << "\n";
    return 0;
} 