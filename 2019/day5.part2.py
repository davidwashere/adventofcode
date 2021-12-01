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

# contents = "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99"
numstrs = contents.split(',')

data = []

for string in numstrs:
    data.append(int(string))

pos = 0
while True:
    opcode, mode_param1, mode_param2, mode_param3 = parse_opcode(data[pos])

    print(f"{data}")
    print(f"INST opcode[{opcode}] params[{mode_param1}, {mode_param2}, {mode_param3}]")

    if opcode == 99:
        break

    if opcode == 1 or opcode == 2:
        param = data[pos+1]
        print(f"   NEXT [{param}, ", end='')
        op1 = data[param] if mode_param1 == 0 else param

        param = data[pos+2]
        print(f"{param}, ", end='')
        op2 = data[param] if mode_param2 == 0 else param

        param = data[pos+3]
        print(f"{param}]")
        # destindex = data[param] if mode_param3 == 0 else param
        destindex = param

        if opcode == 1:
            value = op1 + op2
        elif opcode == 2:
            value = op1 * op2

        print(f"   Storing {value} at {destindex}")
        data[destindex] = value
        
        pos += 4
    
    elif opcode == 3 or opcode == 4: # input, prompt for integer
        op1 = data[pos+1]
        print(f"   NEXT [{op1}]")

        if opcode == 3:
            print("Getting input")
            data[op1] = int(input("Enter Input: "))
        elif opcode == 4:
            value = op1 if mode_param1 == 1 else data[op1]
            print(f"\n*** Output: {value} *** \n")

        pos += 2
    
    elif opcode == 5:
        param = data[pos+1]
        print(f"   NEXT [{param}, ", end='')
        op1 = data[param] if mode_param1 == 0 else param

        param = data[pos+2]
        print(f"{param}]")
        op2 = data[param] if mode_param2 == 0 else param
    
        if op1 != 0:
            print(f"   Moving current pos [{pos}] to [{op2}]")
            pos = op2
        else:
            pos += 3

    elif opcode == 6:
        param = data[pos+1]
        print(f"   NEXT [{param}, ", end='')
        op1 = data[param] if mode_param1 == 0 else param

        param = data[pos+2]
        print(f"{param}]")
        op2 = data[param] if mode_param2 == 0 else param
    
        if op1 == 0:
            print(f"   Moving current pos [{pos}] to [{op2}]")
            pos = op2
        else:
            pos += 3

    elif opcode == 7:
        param = data[pos+1]
        print(f"   NEXT [{param}, ", end='')
        op1 = data[param] if mode_param1 == 0 else param

        param = data[pos+2]
        print(f"{param}, ", end='')
        op2 = data[param] if mode_param2 == 0 else param

        param = data[pos+3]
        print(f"{param}]")
        # op3 = data[param] if mode_param3 == 0 else param
        op3 = param
    
        if op1 < op2:
            print(f"   Setting 1 to pos [{op3}]")
            data[op3] = 1
        else:
            print(f"   Setting 0 to pos [{op3}]")
            data[op3] = 0
        
        pos += 4

    elif opcode == 8:
        param = data[pos+1]
        print(f"   NEXT [{param}, ", end='')
        op1 = data[param] if mode_param1 == 0 else param

        param = data[pos+2]
        print(f"{param}, ", end='')
        op2 = data[param] if mode_param2 == 0 else param

        param = data[pos+3]
        print(f"{param}]")
        # op3 = data[param] if mode_param3 == 0 else param
        op3 = param
    
        if op1 == op2:
            print(f"   Setting 1 to pos [{op3}]")
            data[op3] = 1
        else:
            print(f"   Setting 0 to pos [{op3}]")
            data[op3] = 0
        
        pos += 4