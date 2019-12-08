import sys
from itertools import permutations 

all_phase_permutations = list(permutations([5, 6, 7, 8, 9]))

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

# contents = "3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5"
numstrs = contents.split(',')

orig_data = []

for string in numstrs:
    orig_data.append(int(string))

class Amplifier:
    def __str__(self):
        return f"Amplifier(phase={self.phase})"
    
    def __repr__(self):
        return self.__str__()

    def __init__(self, phase):
        self.data = list(orig_data)
        self.pos = 0
        # self.outputs = []
        self.inputs = [phase]
        self.done = False
        self.phase = phase
        self.last_output = None
    
    # def last_output(self):
        # return self.outputs[-1]
    
    def is_done(self):
        return self.done
    
    def queue_input(self, innput):
        self.inputs.append(innput)
    
    def reset_inputs(self):
        self.inputs = []
    
    def process(self, innput=None):
        if input:
            self.inputs.append(innput)

        inpos = 0
        while True:
            opcode, mode_param1, mode_param2, mode_param3 = parse_opcode(self.data[self.pos])

            if DEBUG:
                print(f"{self.data}")
                print(f"INST opcode[{opcode}] params[{mode_param1}, {mode_param2}, {mode_param3}]")

            if opcode == 99:
                self.done = True
                break

            if opcode == 1 or opcode == 2:
                param = self.data[self.pos+1]
                if DEBUG:
                    print(f"   NEXT [{param}, ", end='')
                op1 = self.data[param] if mode_param1 == 0 else param

                param = self.data[self.pos+2]
                if DEBUG:
                    print(f"{param}, ", end='')
                op2 = self.data[param] if mode_param2 == 0 else param

                param = self.data[self.pos+3]
                if DEBUG:
                    print(f"{param}]")
                # destindex = data[param] if mode_param3 == 0 else param
                destindex = param

                if opcode == 1:
                    value = op1 + op2
                elif opcode == 2:
                    value = op1 * op2

                if DEBUG:
                    print(f"   Storing {value} at {destindex}")
                self.data[destindex] = value
                
                self.pos += 4
            
            elif opcode == 3 or opcode == 4: # input, prompt for integer
                op1 = self.data[self.pos+1]
                if DEBUG:
                    print(f"   NEXT [{op1}]")

                if opcode == 3:
                    if DEBUG:
                        print("   Getting input")
                    # data[op1] = int(input("Enter Input: "))
                    if inpos > len(self.inputs)-1:
                        print("This probably shouldn't happen, no inputs left to process")
                        self.reset_inputs()
                        return None

                    self.data[op1] = self.inputs[inpos]
                    inpos += 1
                    if DEBUG:
                        print(f"   Input Recieved: {self.data[op1]}")
                elif opcode == 4:
                    value = op1 if mode_param1 == 1 else self.data[op1]
                    if DEBUG:
                        print(f"\n*** Output: {value} *** \n")
                    self.last_output = value
                    self.pos += 2
                    self.reset_inputs()
                    return value

                self.pos += 2
            
            elif opcode == 5:
                param = self.data[self.pos+1]
                if DEBUG:
                    print(f"   NEXT [{param}, ", end='')
                op1 = self.data[param] if mode_param1 == 0 else param

                param = self.data[self.pos+2]
                if DEBUG:
                    print(f"{param}]")
                op2 = self.data[param] if mode_param2 == 0 else param
            
                if op1 != 0:
                    if DEBUG:
                        print(f"   Moving current pos [{self.pos}] to [{op2}]")
                    self.pos = op2
                else:
                    self.pos += 3

            elif opcode == 6:
                param = self.data[self.pos+1]
                if DEBUG:
                    print(f"   NEXT [{param}, ", end='')
                op1 = self.data[param] if mode_param1 == 0 else param

                param = self.data[self.pos+2]
                if DEBUG:
                    print(f"{param}]")
                op2 = self.data[param] if mode_param2 == 0 else param
            
                if op1 == 0:
                    if DEBUG:
                        print(f"   Moving current pos [{self.pos}] to [{op2}]")
                    self.pos = op2
                else:
                    self.pos += 3

            elif opcode == 7:
                param = self.data[self.pos+1]
                if DEBUG:
                    print(f"   NEXT [{param}, ", end='')
                op1 = self.data[param] if mode_param1 == 0 else param

                param = self.data[self.pos+2]
                if DEBUG:
                    print(f"{param}, ", end='')
                op2 = self.data[param] if mode_param2 == 0 else param

                param = self.data[self.pos+3]
                if DEBUG:
                    print(f"{param}]")
                # op3 = data[param] if mode_param3 == 0 else param
                op3 = param
            
                if op1 < op2:
                    if DEBUG:
                        print(f"   Setting 1 to pos [{op3}]")
                    self.data[op3] = 1
                else:
                    if DEBUG:
                        print(f"   Setting 0 to pos [{op3}]")
                    self.data[op3] = 0
                
                self.pos += 4

            elif opcode == 8:
                param = self.data[self.pos+1]
                if DEBUG:
                    print(f"   NEXT [{param}, ", end='')
                op1 = self.data[param] if mode_param1 == 0 else param

                param = self.data[self.pos+2]
                if DEBUG:
                    print(f"{param}, ", end='')
                op2 = self.data[param] if mode_param2 == 0 else param

                param = self.data[self.pos+3]
                if DEBUG:
                    print(f"{param}]")
                # op3 = data[param] if mode_param3 == 0 else param
                op3 = param
            
                if op1 == op2:
                    if DEBUG:
                        print(f"   Setting 1 to pos [{op3}]")
                    self.data[op3] = 1
                else:
                    if DEBUG:
                        print(f"   Setting 0 to pos [{op3}]")
                    self.data[op3] = 0
                
                self.pos += 4

max_output = -1
for phase_permutation in all_phase_permutations:
# for phase_permutation in [[9, 8, 7, 6, 5]]:

    amplifiers = []
    for phase in phase_permutation:
        amplifiers.append(Amplifier(phase))

    amplifiers[0].queue_input(0)

    last_output = None
    while not amplifiers[-1].is_done():
        for amplifier in amplifiers:
            last_output = amplifier.process(last_output)
    
    if amplifiers[-1].last_output > max_output:
        max_output = amplifiers[-1].last_output
        
print(f"\nFinal Output: {max_output}\n")