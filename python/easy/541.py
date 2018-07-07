class Solution:
    def reverseStr(self, s, k):
        """
        :type s: str
        :type k: int
        :rtype: str
        """
        ans = ''
        end = 0
        while (end + 2*k) < len(s):
            ans += (s[end:end+k:][::-1] + s[end+k:end+2*k:])
            end += 2*k
        if end != len(s):
            if len(s) - end > k:
                ans += s[end:end+k:][::-1]
                ans += s[end+k::]
            else:   ans+= s[end::][::-1]
        else: ans += (s[end:end+k:][::-1] + s[end+k:end+2*k:])
        return ans