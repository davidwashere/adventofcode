package disjointset

// https://en.wikipedia.org/wiki/Disjoint-set_data_structure

type DisjointSet[T comparable] struct {
	parent map[T]T
	rank   map[T]int
	size   map[T]int // tracks size of each root's component
}

func New[T comparable]() *DisjointSet[T] {
	return &DisjointSet[T]{
		parent: make(map[T]T),
		rank:   make(map[T]int),
		size:   make(map[T]int),
	}
}

// AddNode adds a node if it doesn't exist
func (ds *DisjointSet[T]) AddNode(node T) {
	if _, exists := ds.parent[node]; !exists {
		ds.parent[node] = node
		ds.rank[node] = 0
		ds.size[node] = 1
	}
}

// Find returns the root of the node's graph (with path compression)
func (ds *DisjointSet[T]) Find(node T) T {
	ds.AddNode(node)
	if ds.parent[node] != node {
		ds.parent[node] = ds.Find(ds.parent[node]) // path compression
	}
	return ds.parent[node]
}

// Union connects two nodes into the same graph
func (ds *DisjointSet[T]) Union(a, b T) {
	rootA := ds.Find(a)
	rootB := ds.Find(b)

	if rootA == rootB {
		return // already in the same graph
	}

	// union by rank
	if ds.rank[rootA] < ds.rank[rootB] {
		ds.parent[rootA] = rootB
		ds.size[rootB] += ds.size[rootA]
	} else if ds.rank[rootA] > ds.rank[rootB] {
		ds.parent[rootB] = rootA
		ds.size[rootA] += ds.size[rootB]
	} else {
		ds.parent[rootB] = rootA
		ds.size[rootA] += ds.size[rootB]
		ds.rank[rootA]++
	}
}

// CountNodes returns the number of nodes in the graph containing the given node
func (ds *DisjointSet[T]) CountNodes(node T) int {
	root := ds.Find(node)
	return ds.size[root]
}

// GetAllGraphs returns a map of root -> list of nodes in that graph
func (ds *DisjointSet[T]) GetAllGraphs() map[T][]T {
	graphs := make(map[T][]T)
	for node := range ds.parent {
		root := ds.Find(node)
		graphs[root] = append(graphs[root], node)
	}
	return graphs
}

// NumGraphs returns to total number of graphs in the disjoint set
func (ds *DisjointSet[T]) NumGraphs() int {
	graphs := ds.GetAllGraphs()
	return len(graphs)
}

// NumNodes returns the total number of nodes in the disjoint set
func (ds *DisjointSet[T]) NumNodes() int {
	graphs := ds.GetAllGraphs()
	sum := 0
	for _, g := range graphs {
		sum += len(g)
	}
	return sum
}

// Connected checks if two nodes are in the same graph
func (ds *DisjointSet[T]) Connected(a, b T) bool {
	return ds.Find(a) == ds.Find(b)
}
