class Solution:
    def subdomainVisits(self, cpdomains):
        """
        :type cpdomains: List[str]
        :rtype: List[str]
        """
        ans = {}
        for i in cpdomains:
            times, domains = i.split(' ')
            domain = domains.split('.')
            url = ''
            for j in reversed(domain):
                if len(url) == 0:
                    url = j
                else:
                    url = j + '.' + url
                if url not in ans:
                    ans[url] = int(times)
                else:
                    ans[url] += int(times)
        
        return [str(value) + " " + str(key) for key, value in ans.items()]
            