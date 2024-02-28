class Solution {
public:
    int numSubarrayProductLessThanK(vector<int>& nums, int k) {
        if(k <= 1) return 0;
        int ans = 0, left = 0, prod = 1, n = nums.size();
        for(int right = 0; right < n; right++){
            prod *= nums[right];
            while(prod >= k){
                prod /= nums[left];
                left += 1;
            }
            ans += right - left + 1;
        }
        return ans;
    }
};