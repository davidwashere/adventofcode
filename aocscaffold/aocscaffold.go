package main

import (
	"aoc/util"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const dayTemplate = `
package day{{.DayName}}

import (
	"aoc/util"
	"fmt"
)

func part1(inputfile string) int {
	data, _ := util.ReadFileToStringSlice(inputfile)
	// data, _ := util.ReadFileToIntSlice(inputfile)
	// data, _ := util.ReadFileToStringSliceWithDelim(inputfile, "\n\n")
	// grid := util.NewInfinityGridFromFile(inputfile, ".")

	for _, line := range data {
		tokens := util.ParseTokens(line)
		// ints := util.ParseInts(line)
		// strs := util.ParseStrs(line)
		// words := util.ParseWords(line)

		fmt.Println(tokens)
	}

	result := 0

	return result
}

func part2(inputfile string) int {
	return 0
}
`

const dayTestTemplate = `
package day{{.DayName}}

import (
	"fmt"
	"testing"
)

func TestP1(t *testing.T) {
	got := part1("sample.txt")
	fmt.Printf("Got: %v\n", got)
	// want := 0
	// if got != want {
	// 	t.Errorf("Got %v but want %v", got, want)
	// }
}

func TestP1_Actual(t *testing.T) {
	got := part1("input.txt")
	fmt.Printf("Got: %v\n", got)
	// want := 0
	// if got != want {
	// 	t.Errorf("Got %v but want %v", got, want)
	// }
}

func TestP2(t *testing.T) {
	got := part2("sample.txt")
	fmt.Printf("Got: %v\n", got)
	// want := 0
	// if got != want {
	// 	t.Errorf("Got %v but want %v", got, want)
	// }
}

func TestP2_Actual(t *testing.T) {
	got := part2("input.txt")
	fmt.Printf("Got: %v\n", got)
	// want := 0
	// if got != want {
	// 	t.Errorf("Got %v but want %v", got, want)
	// }
}
`

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func usage() {
	prog := filepath.Base(os.Args[0])
	fmt.Printf("Usage: %v DAY#\n", prog)
}

type tmplfields struct {
	DayName     string
	PackageName string
}

func createDir(packageName string) {
	exists, err := util.Exists(packageName)
	check(err)
	if exists {
		fmt.Printf("Directory %v already exists\n", packageName)
		os.Exit(1)
	}

	err = os.MkdirAll(packageName, 0700)
	check(err)
}

func clean(packageName string) {
	exists, err := util.Exists(packageName)
	check(err)
	if !exists {
		return
	}

	err = os.RemoveAll(packageName)
	check(err)
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	day, err := strconv.Atoi(os.Args[1])
	check(err)

	dayS := fmt.Sprintf("%02d", day)

	now := time.Now()

	yearS := fmt.Sprintf("%04d", now.Year())

	packageName := fmt.Sprintf("%s/day%s", yearS, dayS)

	data := tmplfields{}
	data.DayName = dayS
	data.PackageName = packageName

	// clean(packageName) // TODO: Remove after testing complete
	createDir(packageName)

	filename := fmt.Sprintf("day%s.go", dayS)
	f, err := os.Create(filepath.Join(packageName, filename))
	check(err)

	t := template.Must(template.New("dayTemplate").Parse(strings.TrimSpace(dayTemplate)))
	err = t.Execute(f, data)
	f.Close()
	check(err)

	filename = fmt.Sprintf("day%s_test.go", dayS)
	f, err = os.Create(filepath.Join(packageName, filename))
	check(err)

	t = template.Must(template.New("dayTestTemplate").Parse(strings.TrimSpace(dayTestTemplate)))
	err = t.Execute(f, data)
	f.Close()
	check(err)

	f, err = os.Create(filepath.Join(packageName, "input.txt"))
	check(err)
	f.Close()

	f, err = os.Create(filepath.Join(packageName, "sample.txt"))
	check(err)
	f.Close()

	fmt.Printf("Scaffold %v successfully\n", packageName)
}
