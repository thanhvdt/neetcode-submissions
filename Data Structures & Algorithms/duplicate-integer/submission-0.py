class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        numDict = {}
        for num in nums:
            if num not in numDict.keys():
                numDict[num] = 1
                continue
            return True 
        return False
