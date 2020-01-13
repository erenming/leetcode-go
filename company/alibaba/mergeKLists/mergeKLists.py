# Definition for singly-linked list.
import heapq


class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def mergeKLists(self, lists):
        dummy = ListNode(-1)
        cur = dummy
        pq = []
        for ll in lists:
            if ll:
                pq.append((ll.val, ll))
        heapq.heapify(pq)
        while len(pq) != 0:
            _, top = heapq.heappop(pq)
            cur.next = top
            cur = cur.next
            if top.next:
                heapq.heappush(pq, (top.next.val, top.next))
        return dummy.next
