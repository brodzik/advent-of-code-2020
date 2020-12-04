package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func hasKey(passport map[string]string, key string) bool {
	_, ok := passport[key]
	return ok
}

func getValue(passport map[string]string, key string) string {
	value, _ := passport[key]
	return value
}

func isValidColor(hcl string) bool {
	if len(hcl) == 7 && hcl[0] == '#' {
		for _, v := range hcl[1:] {
			if !(v >= 'a' && v <= 'f') && !(v >= '0' && v <= '9') {
				return false
			}
		}

		return true
	} else {
		return false
	}
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	data := ""
	valid := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			passport := make(map[string]string)
			for _, field := range strings.Split(data[:len(data)-1], " ") {
				s := strings.Split(field, ":")
				passport[s[0]] = s[1]
			}

			if hasKey(passport, "byr") && hasKey(passport, "iyr") && hasKey(passport, "eyr") && hasKey(passport, "hgt") && hasKey(passport, "hcl") && hasKey(passport, "ecl") && hasKey(passport, "pid") {
				byr, _ := strconv.Atoi(getValue(passport, "byr"))
				iyr, _ := strconv.Atoi(getValue(passport, "iyr"))
				eyr, _ := strconv.Atoi(getValue(passport, "eyr"))
				hgt := getValue(passport, "hgt")
				hcl := getValue(passport, "hcl")
				ecl := getValue(passport, "ecl")
				pid := getValue(passport, "pid")

				cm := strings.Contains(hgt, "cm")
				h, _ := strconv.Atoi(hgt[:len(hgt)-2])

				if byr >= 1920 && byr <= 2002 && iyr >= 2010 && iyr <= 2020 && eyr >= 2020 && eyr <= 2030 && (ecl == "amb" || ecl == "blu" || ecl == "brn" || ecl == "gry" || ecl == "grn" || ecl == "hzl" || ecl == "oth") && len(pid) == 9 && isValidColor(hcl) && ((cm && h >= 150 && h <= 193) || (!cm && h >= 59 && h <= 76)) {
					valid += 1
				}
			}

			data = ""
		} else {
			data += line + " "
		}
	}

	fmt.Println(valid)
}
