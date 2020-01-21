# Definition for a Node.
class Node:
    def __init__(self, x, next=None, random=None):
        self.val = int(x)
        self.next = next
        self.random = random


class Solution(object):
    def copyRandomList(self, head):
        """
        :type head: Node
        :rtype: Node
        """
        if not head:
            return None
        m = {}
        i = j = head
        while i:
            m[i] = Node(i.val)
            i = i.next

        while j:
            m[j].next = m.get(j.next)
            m[j].random = m.get(j.random)
            j = j.next
        return m[head]
