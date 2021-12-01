package util

import (
	"fmt"
	"strings"
)

// Tree represents a set of relationships between objects, objects
// have parents and children
//
// When children are added, parent relationships are automatically set, and vice versa when
// parents are added
type Tree struct {
	objects  map[treeT]treeT
	children map[treeT][]treeT
	parents  map[treeT][]treeT
}

// NewTree .
func NewTree() Tree {
	return Tree{
		objects:  map[treeT]treeT{},
		children: map[treeT][]treeT{},
		parents:  map[treeT][]treeT{},
	}
}

type treeT interface{}

func mapKeysSlice(m map[treeT]struct{}) []treeT {
	keys := make([]treeT, len(m))

	i := 0
	for k := range m {
		keys[i] = k
		i++
	}

	return keys
}

// GetChildren returns slice of child ids for given id
func (t *Tree) GetChildren(id treeT) []treeT {
	return t.children[id]
}

// GetChildrenInt same as GetChildren with result type asserted to int
func (t *Tree) GetChildrenInt(id treeT) []int {
	orig := t.GetChildren(id)
	asserted := make([]int, len(orig))

	for i, item := range orig {
		asserted[i] = item.(int)
	}

	return asserted
}

// GetChildrenString same as GetChildren with result type asserted to string
func (t *Tree) GetChildrenString(id treeT) []string {
	orig := t.GetChildren(id)
	asserted := make([]string, len(orig))

	for i, item := range orig {
		asserted[i] = item.(string)
	}

	return asserted
}

// GetAllChildren returns slice of child ids (and ids of childs children recursively) for given id
func (t *Tree) GetAllChildren(id treeT) []treeT {
	childs := []treeT{}

	for _, child := range t.GetChildren(id) {
		t.recurChildren(&childs, child)
	}

	return childs
}

// GetAllChildrenInt same as GetAllChildren with result type asserted to int
func (t *Tree) GetAllChildrenInt(id treeT) []int {
	orig := t.GetAllChildren(id)
	asserted := make([]int, len(orig))

	for i, item := range orig {
		asserted[i] = item.(int)
	}

	return asserted
}

// GetAllChildrenString same as GetAllChildren with result type asserted to string
func (t *Tree) GetAllChildrenString(id treeT) []string {
	orig := t.GetAllChildren(id)
	asserted := make([]string, len(orig))

	for i, item := range orig {
		asserted[i] = item.(string)
	}

	return asserted
}

func (t *Tree) recurChildren(childs *[]treeT, id treeT) {
	*childs = append(*childs, id)

	for _, child := range t.GetChildren(id) {
		t.recurChildren(childs, child)
	}
}

// GetAllUniqueChildren same as GetAllChildren with duplicates removed
func (t *Tree) GetAllUniqueChildren(id treeT) []treeT {
	childs := map[treeT]struct{}{}

	for _, child := range t.GetChildren(id) {
		t.recurUniqueChildren(childs, child)
	}

	return mapKeysSlice(childs)
}

// GetAllUniqueChildrenInt same as GetAllUniqueChildren with result type asserted to int
func (t *Tree) GetAllUniqueChildrenInt(id treeT) []int {
	orig := t.GetAllUniqueChildren(id)
	asserted := make([]int, len(orig))

	for i, item := range orig {
		asserted[i] = item.(int)
	}

	return asserted
}

// GetAllUniqueChildrenString same as GetAllUniqueChildren with result type asserted to string
func (t *Tree) GetAllUniqueChildrenString(id treeT) []string {
	orig := t.GetAllUniqueChildren(id)
	asserted := make([]string, len(orig))

	for i, item := range orig {
		asserted[i] = item.(string)
	}

	return asserted
}

func (t *Tree) recurUniqueChildren(childs map[treeT]struct{}, id treeT) {
	if _, ok := childs[id]; ok {
		return
	}
	childs[id] = struct{}{}

	for _, child := range t.GetChildren(id) {
		t.recurUniqueChildren(childs, child)
	}
}

// GetAllParents returns slice of parent ids (and ids of parents parent recursively) for given id
func (t *Tree) GetAllParents(id treeT) []treeT {
	parents := []treeT{}

	for _, parent := range t.GetParents(id) {
		t.recurParents(&parents, parent)
	}

	return parents
}

// GetAllParentsInt same as GetAllParents with result type asserted to int
func (t *Tree) GetAllParentsInt(id treeT) []int {
	orig := t.GetAllParents(id)
	asserted := make([]int, len(orig))

	for i, item := range orig {
		asserted[i] = item.(int)
	}

	return asserted
}

// GetAllParentsString same as GetAllParents with result type asserted to string
func (t *Tree) GetAllParentsString(id treeT) []string {
	orig := t.GetAllParents(id)
	asserted := make([]string, len(orig))

	for i, item := range orig {
		asserted[i] = item.(string)
	}

	return asserted
}

func (t *Tree) recurParents(parents *[]treeT, id treeT) {
	*parents = append(*parents, id)

	for _, parent := range t.GetParents(id) {
		t.recurParents(parents, parent)
	}
}

// GetAllUniqueParents same as GetAllParents with duplicates removed
func (t *Tree) GetAllUniqueParents(id treeT) []treeT {
	parents := map[treeT]struct{}{}

	for _, parent := range t.GetParents(id) {
		t.recurUniqueParents(parents, parent)
	}

	return mapKeysSlice(parents)
}

// GetAllUniqueParentsInt same as GetAllUniqueParents with result type asserted to int
func (t *Tree) GetAllUniqueParentsInt(id treeT) []int {
	orig := t.GetAllUniqueParents(id)
	asserted := make([]int, len(orig))

	for i, item := range orig {
		asserted[i] = item.(int)
	}

	return asserted
}

// GetAllUniqueParentsString same as GetAllUniqueParents with result type asserted to string
func (t *Tree) GetAllUniqueParentsString(id treeT) []string {
	orig := t.GetAllUniqueParents(id)
	asserted := make([]string, len(orig))

	for i, item := range orig {
		asserted[i] = item.(string)
	}

	return asserted
}

func (t *Tree) recurUniqueParents(parents map[treeT]struct{}, id treeT) {
	if _, ok := parents[id]; ok {
		return
	}
	parents[id] = struct{}{}

	for _, parent := range t.GetParents(id) {
		t.recurUniqueParents(parents, parent)
	}
}

// GetParents returns slice of parent ids for given id
func (t *Tree) GetParents(id treeT) []treeT {
	return t.parents[id]
}

// GetParentsInt same as GetParents with result type asserted to int
func (t *Tree) GetParentsInt(id treeT) []int {
	orig := t.GetParents(id)
	asserted := make([]int, len(orig))

	for i, item := range orig {
		asserted[i] = item.(int)
	}

	return asserted
}

// GetParentsString same as GetParents with result type asserted to string
func (t *Tree) GetParentsString(id treeT) []string {
	orig := t.GetParents(id)
	asserted := make([]string, len(orig))

	for i, item := range orig {
		asserted[i] = item.(string)
	}

	return asserted
}

// Set inserts an object into the tree without any child/parent links (orphan)
func (t *Tree) Set(id treeT, obj treeT) {
	t.objects[id] = obj
}

// Get returns the object associated with the id added into the tree via Set
func (t *Tree) Get(id treeT) treeT {
	return t.objects[id]
}

// GetInt same as Get with result type asserted to int
func (t *Tree) GetInt(id treeT) int {
	return t.objects[id].(int)
}

// GetString same as Get with result type asserted to string
func (t *Tree) GetString(id treeT) string {
	return t.objects[id].(string)
}

// AddChild adds a child id to a parent, if parent doesn't exist the parent is added
// to the tree with a nil object
func (t *Tree) AddChild(parent, child treeT) {
	t.addChild(parent, child, false)
}

// AddUniqueChild same as AddChild but will not add the child id if the parent already has the link
func (t *Tree) AddUniqueChild(parent, child treeT) {
	t.addChild(parent, child, true)
}

func (t *Tree) addChild(parent, child treeT, unique bool) {
	if _, ok := t.objects[parent]; !ok {
		t.Set(parent, nil)
	}

	if _, ok := t.objects[child]; !ok {
		t.Set(child, nil)
	}

	if unique {
		if !t.isInFace(t.children[parent], child) {
			t.children[parent] = append(t.children[parent], child)
		}
	} else {
		t.children[parent] = append(t.children[parent], child)
	}

	if !t.isInFace(t.parents[child], parent) {
		t.parents[child] = append(t.parents[child], parent)
	}
}

// AddParent adds a parent id to a child, if child doesn't exist the child is added
// to the tree with a nil object
func (t *Tree) AddParent(child, parent treeT) {
	t.addParent(child, parent, false)
}

// AddUniqueParent same as AddParent but will not add the parent id if the child already has the link
func (t *Tree) AddUniqueParent(child, parent treeT) {
	t.addParent(child, parent, true)
}

func (t *Tree) addParent(child, parent treeT, unique bool) {
	if _, ok := t.objects[child]; !ok {
		t.Set(child, nil)
	}

	if _, ok := t.objects[parent]; !ok {
		t.Set(parent, nil)
	}

	if !unique {
		t.parents[child] = append(t.parents[child], parent)
	} else {
		if !t.isInFace(t.parents[child], parent) {
			t.parents[child] = append(t.parents[child], parent)
		}
	}

	if !t.isInFace(t.children[parent], child) {
		t.children[parent] = append(t.children[parent], child)
	}
}

// isInFace returns true if val found in slice, false otherwise
func (t Tree) isInFace(slice []treeT, val treeT) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}

	return false
}

// StringNode returns a string representation of a node in the tree
func (t Tree) StringNode(id treeT) string {
	var sb strings.Builder
	key := id
	val := t.objects[id]
	sb.WriteString(fmt.Sprintf("%v: %+v\n", key, val))
	children := t.children[key]
	sb.WriteString(fmt.Sprintf("  childrn: %+v\n", children))
	parents := t.parents[key]
	sb.WriteString(fmt.Sprintf("  parents: %+v", parents))

	return sb.String()
}

func (t Tree) String() string {
	var sb strings.Builder
	for key := range t.objects {
		if sb.Len() > 0 {
			sb.WriteString("\n")
		}
		sb.WriteString(t.StringNode(key))
	}
	return sb.String()
}

// CountPaths with return the number of distinct paths from and to the given node ids
// It is assumed that `to` is a child of `from` in the tree and the nodes are acyclic
func (t Tree) CountPaths(from treeT, to treeT) int {
	counts := map[treeT]int{}
	count := 0
	for _, child := range t.GetChildren(from) {
		count += t.recurCountPaths(counts, child, to)
	}

	return count
}

func (t Tree) recurCountPaths(counts map[treeT]int, id, last treeT) int {
	val, ok := counts[id]

	if ok {
		return val
	}

	if id == last {
		return 1
	}

	for _, child := range t.GetChildren(id) {
		counts[id] += t.recurCountPaths(counts, child, last)
	}

	return counts[id]
}
