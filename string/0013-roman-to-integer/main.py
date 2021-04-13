class Solution(object):
    def romanToInt(self, s):
        m = {
            'I': 1,
            'V': 5,
            'X': 10,
            'L': 50,
            'C': 100,
            'D': 500,
            'M': 1000,
        }
        total = 0
        for i in range(len(s) - 1):
            v = m[s[i]]
            if v < m[s[i+1]]:
                total -= v
            else:
                total += v
        # 最后一个一定是加
        total += m[s[-1]]
        return total


if __name__ == '__main__':
    s = Solution()
    print(s.romanToInt('MCMXCIV'))
    print(s.romanToInt('VII'))
    print(s.romanToInt('IV'))