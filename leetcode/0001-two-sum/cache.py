#
# @lc app=leetcode.cn id=1 lang=python
#
# [1] 两数之和
#

# @lc code=start
class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        cache = {} # value, index
        for i, num in enumerate(nums):
            if target - num in cache:
                return [cache[target - num], i]
            else:
                cache[num] = i
        return []

# @lc code=end

