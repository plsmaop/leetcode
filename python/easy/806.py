class Solution:
    def numberOfLines(self, widths, S):
        """
        :type widths: List[int]
        :type S: str
        :rtype: List[int]
        """
        line = 0
        sum = 0
        for s in S:
            width = widths[ord(s) - 97]
            if sum + width > 100:
                line += 1
                sum = width
            else: sum += width
        return [line+1, sum % 100]