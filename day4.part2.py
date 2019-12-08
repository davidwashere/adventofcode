low = 367479
high = 893698
    
def has_small_double(num): 
    numstr = str(num)

    # digit > occurances
    occurances = [0] * 10

    for digit in numstr:
        digit = int(digit)
        occurances[digit] += 1

    for item in occurances:
        if item == 2:
            return True

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
    if has_small_double(num) and never_decreases(num):
        return True

    return False

# print(has_small_double(112233))
# print(has_small_double(123444))
# print(has_small_double(111122))

count = 0
for num in range(low, high):
    if is_valid(num):
        count += 1
        print(f'{num} is valid!')

print(f'\nFinal Count of Valid Entries: {count}')

