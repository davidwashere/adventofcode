package util

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// ParseFile parses file line by line calling `lineHandler` with the []byte's of each line
func ParseFile(filename string, lineHandler func([]byte)) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineHandler(scanner.Bytes())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// ParseFileAsString same as `ParseFile` but converts line to a string before calling `lineHandler`
func ParseFileAsString(filename string, lineHandler func(string)) error {
	return ParseFile(filename, func(line []byte) {
		lineHandler(string(line))
	})
}

// ReadFileToStringSlice Read entire file into string slice, one entry per
// line in file
func ReadFileToStringSlice(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

// ReadFileToIntSlice Read entire file into int slice, one entry per
// line in file
func ReadFileToIntSlice(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		item := scanner.Text()
		num, err := strconv.Atoi(item)
		if err != nil {
			return nil, err
		}
		data = append(data, num)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

// ReadFileToStringSliceWithDelim Read entire file into string slice, entries separated
// by delim
func ReadFileToStringSliceWithDelim(filename string, delim string) ([]string, error) {
	dataB, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	raw := string(dataB)
	data := strings.Split(raw, delim)

	return data, nil
}

// ConvertFromCRLFtoLF will convert a files line endings from CRLF to LF
func ConvertFromCRLFtoLF(filename string) error {
	dataB, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	raw := string(dataB)
	raw = strings.ReplaceAll(raw, "\r\n", "\n")
	err = ioutil.WriteFile(filename, []byte(raw), 0700)
	return err
}

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}
