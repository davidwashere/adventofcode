import sys

filename = 'day6.data.actual.txt'
# filename = 'day6.data.sample.txt'

with open(filename, 'rt') as f:
    contents = f.read()

lines = contents.splitlines()

class Node:
    def __init__(self, id):
        self.depth = 0
        self.id = id
        self.orbits = None
        self.nodes_orbiting_me = []
    
    def now_orbits(self, orbits):
        if self.orbits:
            print(f"{self.id} is already orbiting {self.orbits.id}")
            sys.exit(1)

        orbits.orbitted_by(self)
        self.orbits = orbits
        self.depth = orbits.depth + 1

        self.update_depths_of_nodes_orbiting()

    def update_depths_of_nodes_orbiting(self):
        for node in self.nodes_orbiting_me:
            node.depth = self.depth + 1
            node.update_depths_of_nodes_orbiting()

    def orbitted_by(self, node):
        self.nodes_orbiting_me.append(node)
        self.update_depths_of_nodes_orbiting()
    
    def __str__(self):
        orbits_id = None if not self.orbits else self.orbits.id

        if not self.orbitted_by:
            orbitted_by = None
        else:
            orbitted_by = ""
            for node in self.nodes_orbiting_me:
                orbitted_by += f"{node.id}, "

        
        return f"Node(id={self.id}, orbits={orbits_id}, depth={self.depth}, orbitted_by={orbitted_by})"

    def __repr__(self):
        return self.__str__()

nodes = {}

for line in lines:
    split = line.split(")")
    left = split[0]
    right = split[1]
    print(f"{left} ))) {right}")

    if left not in nodes:
        # print(f"{left} not found, creating - LEFT")
        lnode = Node(left)
        nodes[left] = lnode
    
    lnode = nodes[left]
    
    if right not in nodes:
        # print(f"{right} not found, creating - RIGHT")
        rnode = Node(right)
        rnode.now_orbits(lnode)

        nodes[right] = rnode
    else:
        rnode = nodes[right]
        rnode.now_orbits(lnode)
    print(f"LEFT[{lnode}], RIGHT[{rnode}]")

sum = 0
for node in nodes:
    sum += nodes[node].depth
    # print(nodes[node])

print(f"Depth: {sum}")