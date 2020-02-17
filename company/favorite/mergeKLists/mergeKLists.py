import heapq


# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    def mergeKLists(self, lists):
        """
        :type lists: List[ListNode]
        :rtype: ListNode
        """
        dummy = ListNode(-1)
        cur = dummy
        h = []
        for i in lists:
            if i:
                h.append((i.val, i))
        heapq.heapify(h)
        while len(h) != 0:
            _, item = heapq.heappop(h)
            cur.next = item
            cur = cur.next
            if item.next:
                heapq.heappush(h, (item.next.val, item.next))
        return dummy.next
