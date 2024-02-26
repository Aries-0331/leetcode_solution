class Solution {
public:
    vector<vector<int>> threeSum(vector<int>& nums) {
        vector<vector<int>> ans;
        sort(nums.begin(), nums.end());
        int sum = 0, n = nums.size();
        for (int i = 0; i < n - 2; i++){
            if(nums[i] + nums[n-1] + nums[n-2] < 0) continue;
            if(nums[i] + nums[i+1] + nums[i+2] > 0) break;
            if(i > 0 && nums[i] == nums[i-1]) continue;

            int left = i + 1, right = n - 1;
            while (left < right){
                sum = nums[i] + nums[left] + nums[right];
                if(sum < 0) left++;
                else if(sum > 0) right--;
                else{
                    ans.push_back({nums[i], nums[left], nums[right]});
                    for(++left; left < right && nums[left] == nums[left-1]; ++left);
                    for(--right; right > left && nums[right] == nums[right+1]; --right);
                }
            }
        }
        return ans;
    }
};