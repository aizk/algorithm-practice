class Solution(object):
    def isValid(self, s):
        """
        :type s: str
        :rtype: bool
        """
        stack = []
        for c in s:
            if c in '([{':
                stack.append(c)
            if c in ')]}':
                # 边界条件，为空弹不出括号，必然无法匹配
                if not stack:
                    return False
                y = stack.pop() + c
                if y not in ['()', '[]', '{}']:
                    return False
        if not stack:
            return True
        else:
            return False


if __name__ == '__main__':
    s = Solution()
    print(s.isValid('['))
