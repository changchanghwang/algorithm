class Solution:
   def encode(self, strs: List[str]) -> str:
        text = ""
        for str in strs:
            text += f"{len(str)}:{str}"
        return text

    def decode(self, s: str) -> List[str]:
        ls, start = [], 0
        while start < len(s):
            mid = s.find(":", start)
            length = int(s[start:mid])
            ls.append(s[mid + 1 : mid + 1 + length])
            start = mid + 1 + length
        return ls