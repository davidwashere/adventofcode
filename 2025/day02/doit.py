import bisect, itertools

maxdig = 10 # max no. digits in input
maxnum = 10 ** (maxdig // 2)

def solve(p: bool) -> int:
    tab = sorted({ int(j * str(i))
                   for i in range(1, maxnum)
                   for j in range(2, maxdig // len(str(i)) + 1 if p else 3) })
    dp = (0, *itertools.accumulate(tab))
    return sum(dp[bisect.bisect_right(tab, b)] - dp[bisect.bisect_left(tab, a)]
               for r in inp for a, b in [map(int, r.split("-"))])

inp = open(0).read().split(",")
print(f"silver: {solve(0)}, gold: {solve(1)}")