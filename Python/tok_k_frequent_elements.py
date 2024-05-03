# Given an integer array nums and an integer k, return the k most frequent
# elements. You may return the answer in any order.
#
#
# Example 1:
#
# Input: nums = [1, 1, 1, 2, 2, 3], k = 2
# Output: [1, 2]
#
# Example 2:
#
# Input: nums = [1], k = 1
# Output: [1]
from typing import List
from collections import defaultdict


class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        if len(nums) <= k:
            return nums

        cache = defaultdict(int)

        for el in nums:
            cache[el] += 1

        sorted_els = sorted(cache.items(), key=lambda x: x[1], reverse=True)
        return [sorted_els[i][0] for i in range(k)]

        # cache_freq = defaultdict(int)
        # cache_index = dict()
        # elems = []
        #
        # for i in range(k):
        #     el = nums[i]
        #     elems.append(nums[i])
        #     cache_index[el] = i
        #     cache_freq[el] += 1
        #
        # for num in nums[k:]:
        #     cache_freq[num] += 1
