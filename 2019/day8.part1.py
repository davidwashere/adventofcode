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

# print(len(layers))
# print(layers[0])
# print(layers[1])
# print(layers[-1])

min_zero_digits = 99999999
min_zero_digits_layer = None

for i, layer in enumerate(layers):
    zero_count = 0
    for digit in layer:
        if digit == 0:
            zero_count += 1

    print(f"Layer {i+1} has {zero_count} zeros")
    if zero_count < min_zero_digits:
        min_zero_digits = zero_count
        min_zero_digits_layer = layer
    
print(f"Least amount of zeros: {min_zero_digits}")

num_ones = 0
num_twos = 0
for digit in min_zero_digits_layer:
    if digit == 1:
        num_ones += 1
    elif digit == 2:
        num_twos += 1

print(f"Result of #1's * #2's: {num_ones * num_twos}")