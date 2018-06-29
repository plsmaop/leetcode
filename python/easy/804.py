class Solution:
    def uniqueMorseRepresentations(self, words):
        """
        :type words: List[str]
        :rtype: int
        """
        transfer = [".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."]
        
        table = set({})
        
        for word in words:
            code = ''
            for i in word:
                code += transfer[(ord(i) - 97)]
            if code not in table:
                table.add(code)
        
        return len(table)