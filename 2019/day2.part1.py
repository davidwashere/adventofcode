# input = 'day2.data.sample.txt'
input = 'day2.data.actual.txt'

with open(input, 'rt') as f:
    contents = f.read()

numstrs = contents.split(',')

data = []

for string in numstrs:
    data.append(int(string))

pos = 0
while True:
    opcode = data[pos]

    if opcode == 99:
        break

    op1index = data[pos+1]
    op1 = data[op1index]

    op2index = data[pos+2]
    op2 = data[op2index]

    destindex = data[pos+3]

    if opcode == 1:
        data[destindex] = op1 + op2
    elif opcode == 2:
        data[destindex] = op1 * op2
    
    pos += 4


print(data)
