low = 367479
high = 893698
    
def has_double(num): 
    numstr = str(num)

    prev = 'x'
    for digit in numstr:
        if digit == prev:
            return True
        prev = digit
    
    return False

def never_decreases(num):
    numstr = str(num)

   
    prev = -1
    for digit in numstr:
        digit = int(digit)
        if digit < prev:
            return False
        prev = digit

    
    return True

def is_valid(num):
    if has_double(num) and never_decreases(num):
        return True

    return False

count = 0
for num in range(low, high):
    if is_valid(num):
        count += 1
        print(f'{num} is valid!')

print(f'\nFinal Count of Valid Entries: {count}')

