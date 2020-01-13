import heapq
from typing import List


# heap implementation
class Solution:
    def findKthLargest(self, nums: List[int], k: int) -> int:
        h = []
        for i in nums:
            if len(h) < k:
                heapq.heappush(h, i)
                continue
            if i > h[0]:
                heapq.heappushpop(h, i)
                continue
        return h[0]


if __name__ == '__main__':
    Solution().findKthLargest([3, 2, 1, 5, 6, 4], 2)
