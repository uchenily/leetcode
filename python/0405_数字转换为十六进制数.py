class Solution:
    def toHex(self, num: int) -> str:
        if num == 0:
            return "0"
        result = []
        base = "0123456789abcdef"
        if num < 0:
            # num += 0xffffffff + 0x00000001
            num += 2 ** 32
        while num > 0:
            # result.append(base[num % 16])
            result.append(base[num & 0xf])
            # num //= 16
            num >>= 4
        result.reverse()
        return ''.join(result)
