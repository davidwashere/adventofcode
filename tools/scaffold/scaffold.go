package main

import (
	"aoc/util"
	_ "embed"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

//go:embed day.tmpl
var dayTmpl []byte

//go:embed dayTest.tmpl
var dayTestTmpl []byte

type tmplfields struct {
	DayName string
	OutPath string
}

func usage() {
	prog := filepath.Base(os.Args[0])
	fmt.Printf("Usage: %v DAY#\n", prog)
}

func createDir(path string) {
	exists, err := util.Exists(path)
	util.Check(err)
	if exists {
		fmt.Printf("Directory %v already exists\n", path)
		os.Exit(1)
	}

	err = os.MkdirAll(path, 0700)
	util.Check(err)
}

func writeFileFromTemplate(outFileFormat string, tmpl []byte, fields tmplfields) {
	filename := fmt.Sprintf(outFileFormat, fields.DayName)
	f, err := os.Create(filepath.Join(fields.OutPath, filename))
	util.Check(err)
	defer f.Close()

	t := template.Must(template.New("t").Parse(string(tmpl)))
	err = t.Execute(f, fields)
	util.Check(err)
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	day, err := strconv.Atoi(os.Args[1])
	util.Check(err)

	// Two digit day as string: 01, 02, 10, 12, etc
	dayS := fmt.Sprintf("%02d", day)

	now := time.Now()

	// Current year as four digit string: "2021"
	yearS := fmt.Sprintf("%04d", now.Year())

	// Path to store output 2021/day06
	outPath := fmt.Sprintf("%s/day%s", yearS, dayS)

	// Fields needed by templates
	data := tmplfields{}
	data.DayName = dayS
	data.OutPath = outPath

	// Create the directory for the day
	createDir(outPath)

	// Write files to the outPath
	writeFileFromTemplate("day%s.go", dayTmpl, data)
	writeFileFromTemplate("day%s_test.go", dayTestTmpl, data)

	f, err := os.Create(filepath.Join(outPath, "input.txt"))
	util.Check(err)
	f.Close()

	f, err = os.Create(filepath.Join(outPath, "sample.txt"))
	util.Check(err)
	f.Close()

	fmt.Printf("Scaffold %v successfully\n", outPath)
}
