package day12

import (
	"aoc/util"
	"strings"
	"unicode"
)

// start,A,b,A,c,A,end
// start,A,b,A,end
// start,A,b,end
// start,A,c,A,b,A,end
// start,A,c,A,b,end
// start,A,c,A,end
// start,A,end
// start,b,A,c,A,end
// start,b,A,end
// start,b,end

func part1(inputfile string) int {
	data, _ := util.ReadFileToStringSlice(inputfile)
	t := util.NewTree()

	for _, line := range data {
		tokens := util.ParseTokens(line)
		t.AddChild(tokens.Strs[0], tokens.Strs[1])
	}

	allPaths := &[][]string{}
	recurPaths(&t, "start", []string{}, map[string]bool{}, allPaths)

	result := len(*allPaths)

	return result
}

func part2(inputfile string) int {
	data, _ := util.ReadFileToStringSlice(inputfile)
	t := util.NewTree()

	for _, line := range data {
		tokens := util.ParseTokens(line)
		t.AddChild(tokens.Strs[0], tokens.Strs[1])
	}

	allPaths := map[string]bool{}
	recurPathsP2(&t, "start", []string{}, map[string]int{}, allPaths, false)

	// for k, _ := range allPaths {
	// 	fmt.Println(k)
	// }

	result := len(allPaths)

	return result
}

func copyMap(s map[string]bool) map[string]bool {
	d := map[string]bool{}
	for k, v := range s {
		d[k] = v
	}

	return d
}

func copyMapInt(s map[string]int) map[string]int {
	d := map[string]int{}
	for k, v := range s {
		d[k] = v
	}

	return d
}

func copySlice(s []string) []string {
	d := make([]string, len(s))
	copy(d, s)
	return d
}

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func recurPathsP2(t *util.Tree, curNode string, path []string, visited map[string]int, allPaths map[string]bool, twice bool) {
	if IsLower(curNode) {
		if curNode == "end" {
			path = append(path, curNode)
			pathS := strings.Join(path, " ")
			allPaths[pathS] = true
			return
		}

		// if the cur lower-case node has already been visited
		if visited[curNode] > 0 {
			if curNode == "start" {
				// start cannot be visited more than once
				return
			}

			// if already have a node that has been visited twice stop
			if twice {
				return
			}
		}

		// otherwise, mark this one as visited
		visited[curNode]++
		if visited[curNode] > 1 {
			twice = true
		}
	}

	path = append(path, curNode)

	for _, cRaw := range t.GetChildren(curNode) {
		c := cRaw.(string)

		nPath := copySlice(path)
		nVisited := copyMapInt(visited)

		recurPathsP2(t, c, nPath, nVisited, allPaths, twice)
	}

	for _, cRaw := range t.GetParents(curNode) {
		c := cRaw.(string)

		nPath := copySlice(path)
		nVisited := copyMapInt(visited)

		recurPathsP2(t, c, nPath, nVisited, allPaths, twice)
	}
}

func recurPaths(t *util.Tree, curNode string, path []string, visited map[string]bool, allPaths *[][]string) {
	if IsLower(curNode) {
		if visited[curNode] {
			// end this branch
			return
		}

		if curNode == "end" {
			path = append(path, curNode)
			*allPaths = append(*allPaths, path)
			return
		}

		// otherwise, mark this one as visited
		visited[curNode] = true
	}

	path = append(path, curNode)

	for _, cRaw := range t.GetChildren(curNode) {
		c := cRaw.(string)

		nPath := copySlice(path)
		nVisited := copyMap(visited)

		recurPaths(t, c, nPath, nVisited, allPaths)
	}

	for _, cRaw := range t.GetParents(curNode) {
		c := cRaw.(string)

		nPath := copySlice(path)
		nVisited := copyMap(visited)

		recurPaths(t, c, nPath, nVisited, allPaths)
	}
}
