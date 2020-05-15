INT_MAX = 2147483647
INT_MIN = -2147483648

class Solution(object):
    def myAtoi(self, str):
        """
        :type str: str
        :rtype: int
        """
        s = str.strip()
        if len(s) == 0: return 0
        neg = False
        # 判断符号
        fc = s[0]
        if fc == '+': s = s[1:]
        if fc == '-':
            s = s[1:]
            neg = True
        
        n = 0
        for ch in s:
            c = ord(ch) - 48
            if c > 9 or c < 0:
                # 不是数字字符 空格 32 小于 48
                break
            # 处理超大的数
            if INT_MAX / 10 < n or (INT_MAX/10 == n and INT_MAX % 10 < c):
                if neg:
                    return INT_MIN
                else:
                    return INT_MAX
            
            n = n*10 + c
        
        if neg:
            n = -n

        return n