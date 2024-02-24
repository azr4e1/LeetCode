# You have a long flowerbed in which some of the plots are planted, and some are not. However, flowers cannot be planted in adjacent plots.
#
# Given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1 means not empty, and an integer n, return true if n new flowers can be planted in the flowerbed without violating the no-adjacent-flowers rule and false otherwise.
#
#  
#
# Example 1:
#
# Input: flowerbed = [1,0,0,0,1], n = 1
# Output: true
#
# Example 2:
#
# Input: flowerbed = [1,0,0,0,1], n = 2
# Output: false
#
#  
#
# Constraints:
#
#     1 <= flowerbed.length <= 2 * 104
#     flowerbed[i] is 0 or 1.
#     There are no two adjacent flowers in flowerbed.
#     0 <= n <= flowerbed.length
#
#

class Solution:
    def checkAround(self, flowerbed, i):
        if flowerbed[i-1] == 0 and flowerbed[i] == 0 and flowerbed[i+1] == 0:
            return True
        else:
            return False

    def canPlaceFlowers_slow(self, flowerbed: list[int], n: int) -> bool:
        length = 0
        flowerbed_length = len(flowerbed)

        if flowerbed_length == 1 and flowerbed[0] == 0 and n == 1:
            return True
        elif flowerbed_length == 1 and n == 0:
            return True
        elif flowerbed_length == 1:
            return False
        else:
            for i in range(0, len(flowerbed)):
                if i == 0:
                    if flowerbed[1] == 0 and flowerbed[0] == 0:
                        flowerbed[0] = 1
                        length += 1
                elif i == flowerbed_length-1:
                    if flowerbed[i-1] == 0 and flowerbed[i] == 0:
                        flowerbed[i] = 1
                        length += 1
                else:
                    if self.checkAround(flowerbed, i):
                        flowerbed[i] = 1
                        length += 1

                if length == n:
                    break

            print(flowerbed)
            if length >= n:
                return True
            else:
                return False

    def canPlaceFlowers(self, flowerbed: list[int], n: int) -> bool:
        # for i in range(len(flowerbed)):
        #     if n==0:
        #         return True
        #     elif i==0 and flowerbed[i+1]==0:
        #         flowebed[i]=1
        #         n-=1
        #     elif flowerbed[i-1]==0 and flowerbed[i+1]==0:
        #         flowerbed[i]=1
        #         n-=1
        #     elif i==len(flowerbed)-1 and flowerbed[i-1]==0:
        #         flowerbed[i-1]=1
        #         n-=1
        c = 1
        f = 0
        for i in flowerbed:
            if i:
                c = 0
            else:
                c += 1
                if c == 3:
                    f += 1
                    c = 1
        if not flowerbed[-1]:
            c += 1
            if c == 3:
                f += 1
        return f >= n


sol = Solution()

flowerbed = [1,0,0,0,1]
n = 1
print(sol.canPlaceFlowers(flowerbed, n))

flowerbed = [1,0,0,0,1]
n = 2
print(sol.canPlaceFlowers(flowerbed, n))

flowerbed = [1,0,0,0,1,0,0]
n = 2
print(sol.canPlaceFlowers(flowerbed, n))

flowerbed = [1]
n = 0
print(sol.canPlaceFlowers(flowerbed, n))

flowerbed = [0,0,0,0,1]
n = 2
print(sol.canPlaceFlowers(flowerbed, n))

flowerbed = [0,1,0]
n = 1
print(sol.canPlaceFlowers(flowerbed, n))
