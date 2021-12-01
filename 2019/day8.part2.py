import sys


filename = 'day8.data.actual.txt'
# filename = 'day7.data.sample.txt'

with open(filename, 'rt') as f:
    contents = f.read()

width = 25
height = 6

size = width * height

layers = []
layer = []
count = 0
for char_digit in contents:
    if count % (size) == 0 and count > 0:
        layers.append(layer)
        layer = []

    layer.append(int(char_digit))
    count += 1

layers.append(layer)

def print_image(image):
    for i, digit in enumerate(image):
        if i % width == 0:
            print()

        if digit == 1:
            print('X', end='')
        elif digit == 0:
            print(' ', end='')
    print()

final = [2] * size

for layer in layers:
    for i, digit in enumerate(layer):
        if digit != 2 and final[i] == 2:
            final[i] = digit
        
print_image(final)

