class Solution(object):
    def getLeastNumbers(self, arr, k):
        """
        :type arr: List[int]
        :type k: int
        :rtype: List[int]
        """
        if len(arr) == k:
            return arr

        s = 0
        e = len(arr) - 1
        while True:
            j = self.partition(arr, s, e)
            if j == k:
                return arr[:j]
            elif j < k:
                s = j + 1
            elif j > k:
                e = j - 1

    def partition(self, arr, s, e):
        v = arr[e]
        i = s - 1
        for j in range(s, e):
            if arr[j] < v:
                i += 1
                # exchange
                t = arr[j]
                arr[j] = arr[i]
                arr[i] = t
        t = arr[e]
        arr[e] = arr[i + 1]
        arr[i + 1] = t
        return i + 1


if __name__ == '__main__':
    s = Solution()
    print(s.getLeastNumbers([3, 2, 1], 2))
    print(s.getLeastNumbers([0, 0, 2, 3, 2, 1, 1, 2, 0, 4], 10))
