class Solution {
public:
    int minSubArrayLen(int target, vector<int>& nums) {
        int sum = 0, left = 0, right = 0, n = nums.size(), ans = n + 1;
        for(right; right < n; right++){
            sum += nums[right];
            while(sum - nums[left] >= target){
                sum -= nums[left];
                left += 1;
            }
            if(sum >= target) ans = min(ans, right - left + 1);
        }
        return ans <= n ? ans : 0;
    }
};