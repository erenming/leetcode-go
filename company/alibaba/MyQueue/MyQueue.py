class MyQueue(object):

    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.s_push = []
        self.s_pop = []

    def push(self, x):
        """
        Push element x to the back of queue.
        :type x: int
        :rtype: None
        """
        self.s_push.append(x)

    def pop(self):
        """
        Removes the element from in front of queue and returns that element.
        :rtype: int
        """
        self.__shift()
        return self.s_pop.pop(-1)

    def peek(self):
        """
        Get the front element.
        :rtype: int
        """
        self.__shift()
        return self.s_pop[-1]

    def empty(self):
        """
        Returns whether the queue is empty.
        :rtype: bool
        """
        return len(self.s_pop) == 0 and len(self.s_push) == 0

    def __shift(self):
        if len(self.s_pop) == 0:
            while len(self.s_push) != 0:
                self.s_pop.append(self.s_push.pop(-1))

# Your MyQueue object will be instantiated and called as such:
# obj = MyQueue()
# obj.push(x)
# param_2 = obj.pop()
# param_3 = obj.peek()
# param_4 = obj.empty()
