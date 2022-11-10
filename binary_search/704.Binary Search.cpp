#include <vector>
using namespace std;

class Solution {
public:
    int search(vector<int>& nums, int target) {
        int num, mid, left = 0;
        int right = nums.size() - 1;
        while(left <= right) {
            mid = (right - left)/2 + left;
            num = nums[mid];
            if(num == target)
                return mid;
            else if(num < target)
                left = mid + 1;
            else
                right = mid - 1;
        }
        return -1;
    }
};