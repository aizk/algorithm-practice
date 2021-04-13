import heapq
from typing import List


class Solution:
    def getLeastNumbers(self, arr: List[int], k: int) -> List[int]:
        if k == 0:
            return []
        # hp 就是堆
        hp = [-x for x in arr[:k]]
        heapq.heapify(hp)
        for i in range(k, len(arr)):
            # 最小的负数就是最大的整数，如果比当前元素大，就要被替换掉，因此要找最小的 k 个数
            if -hp[0] > arr[i]:
                heapq.heappop(hp)
                heapq.heappush(hp, -arr[i])
        return [-x for x in hp]

