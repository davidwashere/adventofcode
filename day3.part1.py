def calc_distance(x2, y2):
    return abs(x2) + abs(y2)

filename, max_coord = 'day3.data.sample.txt', 1000
filename, max_coord = 'day3.data.actual.txt', 20000

with open(filename, 'rt') as f:
    contents = f.read()

lines = contents.splitlines()

# for line in lines:
#     sections = line.split(',')

#     for section in sections:
#         direction = section[0]
#         distance = int(section[1:])
#         print(f'direction={direction}, distance={distance}') 

intersections = list()

print(f'Creating Grids... ', end='')
grid_posX_posY = [[False for i in range(max_coord)] for j in range(max_coord)]
grid_posX_negY = [[False for i in range(max_coord)] for j in range(max_coord)]
grid_negX_posY = [[False for i in range(max_coord)] for j in range(max_coord)]
grid_negX_negY = [[False for i in range(max_coord)] for j in range(max_coord)]
print(f'done')


for wire_num, wire in enumerate(lines):
    print(f'\nStarting Wire #{wire_num}')
    sections = wire.split(',')
    curX = 0
    curY = 0
    try:
        for section in sections:
            direction = section[0]
            distance = int(section[1:])

            print(f'{direction}{distance} Start[{curX}, {curY}]') 

            for i in range(0,distance):
                if direction == 'R':
                    curX += 1
                    # curX += distance
                elif direction == 'L':
                    curX -= 1
                    # curX -= distance
                elif direction == 'U':
                    curY += 1
                    # curY += distance
                elif direction == 'D':
                    curY -= 1
                    # curY += distance

                if curX >= 0 and curY >= 0:
                    grid = grid_posX_posY
                elif curX >= 0 and curY < 0:
                    grid = grid_posX_negY
                elif curX < 0 and curY < 0:
                    grid = grid_negX_negY
                elif curX < 0 and curY >= 0:
                    grid = grid_negX_posY

                if wire_num != 0 and grid[curX][curY] == True:
                    print(f'   Intersction at {curX},{curY}')
                    intersections.append((curX, curY))
                elif wire_num != 1:
                    grid[curX][curY] = True


            print(f'   Ending at [{curX}, {curY}]') 
    except:
        print(f'ERROR: {curX}, {curY}')
        raise

min = 10000
print('\nCalculating Distances...')
for intersection in intersections:
    distance = calc_distance(intersection[0], intersection[1])
    print(f'{intersection} - distance: {distance}')
    if distance < min:
        min = distance

print(f'\nShortest distance = {min}')

