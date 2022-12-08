package day07

import (
	"aoc/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type dir struct {
	name      string
	files     []*file
	dirs      []*dir
	totalSize int
}

type file struct {
	name string
	size int
}

func part1(inputfile string) int {
	data, _ := util.ReadFileToStringSlice(inputfile)

	dirs := map[string]*dir{}
	path := util.NewStringStack()
	for _, line := range data {
		parts := strings.Split(line, " ")

		if parts[0] == "$" {
			if parts[1] == "cd" {
				if parts[2] == ".." {
					path.Pop()
				} else {
					// changed to dir
					name := parts[2]
					path.Push(name)

					key := pathKey(path)

					d := new(dir)
					d.name = name

					_, ok := dirs[key]
					if ok {
						panic("Visited dir twice - needs handling")
					}
					dirs[key] = d

					if len(key) > 1 {
						// not root dir, so lets add this dir to parent
						pKey := parentKey(path)
						pDir := dirs[pKey]
						pDir.dirs = append(pDir.dirs, d)
					}
				}
			}
			// ls is ignored
		} else if parts[0] == "dir" {
			// is dir
			// ignore 'dir' - will create when change into the dir - assuming don't care unless look at files under it
			// if next star asks to do something with empty dirs.. poo poo on that
		} else {
			// is file
			name := parts[1]

			sizeStr := parts[0]
			size, _ := strconv.Atoi(sizeStr)

			f := &file{name, size}

			d := dirs[pathKey(path)]
			d.files = append(d.files, f)
			d.totalSize += f.size
		}
	}

	result := DFS(dirs["/"], 100000)

	return result
}

func DFS(root *dir, maxSize int) int {

	total := 0

	var dfsR func(cur *dir) int
	dfsR = func(cur *dir) int {
		if len(cur.dirs) == 0 {
			return cur.totalSize
		}

		totalSubSize := 0
		for _, d := range cur.dirs {
			r := dfsR(d)
			if r < maxSize {
				total += r
				fmt.Printf("adding %v of size %v\n", d.name, r)
			}

			totalSubSize += r
		}

		return totalSubSize + cur.totalSize
	}

	r := dfsR(root)
	if r < maxSize {
		total += r
	}

	return total
}

func pathKey(path util.StringStack) string {
	key := ""
	started := false
	for _, e := range path {
		if started {
			key += ":"
		}
		key += e
		started = true
	}

	return key
}

func parentKey(path util.StringStack) string {
	pPath := make(util.StringStack, len(path))
	copy(pPath, path)
	pPath.Pop()

	return pathKey(pPath)
}

func part2(inputfile string) int {
	data, _ := util.ReadFileToStringSlice(inputfile)

	dirs := map[string]*dir{}
	path := util.NewStringStack()
	for _, line := range data {
		parts := strings.Split(line, " ")

		if parts[0] == "$" {
			if parts[1] == "cd" {
				if parts[2] == ".." {
					path.Pop()
				} else {
					// changed to dir
					name := parts[2]
					path.Push(name)

					key := pathKey(path)

					d := new(dir)
					d.name = name

					_, ok := dirs[key]
					if ok {
						panic("Visited dir twice - needs handling")
					}
					dirs[key] = d

					if len(key) > 1 {
						// not root dir, so lets add this dir to parent
						pKey := parentKey(path)
						pDir := dirs[pKey]
						pDir.dirs = append(pDir.dirs, d)
					}
				}
			}
			// ls is ignored
		} else if parts[0] == "dir" {
		} else {
			// is file
			name := parts[1]

			sizeStr := parts[0]
			size, _ := strconv.Atoi(sizeStr)

			f := &file{name, size}

			d := dirs[pathKey(path)]
			d.files = append(d.files, f)
			d.totalSize += f.size
		}
	}

	root := dirs["/"]

	totalSize := totalSize(root)
	unused := 70000000 - totalSize

	// delete a directory that can free up at least needed amt
	needed := 30000000 - unused

	candidates := dirsWithSizeGreaterThan(root, needed)
	sort.Ints(candidates)

	return candidates[0]
}

func totalSize(root *dir) int {
	var dfsR func(cur *dir) int
	dfsR = func(cur *dir) int {
		if len(cur.dirs) == 0 {
			return cur.totalSize
		}

		totalSubSize := 0
		for _, d := range cur.dirs {
			r := dfsR(d)
			totalSubSize += r
		}

		return totalSubSize + cur.totalSize
	}

	r := dfsR(root)

	return r
}

func dirsWithSizeGreaterThan(root *dir, minSize int) []int {
	candidates := []int{}

	var dfsR func(cur *dir) int
	dfsR = func(cur *dir) int {
		if len(cur.dirs) == 0 {
			return cur.totalSize
		}

		totalSubSize := 0
		for _, d := range cur.dirs {
			r := dfsR(d)
			if r > minSize {
				candidates = append(candidates, r)
			}

			totalSubSize += r
		}

		return totalSubSize + cur.totalSize
	}

	r := dfsR(root)
	if r > minSize {
		candidates = append(candidates, r)
	}

	return candidates
}
