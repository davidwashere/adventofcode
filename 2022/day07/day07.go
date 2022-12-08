package day07

import (
	"aoc/util"
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

func buildFS(inputfile string) map[string]*dir {
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

	var populateAggSizes func(cur *dir)
	populateAggSizes = func(cur *dir) {
		if len(cur.dirs) == 0 {
			return
		}

		totalSubSize := 0
		for _, d := range cur.dirs {
			populateAggSizes(d)
			totalSubSize += d.totalSize
		}

		cur.totalSize += totalSubSize
	}

	root := dirs["/"]
	populateAggSizes(root)

	return dirs
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

func part1(inputfile string) int {
	dirs := buildFS(inputfile)

	result := 0
	for _, d := range dirs {
		if d.totalSize < 100000 {
			result += d.totalSize
		}
	}

	return result
}

func part2(inputfile string) int {
	dirs := buildFS(inputfile)

	root := dirs["/"]

	unused := 70000000 - root.totalSize

	needToDeleteThisMuch := 30000000 - unused

	low := util.MaxInt
	for _, d := range dirs {
		if d.totalSize >= needToDeleteThisMuch {
			low = util.Min(low, d.totalSize)
		}
	}

	return low
}
