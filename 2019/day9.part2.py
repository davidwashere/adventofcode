import sys

DEBUG = True

def parse_opcode(opcode):
    """
    opcode : int

    Returns
    --------------
    [opcode, mode_param1, mode_param2, mode_param3]

    A mode_param of 0 == get value from index represented by parameter
    A mode_param of 1 == the parameter is the literal value
    """

    if DEBUG:
        print(f"\nParsing opcode: {opcode}")

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
    mode_param3 = int(remaining[-1])
    return [int(opcode), mode_param1, mode_param2, mode_param3]
    

class Amplifier:
    def __init__(self):
        self.data = list(orig_data)
        self.pos = 0
        self.done = False
        self.last_output = None
        self.relbase = 0
        self.inputs = []
        self.extra_data = {}
    
    def is_done(self):
        return self.done
    
    def queue_input(self, innput):
        self.inputs.append(innput)
    
    def reset_inputs(self):
        self.inputs = []
    
    def process(self, innput=None):
        if innput:
            self.inputs.append(innput)

        inpos = 0
        while True:
            opcode, mode_param1, mode_param2, mode_param3 = parse_opcode(self.data[self.pos])

            if DEBUG:
                print(f"{self.data[:40]}")
                print(f"INST opcode[{opcode}] params[{mode_param1}, {mode_param2}, {mode_param3}]")
                print(f"   POS: {self.pos}")

            if opcode == 99:
                self.done = True
                break

            if opcode == 1 or opcode == 2: # opcode 1 = add 1st and 2nd together, store at 3rd
                                           # opcode 2 = multiple instead
                param = self.data[self.pos+1]
                if DEBUG:
                    print(f"   NEXT [{param}, ", end='')
                op1 = self._get_op_value(mode_param1, param)

                param = self.data[self.pos+2]
                if DEBUG:
                    print(f"{param}, ", end='')
                op2 = self._get_op_value(mode_param2, param)

                param = self.data[self.pos+3]
                if DEBUG:
                    print(f"{param}]")
                destindex = self._get_op_index(mode_param3, param)

                if opcode == 1:
                    value = op1 + op2
                elif opcode == 2:
                    value = op1 * op2

                if DEBUG:
                    print(f"   Storing {value} at {destindex}")
                self._store(destindex, value)
                
                self.pos += 4
            
            elif opcode == 3 or opcode == 4: 
                op1 = self.data[self.pos+1]
                if DEBUG:
                    print(f"   NEXT [{op1}]")

                if opcode == 3: # Grab the input from buffer and store at 1st param
                    if DEBUG:
                        print("   Getting input")
                    if inpos > len(self.inputs)-1:
                        print("This probably shouldn't happen, no inputs left to process")
                        self.reset_inputs()
                        return None

                    indata = self.inputs[inpos]
                    destindex = self._get_op_index(mode_param1, op1)

                    if DEBUG:
                        print(f"   Input Recieved: {indata}")
                        print(f"   Storing {indata} at: {destindex}")
                        print(f"   Relbase: {self.relbase}")

                    self._store(destindex, indata)
                    inpos += 1
                    self.pos += 2
                elif opcode == 4: # output the value based on 1st param
                    value = self._get_op_value(mode_param1, op1)
                    print(f"\n*** Output: {value} *** \n")
                    self.last_output = value
                    self.pos += 2
                    self.reset_inputs()
                    # return value
            
            elif opcode == 5: # Jump when 1st param non-zero to 2nd param
                param = self.data[self.pos+1]
                if DEBUG:
                    print(f"   NEXT [{param}, ", end='')
                op1 = self._get_op_value(mode_param1, param)

                param = self.data[self.pos+2]
                if DEBUG:
                    print(f"{param}]")
                op2 = self._get_op_value(mode_param2, param)
            
                if op1 != 0:
                    if DEBUG:
                        print(f"   Moving current pos [{self.pos}] to [{op2}]")
                    self.pos = op2
                else:
                    self.pos += 3

            elif opcode == 6: # Jump when 1st param is zero to 2nd param
                param = self.data[self.pos+1]
                if DEBUG:
                    print(f"   NEXT [{param}, ", end='')
                op1 = self._get_op_value(mode_param1, param)

                param = self.data[self.pos+2]
                if DEBUG:
                    print(f"{param}]")
                op2 = self._get_op_value(mode_param2, param)
            
                if op1 == 0:
                    if DEBUG:
                        print(f"   Moving current pos [{self.pos}] to [{op2}]")
                    self.pos = op2
                else:
                    self.pos += 3

            elif opcode == 7: # if 1st < 2nd store 1 in 3rd, else store 0 in 3rd
                param = self.data[self.pos+1]
                if DEBUG:
                    print(f"   NEXT [{param}, ", end='')
                op1 = self._get_op_value(mode_param1, param)

                param = self.data[self.pos+2]
                if DEBUG:
                    print(f"{param}, ", end='')
                op2 = self._get_op_value(mode_param2, param)

                param = self.data[self.pos+3]
                if DEBUG:
                    print(f"{param}]")
                destindex = self._get_op_index(mode_param3, param)
            
                if op1 < op2:
                    if DEBUG:
                        print(f"   Setting 1 to pos [{destindex}]")
                    self._store(destindex, 1)
                else:
                    if DEBUG:
                        print(f"   Setting 0 to pos [{destindex}]")
                    self._store(destindex, 0)
                
                self.pos += 4

            elif opcode == 8: # if 1st == 2nd store 1 in 3rd, else store 0 in 3rd
                param = self.data[self.pos+1]
                if DEBUG:
                    print(f"   NEXT [{param}, ", end='')
                op1 = self._get_op_value(mode_param1, param)

                param = self.data[self.pos+2]
                if DEBUG:
                    print(f"{param}, ", end='')
                op2 = self._get_op_value(mode_param2, param)

                param = self.data[self.pos+3]
                if DEBUG:
                    print(f"{param}]")
                op3 = self._get_op_index(mode_param3, param)
            
                if op1 == op2:
                    if DEBUG:
                        print(f"   Setting 1 to pos [{op3}]")
                    self._store(op3, 1)
                else:
                    if DEBUG:
                        print(f"   Setting 0 to pos [{op3}]")
                    self._store(op3, 0)
                
                self.pos += 4

            elif opcode == 9: # Move releative base by 1st param
                param = self.data[self.pos+1]

                currelbase = self.relbase
                self.relbase += self._get_op_value(mode_param1, param)
                if DEBUG:
                    print(f"   NEXT [{param}]")
                    print(f"   Moving RELBASE from {currelbase} to {self.relbase}")

                self.pos += 2
            
            else:
                print(f"Error: Unknown opcode {opcode} at pos {self.pos}")
                raise Exception("")
        
    def _get_op_index(self, mode, param):
        try:
            if mode == 0:
                return param
            
            if mode == 2:
                return self.relbase + param

            raise Exception(f"Cannot get index, sumthing wrong: mode={mode}, param={param}")


        except Exception as err:
            print(f"Exception[A]: mode={mode}, param={param}, len={len(self.data)}")
            raise(err)
            
    def _get_op_value(self, mode, param):
        try:
            if mode == 0:
                return self._get(param)

            if mode == 1:
                return param
            
            if mode == 2:
                return self._get(self.relbase + param)

        except Exception as err:
            print(f"Exception[A]: mode={mode}, param={param}, len={len(self.data)}")
            raise(err)
    
    def _store(self, destindex, value):
        if destindex < len(self.data):
            self.data[destindex] = value
        
        self.extra_data[destindex] = value
    
    def _get(self, index):
        if index < len(self.data):
            return self.data[index]
        
        if index in self.extra_data:
            return self.extra_data[index]

        return 0

        

filename = 'day9.data.actual.txt'

with open(filename, 'rt') as f:
    contents = f.read()

# DEBUG = True
DEBUG = False
# contents = "109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99"
# contents = "1102,34915192,34915192,7,4,7,99,0"
# contents = "104,1125899906842624,99"
numstrs = contents.split(',')

orig_data = []

for string in numstrs:
    orig_data.append(int(string))

# for _ in range(0, 300000):
#     orig_data.append(0)

prog = Amplifier()
prog.queue_input(2)
prog.process()
