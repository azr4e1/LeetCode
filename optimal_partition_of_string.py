# Given a string s, partition the string into one or more substrings such that the characters
# in each substring are unique. That is, no letter appears in a single substring more than once.
#
# Return the minimum number of substrings in such a partition.
#
# Note that each character should belong to exactly one substring in a partition.
#
#
#
# Example 1:
#
# Input: s = "abacaba"
# Output: 4
# Explanation:
# Two possible partitions are ("a","ba","cab","a") and ("ab","a","ca","ba").
# It can be shown that 4 is the minimum number of substrings needed.
#
# Example 2:
#
# Input: s = "ssssss"
# Output: 6
# Explanation:
# The only valid partition is ("s","s","s","s","s","s").
#
#
#
# Constraints:
#
#     1 <= s.length <= 105
#     s consists of only English lowercase letters.
#
#
class Solution:
    def partitionString_slow(self, s: str) -> int:

        partition = 0
        counter = {}
        for letter in s:
            if counter.get(letter, 0) > 0:
                counter = {}
                partition += 1
                counter[letter] = 1
            else:
                counter[letter] = 1

        return partition+1

    def partitionString(self, s: str) -> int:
        out = 1
        sub = ""
        for c in s:
            if c in sub:
                out += 1
                sub = c
            else:
                sub += c
        return out


s = "abacaba"
sol = Solution()
print(sol.partitionString(s))

s = "hdklqkcssgxlvehva"
print(sol.partitionString(s))
