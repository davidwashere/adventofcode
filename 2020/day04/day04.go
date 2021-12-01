package day04

import (
	"aoc2020/util"
	"fmt"
	"strconv"
	"strings"
)

type Passport struct {
	fields map[string]string
}

func NewPassport() Passport {
	return Passport{
		fields: make(map[string]string),
	}
}

func part1(inputfile string) int {
	data, err := util.ReadFileToStringSlice(inputfile)
	util.Check(err)

	keysWanted := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}

	passports := []Passport{}
	passport := NewPassport()

	for _, line := range data {
		if strings.TrimSpace(line) == "" {
			passports = append(passports, passport)
			passport = NewPassport()

			// fmt.Println()
			continue
		}

		spaceSplit := strings.Split(line, " ")
		for _, item := range spaceSplit {
			itemSplit := strings.Split(item, ":")
			key := itemSplit[0]
			val := itemSplit[1]
			// fmt.Println(key)
			passport.fields[key] = val
		}
		// fmt.Println(passport)
	}

	passports = append(passports, passport)

	valid := 0
	for _, p := range passports {
		fmt.Printf("%+v\n", p)
		if util.AllInMap(keysWanted, p.fields) {
			valid++
		}
	}

	return valid
}

func isValid(passport Passport) bool {
	keysWanted := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	eyesWanted := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

	if !util.AllInMap(keysWanted, passport.fields) {
		return false
	}

	num, err := strconv.Atoi(passport.fields["byr"])
	if err != nil || num < 1920 || num > 2002 {
		return false
	}

	num, err = strconv.Atoi(passport.fields["iyr"])
	if err != nil || num < 2010 || num > 2020 {
		return false
	}

	num, err = strconv.Atoi(passport.fields["eyr"])
	if err != nil || num < 2020 || num > 2030 {
		return false
	}

	val := passport.fields["hgt"]
	if strings.HasSuffix(val, "cm") {
		val = strings.TrimSuffix(val, "cm")
		num, err = strconv.Atoi(val)
		if err != nil || num < 150 || num > 193 {
			return false
		}

	} else if strings.HasSuffix(val, "in") {
		val = strings.TrimSuffix(val, "in")
		num, err = strconv.Atoi(val)
		if err != nil || num < 59 || num > 76 {
			return false
		}
	} else {
		return false
	}

	val = passport.fields["hcl"]
	if !strings.HasPrefix(val, "#") {
		return false
	}
	val = strings.TrimPrefix(val, "#")
	if !util.IsHex(val) {
		return false
	}

	val = passport.fields["ecl"]
	if !util.IsStringIn(eyesWanted, val) {
		return false
	}

	val = passport.fields["pid"]
	if len(val) != 9 {
		return false
	}
	_, err = strconv.Atoi(val)
	if err != nil {
		return false
	}

	return true

}

func part2(inputfile string) int {
	data, err := util.ReadFileToStringSlice(inputfile)
	util.Check(err)

	passports := []Passport{}

	passport := NewPassport()

	for _, line := range data {
		if strings.TrimSpace(line) == "" {
			passports = append(passports, passport)
			passport = NewPassport()

			continue
		}

		spaceSplit := strings.Split(line, " ")
		for _, item := range spaceSplit {
			itemSplit := strings.Split(item, ":")
			key := itemSplit[0]
			val := itemSplit[1]
			passport.fields[key] = val
		}
	}

	passports = append(passports, passport)

	valid := 0
	for _, p := range passports {
		if isValid(p) {
			valid++
		}
	}

	return valid
}
