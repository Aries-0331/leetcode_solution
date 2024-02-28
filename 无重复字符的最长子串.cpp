class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        int ans = 0, left = 0, n = s.length();
        unordered_map<char, int> mp;
        for(int right = 0; right < n; right++){
            mp[s[right]] += 1;
            while(mp[s[right]] > 1){
                mp[s[left]] -= 1;
                left += 1;
            }
            ans = max(ans, right - left + 1);
        }
        return ans;
    }
};