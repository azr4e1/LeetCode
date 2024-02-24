# Given a triangle array, return the minimum path sum from top to bottom.
#
# For each step, you may move to an adjacent number of the row below. More formally, if you are on index i on the current row, you may move to either index i or index i + 1 on the next row.
#
#
#
# Example 1:
#
# Input: triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
# Output: 11
# Explanation: The triangle looks like:
#    2
#   3 4
#  6 5 7
# 4 1 8 3
# The minimum path sum from top to bottom is 2 + 3 + 5 + 1 = 11 (underlined above).
#
# Example 2:
#
# Input: triangle = [[-10]]
# Output: -10
#
#
#
# Constraints:
#
#     1 <= triangle.length <= 200
#     triangle[0].length == 1
#     triangle[i].length == triangle[i - 1].length + 1
#     -104 <= triangle[i][j] <= 104
#

from typing import List


class Solution:
    def minimumTotal(self, triangle: List[List[int]]) -> int:
        new_triangle = []
        new_triangle.append(triangle[0])

        for i in range(len(triangle)-1):
            new_triangle.append(self.sumLines(new_triangle[i], triangle[i+1]))

        return min(new_triangle[-1])

    def sumLines(self, line_prev, line_succ):
        new_line = []
        for i in range(len(line_prev)):
            if i == 0:
                new_line.append(line_prev[i] + line_succ[i])
            else:
                new_el = line_prev[i] + line_succ[i]
                if new_el < new_line[i]:
                    new_line[i] = new_el

            new_line.append(line_prev[i] + line_succ[i+1])
        return new_line


sol = Solution()
triangle = [[2], [3, 4], [6, 5, 7], [4, 1, 8, 3]]
print(sol.minimumTotal(triangle))
