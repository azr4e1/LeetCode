# Given an integer array nums, return an array output where output[i] is the product of all the elements of nums except nums[i].
#
# Each product is guaranteed to fit in a 32-bit integer.
#
# Follow-up: Could you solve it in O(n)O(n) time without using the division operation?
from typing import List


class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        cumsum = []
        cumsum_reverse = []
        for i in range(len(nums)):
            if len(cumsum) == 0:
                cumsum.append(nums[i])
            else:
                cumsum.append(cumsum[i-1] * nums[i])
            if len(cumsum_reverse) == 0:
                cumsum_reverse.append(nums[-1 - i])
            else:
                cumsum_reverse.append(cumsum_reverse[i-1] * nums[-1-i])

        final = [cumsum_reverse[-2]]
        for i in range(1, len(nums[:-1])):
            final.append(cumsum[i-1] * cumsum_reverse[-2-i])

        final.append(cumsum[-2])

        return final
