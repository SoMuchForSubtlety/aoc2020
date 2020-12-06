package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/SoMuchForSubtlety/aoc2020/go/pkg/input"
)

func main() {
	passports := strings.Split(input.Read(4), "\n\n")

	var invalid1 int
	var valid2 int
	for _, passport := range passports {
		fields, invalid := fields(strings.ReplaceAll(passport, "\n", " "))
		for _, k := range []string{"byr", "iyr", "eyr", "hcl", "ecl", "hgt", "pid"} {
			if _, ok := fields[k]; !ok {
				invalid1++
				break
			}
		}
		if invalid {
			continue
		}

		year, err := strconv.Atoi(fields["byr"])
		if len(fields["byr"]) != 4 || year > 2002 || year < 1920 || err != nil {
			continue
		}

		iyear, err := strconv.Atoi(fields["iyr"])
		if len(fields["iyr"]) != 4 || iyear > 2020 || iyear < 2010 || err != nil {
			continue
		}

		eyear, err := strconv.Atoi(fields["eyr"])
		if len(fields["eyr"]) != 4 || eyear > 2030 || eyear < 2020 || err != nil {
			continue
		}

		height := fields["hgt"]
		if len(height) <= 2 {
			continue
		} else {
			hightnum, err := strconv.Atoi(height[:len(height)-2])
			unit := height[len(height)-2:]
			if err != nil ||
				(unit == "cm" && (hightnum > 193 || hightnum < 150)) ||
				(unit == "in" && (hightnum > 76 || hightnum < 59)) ||
				(unit != "cm" && unit != "in") {

				continue
			}
		}

		if !regexp.MustCompile("^#[0-9a-f]{6}$").MatchString(fields["hcl"]) ||
			!regexp.MustCompile("^[0-9]{9}$").MatchString(fields["pid"]) {
			continue
		}

		if !(fields["ecl"] == "amb") && !(fields["ecl"] == "blu") && !(fields["ecl"] == "brn") && !(fields["ecl"] == "gry") && !(fields["ecl"] == "grn") && !(fields["ecl"] == "hzl") && !(fields["ecl"] == "oth") {
			continue
		}

		valid2++
	}

	fmt.Println("part 1:", len(passports)-invalid1)
	fmt.Println("part 2:", valid2)
}

func fields(passport string) (map[string]string, bool) {
	fields := make(map[string]string)
	for _, field := range strings.Fields(passport) {
		kv := strings.Split(field, ":")
		if _, ok := fields[kv[0]]; ok {
			return fields, true
		}
		fields[kv[0]] = kv[1]
	}

	return fields, false
}
