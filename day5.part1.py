import sys

def parse_opcode(opcode):
    """
    opcode : int

    Returns
    --------------
    [opcode, mode_param1, mode_param2, mode_param3]

    A mode_param of 0 == get value from index represented by parameter
    A mode_param of 1 == the parameter is the literal value
    """

    if opcode < 100:
        return [opcode, 0, 0, 0]
    
    opcode = str(opcode)
    remaining = opcode[:-2]
    opcode = int(opcode[-2:])
    
    mode_param1 = int(remaining[-1])
    if len(remaining) <= 1:
        return [int(opcode), mode_param1, 0, 0]

    remaining = remaining[:-1]
    mode_param2 = int(remaining[-1])
    if len(remaining) <= 1:
        return [int(opcode), mode_param1, mode_param2, 0]

    remaining = remaining[:-1]
    mode_param2 = int(remaining[-1])
    return [int(opcode), mode_param1, mode_param2, mode_param2]
    

# print(parse_opcode(1002))
# print(parse_opcode(1101))
# sys.exit(0)

filename = 'day5.data.actual.txt'
# filename = 'day2.data.sample.txt'

with open(filename, 'rt') as f:
    contents = f.read()

numstrs = contents.split(',')

data = []

for string in numstrs:
    data.append(int(string))

pos = 0
while True:
    opcode, mode_param1, mode_param2, mode_param3 = parse_opcode(data[pos])

    print(f"INST opcode[{opcode}] params[{mode_param1}, {mode_param2}, {mode_param3}] NEXT data[{data[pos+1]}, {data[pos+2]}, {data[pos+3]}]")

    if opcode == 99:
        break

    if opcode == 1 or opcode == 2:
        param = data[pos+1]
        op1 = data[param] if mode_param1 == 0 else param

        param = data[pos+2]
        op2 = data[param] if mode_param2 == 0 else param

        param = data[pos+3]
        # destindex = data[param] if mode_param3 == 0 else param
        destindex = param

        if opcode == 1:
            data[destindex] = op1 + op2
        elif opcode == 2:
            data[destindex] = op1 * op2
        
        pos += 4
    
    elif opcode == 3 or opcode == 4: # input, prompt for integer
        op1 = data[pos+1]

        if opcode == 3:
            print("Getting input = 1")
            data[op1] = 1
        elif opcode == 4:
            print(f"Output: {data[op1]}")

        pos += 2
    


