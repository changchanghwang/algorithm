class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        newSet = set()
        for num in nums:
            if num in newSet:
                return True
            newSet.add(num)
        return False