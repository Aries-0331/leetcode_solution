#sliding window
class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        if not s: return 0
        left = 0
        lookup = set()
        n = len(s)
        maxLen = 0
        curLen = 0
        for i in range(n):
            curLen += 1
            while s[i] in lookup:
                lookup.remove(s[left])
                left += 1
                curLen -= 1
            if curLen > maxLen: maxLen = curLen
            lookup.add(s[i])
        return maxLen