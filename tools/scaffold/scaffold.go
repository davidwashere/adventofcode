package main

import (
	"aoc/util"
	_ "embed"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"golang.org/x/net/publicsuffix"
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
	fmt.Println("Usage:")
	fmt.Printf("  To scaffold full day (no input):\n")
	fmt.Printf("    %v DAY#\n", prog)
	fmt.Printf("  To scaffold input:\n")
	fmt.Printf("    %v DAY# input\n", prog)
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

func getClient(u string) *http.Client {
	session := os.Getenv("SESSION_COOKIE")

	if len(session) == 0 {
		log.Println("missing SESSION_COOKIE env key")
		os.Exit(1)
	}

	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	util.Check(err)

	url, err := url.Parse(u)
	util.Check(err)

	cookie := http.Cookie{
		Name:  "session",
		Value: session,
	}

	jar.SetCookies(url, []*http.Cookie{&cookie})

	client := &http.Client{
		Jar: jar,
	}

	return client
}

func getInputFromAoC(day, year int) []byte {
	url := fmt.Sprintf("https://adventofcode.com/%v/day/%v/input", year, day)
	log.Printf("pulling from %v", url)

	client := getClient(url)
	// resp, err := http.Get(url)
	resp, err := client.Get(url)
	util.Check(err)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Printf("failed pulling input [%v] %v", resp.StatusCode, resp.Status)
		os.Exit(1)
	}

	log.Printf("pull success, reading data...")
	data, err := ioutil.ReadAll(resp.Body)
	util.Check(err)

	return data
}

func getDayYear() (int, int) {
	day, err := strconv.Atoi(os.Args[1])
	util.Check(err)

	now := time.Now()

	return day, now.Year()
}

func getOutPath(day, year int) string {
	// Two digit day as string: 01, 02, 10, 12, etc
	dayS := fmt.Sprintf("%02d", day)

	// Current year as four digit string: "2021"
	yearS := fmt.Sprintf("%04d", year)

	// Path to store output 2021/day06
	outPath := fmt.Sprintf("%s/day%s", yearS, dayS)

	return outPath
}

func scaffoldDay() {
	day, year := getDayYear()

	outPath := getOutPath(day, year)

	// Fields needed by templates
	data := tmplfields{}
	data.DayName = fmt.Sprintf("%02d", day)
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

	fmt.Printf("scaffold %v successfully\n", outPath)
}

func pullDayInput() {
	day, year := getDayYear()
	outPath := getOutPath(day, year)
	data := getInputFromAoC(day, year)

	filename := filepath.Join(outPath, "input.txt")
	// os.Stat(filename)

	log.Printf("writing data to %v", filename)
	err := ioutil.WriteFile(filename, data, 0644)
	util.Check(err)

	log.Printf("input successfully written")
}

func main() {

	n := len(os.Args)
	if n != 2 && n != 3 {
		usage()
		os.Exit(1)
	}

	if n == 2 {
		scaffoldDay()
		return
	}

	if n == 3 && os.Args[2] == "input" {
		err := util.LoadEnv()
		util.Check(err)
		pullDayInput()
		return
	}

}
