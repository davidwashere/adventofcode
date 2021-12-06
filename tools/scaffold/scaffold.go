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
	util.Check(err)
	if exists {
		fmt.Printf("Directory %v already exists\n", packageName)
		os.Exit(1)
	}

	err = os.MkdirAll(packageName, 0700)
	util.Check(err)
}

func writeFileFromTemplate(outFileFormat string, tmpl []byte, fields tmplfields) {
	filename := fmt.Sprintf(outFileFormat, fields.DayName)
	f, err := os.Create(filepath.Join(fields.PackageName, filename))
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

	dayS := fmt.Sprintf("%02d", day)

	now := time.Now()

	yearS := fmt.Sprintf("%04d", now.Year())

	packageName := fmt.Sprintf("%s/day%s", yearS, dayS)

	data := tmplfields{}
	data.DayName = dayS
	data.PackageName = packageName

	createDir(packageName)

	writeFileFromTemplate("day%s.go", dayTmpl, data)
	writeFileFromTemplate("day%s_test.go", dayTestTmpl, data)

	f, err := os.Create(filepath.Join(packageName, "input.txt"))
	util.Check(err)
	f.Close()

	f, err = os.Create(filepath.Join(packageName, "sample.txt"))
	util.Check(err)
	f.Close()

	fmt.Printf("Scaffold %v successfully\n", packageName)
}
