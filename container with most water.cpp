class Solution {
public:
    int maxArea(vector<int>& height) {
        int left = 0, right = height.size() - 1;
        int current = 0, ans = 0, temp = 0;
        while (left < right){
            current = min(height[left], height[right]) * (right - left);
            ans = max(current, ans);
            (height[left] <= height[right]) ? left++ : right--;
        }
        return ans;
    }
};