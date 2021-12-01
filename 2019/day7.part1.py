import sys
from itertools import permutations 

all_phase_permutations = list(permutations([0, 1, 2, 3, 4]))

DEBUG = False

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
    

filename = 'day7.data.actual.txt'
# filename = 'day7.data.sample.txt'

with open(filename, 'rt') as f:
    contents = f.read()

# contents = "3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0"
numstrs = contents.split(',')

data = []

for string in numstrs:
    data.append(int(string))

orig_data = list(data)

def intcode_execute(phase, innput):

    inputs = [phase, innput]

    inpos = 0
    pos = 0
    last_output = None
    while True:
        opcode, mode_param1, mode_param2, mode_param3 = parse_opcode(data[pos])

        if DEBUG:
            print(f"{data}")
            print(f"INST opcode[{opcode}] params[{mode_param1}, {mode_param2}, {mode_param3}]")

        if opcode == 99:
            return last_output

        if opcode == 1 or opcode == 2:
            param = data[pos+1]
            if DEBUG:
                print(f"   NEXT [{param}, ", end='')
            op1 = data[param] if mode_param1 == 0 else param

            param = data[pos+2]
            if DEBUG:
                print(f"{param}, ", end='')
            op2 = data[param] if mode_param2 == 0 else param

            param = data[pos+3]
            if DEBUG:
                print(f"{param}]")
            destindex = data[param] if mode_param3 == 0 else param
            destindex = param

            if opcode == 1:
                value = op1 + op2
            elif opcode == 2:
                value = op1 * op2

            if DEBUG:
                print(f"   Storing {value} at {destindex}")
            data[destindex] = value
            
            pos += 4
        
        elif opcode == 3 or opcode == 4: # input, prompt for integer
            op1 = data[pos+1]
            if DEBUG:
                print(f"   NEXT [{op1}]")

            if opcode == 3:
                if DEBUG:
                    print("   Getting input")
                # data[op1] = int(input("Enter Input: "))
                data[op1] = inputs[inpos]
                inpos += 1
                if DEBUG:
                    print(f"   Input Recieved: {data[op1]}")
            elif opcode == 4:
                value = op1 if mode_param1 == 1 else data[op1]
                if DEBUG:
                    print(f"\n*** Output: {value} *** \n")
                last_output = value

            pos += 2
        
        elif opcode == 5:
            param = data[pos+1]
            if DEBUG:
                print(f"   NEXT [{param}, ", end='')
            op1 = data[param] if mode_param1 == 0 else param

            param = data[pos+2]
            if DEBUG:
                print(f"{param}]")
            op2 = data[param] if mode_param2 == 0 else param
        
            if op1 != 0:
                if DEBUG:
                    print(f"   Moving current pos [{pos}] to [{op2}]")
                pos = op2
            else:
                pos += 3

        elif opcode == 6:
            param = data[pos+1]
            if DEBUG:
                print(f"   NEXT [{param}, ", end='')
            op1 = data[param] if mode_param1 == 0 else param

            param = data[pos+2]
            if DEBUG:
                print(f"{param}]")
            op2 = data[param] if mode_param2 == 0 else param
        
            if op1 == 0:
                if DEBUG:
                    print(f"   Moving current pos [{pos}] to [{op2}]")
                pos = op2
            else:
                pos += 3

        elif opcode == 7:
            param = data[pos+1]
            if DEBUG:
                print(f"   NEXT [{param}, ", end='')
            op1 = data[param] if mode_param1 == 0 else param

            param = data[pos+2]
            if DEBUG:
                print(f"{param}, ", end='')
            op2 = data[param] if mode_param2 == 0 else param

            param = data[pos+3]
            if DEBUG:
                print(f"{param}]")
            # op3 = data[param] if mode_param3 == 0 else param
            op3 = param
        
            if op1 < op2:
                if DEBUG:
                    print(f"   Setting 1 to pos [{op3}]")
                data[op3] = 1
            else:
                if DEBUG:
                    print(f"   Setting 0 to pos [{op3}]")
                data[op3] = 0
            
            pos += 4

        elif opcode == 8:
            param = data[pos+1]
            if DEBUG:
                print(f"   NEXT [{param}, ", end='')
            op1 = data[param] if mode_param1 == 0 else param

            param = data[pos+2]
            if DEBUG:
                print(f"{param}, ", end='')
            op2 = data[param] if mode_param2 == 0 else param

            param = data[pos+3]
            if DEBUG:
                print(f"{param}]")
            # op3 = data[param] if mode_param3 == 0 else param
            op3 = param
        
            if op1 == op2:
                if DEBUG:
                    print(f"   Setting 1 to pos [{op3}]")
                data[op3] = 1
            else:
                if DEBUG:
                    print(f"   Setting 0 to pos [{op3}]")
                data[op3] = 0
            
            pos += 4

max_output = -1
for phase_permutation in all_phase_permutations:
    output = 0
    for phase in phase_permutation:
        output = intcode_execute(phase, output)
    
    if output > max_output:
        max_output = output

print(f"Final Output: {max_output}")