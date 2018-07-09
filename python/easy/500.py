class Solution:
    def findWords(self, words):
        """
        :type words: List[str]
        :rtype: List[str]
        """
        ans = []
        keyboard = [set('qwertyuiop'), set('asdfghjkl'), set('zxcvbnm')]
        return [word for row in keyboard for word in words if set(word.lower()).issubset(row)]
