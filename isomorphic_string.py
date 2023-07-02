# Given two strings s and t, determine if they are isomorphic.
#
# Two strings s and t are isomorphic if the characters in s can be replaced to get t.
#
# All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character, but a character may map to itself.
#
#
#
# Example 1:
#
# Input: s = "egg", t = "add"
# Output: true
#
# Example 2:
#
# Input: s = "foo", t = "bar"
# Output: false
#
# Example 3:
#
# Input: s = "paper", t = "title"
# Output: true
#
#
#
# Constraints:
#
#     1 <= s.length <= 5 * 104
#     t.length == s.length
#     s and t consist of any valid ascii character.
#
#
from collections import defaultdict


class Solution:
    def isIsomorphic_slow(self, s: str, t: str) -> bool:
        s_t = defaultdict(set)
        t_s = defaultdict(set)
        for i, j in zip(s, t):
            s_t[i].add(j)
            t_s[j].add(i)

        s_t_max = max(map(len, s_t.values()))
        t_s_max = max(map(len, t_s.values()))

        result = False
        if s_t_max * t_s_max == 1:
            result = True

        return result

    def isIsomorphic(self, s: str, t: str) -> bool:
        return [*map(s.index, s)] == [*map(t.index, t)]

    def isIsomorphic2(self, s: str, t: str) -> bool:
        return len(set(s)) == len(set(t)) == len(set(zip(s, t)))


sol = Solution()
s = "egg"
t = "add"
print(sol.isIsomorphic(s, t))

s = "foo"
t = "bar"
print(sol.isIsomorphic(s, t))

s = "paper"
t = "title"
print(sol.isIsomorphic(s, t))

s = "badc"
t = "baba"
print(sol.isIsomorphic(s, t))

s = "abab"
t = "baba"
print(sol.isIsomorphic(s, t))
