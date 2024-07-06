class Solution:

    def encode(self, strs: List[str]) -> str:
        encoding = ""
        for el in strs:
            encoding += el + "\0"

        return encoding

    def decode(self, s: str) -> List[str]:
        return s.split("\0")[:-1]
