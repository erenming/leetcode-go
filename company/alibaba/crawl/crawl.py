# """
# This is HtmlParser's API interface.
# You should not implement it, or speculate about its implementation
# """
# class HtmlParser(object):
#    def getUrls(self, url):
#        """
#        :type url: str
#        :rtype List[str]
#        """

from typing import List


class Solution:
    def crawl(self, startUrl: str, htmlParser: 'HtmlParser') -> List[str]:
        accessed = set()
        domain = self.get_domain(startUrl)
        q = list()
        q.append(startUrl)
        accessed.add(startUrl)
        while len(q) != 0:
            start = q.pop(0)
            urls = htmlParser().getUrls(start)
            for url in urls:
                if self.get_domain(url) == domain and url not in accessed:
                    q.append(url)
                    accessed.add(url)
        return list(accessed)

    def get_domain(self, url: str):
        url = url.lstrip("http://")
        domain = url.split("/")[0]
        return domain
