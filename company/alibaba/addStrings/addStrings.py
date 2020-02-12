class Solution(object):
    def addStrings(self, num1, num2):
        """
        :type num1: str
        :type num2: str
        :rtype: str
        """
        num1 = list(num1)
        num2 = list(num2)
        carry, res = 0, []
        while len(num1) > 0 or len(num2) > 0:
            n1 = ord(num1.pop()) - ord('0') if len(num1) > 0 else 0
            n2 = ord(num2.pop()) - ord('0') if len(num2) > 0 else 0
            tmp = carry + n1 + n2
            res.append(tmp % 10)
            carry = tmp // 10
        if carry > 0:
            res.append(carry)
        return ''.join([str(i) for i in res])[::-1]

if __name__ == '__main__':
    print Solution().addStrings("9", "99")