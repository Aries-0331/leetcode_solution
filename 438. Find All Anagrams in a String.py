# sliding window
class Solution:
    def findAnagrams(self, s: str, p: str) -> List[int]:
        sLen, pLen = len(s), len(p)
        if sLen < pLen:
            return []

        res = []
        sCount = [0] * 26
        pCount = [0] * 26
        for i in range(pLen):
            sCount[ord(s[i]) - 97] += 1
            pCount[ord(p[i]) - 97] += 1

        if sCount == pCount:
            res.append(0)

        for i in range(sLen - pLen):
            sCount[ord(s[i]) - 97] -= 1
            sCount[ord(s[i + pLen]) - 97] += 1

            if sCount == pCount:
                res.append(i+1)
        return res