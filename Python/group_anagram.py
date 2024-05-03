# Given an array of strings strs, group the anagrams together. You can return the
# answer in any order.
#
# An Anagram is a word or phrase formed by rearranging the letters of a different
# word or phrase, typically using all the original letters exactly once.
#
#
#
# Example 1:
#
# Input: strs = ["eat","tea","tan","ate","nat","bat"]
# Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
#
# Example 2:
#
# Input: strs = [""]
# Output: [[""]]
#
# Example 3:
#
# Input: strs = ["a"]
# Output: [["a"]]
from typing import List
from collections import defaultdict


class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        groups = defaultdict(list)
        for ana in strs:
            norm = str(sorted(ana))
            groups[norm].append(ana)

        return list(groups.values())
