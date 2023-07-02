# You are given the root of a binary search tree (BST) and an integer val.
#
# Find the node in the BST that the node's value equals val and return the subtree rooted with that node. If such a node does not exist, return null.
#
# Constraints:
#
#     The number of nodes in the tree is in the range [1, 5000].
#     1 <= Node.val <= 107
#     root is a binary search tree.
#     1 <= val <= 107
#
from typing import Optional


# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class BST:
    def __init__(self, linked_list: list) -> None:
        pass


class Solution:
    def searchBST_linkedLists(self, root: Optional[TreeNode], val: int):
        if val not in root:
            return []
        partial_list = root
        root_el = root[0]
        while root_el != val:
            if val >= root_el:
                order = "greater"
            else:
                order = "less"

            partial_list = self.partial_search(root_el, partial_list, order)

            if partial_list:
                root_el = partial_list[0]

        # print("Partial list:", partial_list)
        smaller_el = self.partial_search(val, partial_list, "less")[0]
        greater_el = self.partial_search(val, partial_list, "greater")[0]

        return [val, smaller_el, greater_el]

    def partial_search(self, root, partial_list, order):
        if order == "greater":
            partial_new_list = [el for el in partial_list[1:] if el >= root]
        else:
            partial_new_list = [el for el in partial_list[1:] if el < root]

        return partial_new_list

    def searchBST(self, root: Optional[TreeNode], val: int):
        while root.val != val:
            if val >= root.val:
                root = root.right
            else:
                root = root.left

            if root is None:
                break
        else:
            # return TreeNode(val=root.val, left=root.left.val, right=root.right.val)
            return root

        return None



sol = Solution()

root = TreeNode(val= 4, left= TreeNode(val= 2, left= TreeNode(val= 1, left= None, right= None), right= TreeNode(val= 3, left= None, right= None)), right= TreeNode(val= 7, left= None, right= None))
val = 2
result = sol.searchBST(root, val)
print(result.val, result.left, result.right)

# root = [4, 2, 7, 1, 3]
# val = 5
# print(sol.searchBST(root, val))
