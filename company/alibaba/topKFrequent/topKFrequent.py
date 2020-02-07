import heapq

class Solution(object):
    def topKFrequent(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: List[int]
        """
        mm = {}
        for v in nums:
            if v in mm:
                mm[v] = mm[v] + 1
            else:
                mm[v] = 1

        h = []
        for key, value in mm.items():
            if len(h) < k:
                heapq.heappush(h, (value, key))
            else:
                if h[0][0] < value:
                    heapq.heappop(h)
                    heapq.heappush(h, (value, key))
        res = [i[1] for i in h]
        res.reverse()
        return res


if __name__ == '__main__':
    Solution().topKFrequent([6,0,1,4,9,7,-3,1,-4,-8,4,-7,-3,3,2,-3,9,5,-4,0], 6)
