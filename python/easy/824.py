class Solution:
    def toGoatLatin(self, S):
        """
        :type S: str
        :rtype: str
        """
        ans = []
        for index, s in enumerate(S.split(' ')):
            append = 'a'*(index+1)
            tmp = ''
            if s[0] in 'aeiouAEIOU':    tmp += s
            else:   tmp += (s[1:] + s[0])
            tmp += ('ma' + append)
            ans.append(tmp)
        return ' '.join(ans)