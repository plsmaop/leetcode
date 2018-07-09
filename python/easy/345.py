class Solution:
    def reverseVowels(self, s):
        """
        :type s: str
        :rtype: str
        """
        vowel = ''
        ind = []
        index = 0
        for i in s:
            case = i.lower()
            if case == 'a' or case == 'e' or case == 'i' or case == 'o' or case == 'u':
                vowel += i
                ind.append(index)
            index += 1
        index = 0
        s_list = list(s)
        for i in reversed(vowel):
            s_list[ind[index]] = i
            index += 1
        return ''.join(s_list)