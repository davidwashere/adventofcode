import math

# data = 'day1.data.sample.txt'
data = 'day1.data.actual.txt'
with open(data, 'rt') as f:
    contents = f.read()

lines = contents.splitlines()

sum = 0
for line in lines:
    num = int(line)
    by3 = num / 3
    down = math.floor(by3)
    sub2 = down - 2
    print( f"{line} - by3 {by3} - round down {down} - minus2 {sub2}" )

    sum += sub2

print(f"\nTotal: {sum}")