class Solution(object):
    def addBinary(self, a, b):
        """
        :type a: str
        :type b: str
        :rtype: str
        """
        a, b = list(a), list(b)
        carry, res = 0, []
        while len(a) > 0 or len(b) > 0:
            n1 = ord(a.pop()) - ord('0') if len(a) > 0 else 0
            n2 = ord(b.pop()) - ord('0') if len(b) > 0 else 0
            tmp = carry + n1 + n2
            res.append(tmp % 2)
            carry = tmp // 2
        if carry:
            res.append(carry)
        return ''.join([str(i) for i in res])[::-1]
