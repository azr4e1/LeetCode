# Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using the letters from magazine and false otherwise.
#
# Each letter in magazine can only be used once in ransomNote.
#
#
#
# Example 1:
#
# Input: ransomNote = "a", magazine = "b"
# Output: false
#
# Example 2:
#
# Input: ransomNote = "aa", magazine = "ab"
# Output: false
#
# Example 3:
#
# Input: ransomNote = "aa", magazine = "aab"
# Output: true
#
#
#
# Constraints:
#
#     1 <= ransomNote.length, magazine.length <= 105
#     ransomNote and magazine consist of lowercase English letters.
#
#
from collections import defaultdict


class Solution:
    def canConstruct_slow(self, ransomNote: str, magazine: str) -> bool:

        result = True

        letters = defaultdict(int)
        for a in magazine:
            letters[a] += 1

        for a in ransomNote:
            if a in letters and letters[a] > 0:
                letters[a] -= 1
            else:
                result = False
                break

        return result

    def canConstruct(self, ransomNote: str, magazine: str) -> bool:
        rcounter = {}
        for l in ransomNote:
            rcounter[l] = ransomNote.count(l)
        
        for k in rcounter:
            if magazine.count(k) < rcounter[k]:
                return False
        
        return True
