package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var zahlenM = map[string]string{
	"one": "o1e", "two": "t2o", "three": "t3e", "four": "f4r", "five": "f5e", "six": "s6x", "seven": "s7n", "eight": "e8t", "nine": "n9e",
}
var str = []string{
	"two1nine",
	"eightwothree",
	"abcone2threexyz",
	"xtwone3four",
	"4nineeightseven2",
	"zoneight234",
	"7pqrstsixteen",
}

func a1() int {
	sum := []string{}
	for _, txt := range splitFile("tag1.txt") {
		digit := digitArr(txt)

		if len(digit) == 1 {
			sum = append(sum, fmt.Sprintf("%s%s", digit[0], digit[0]))
		} else {
			sum = append(sum, fmt.Sprintf("%s%s", digit[0], digit[len(digit)-1]))
		}
	}
	total := 0
	for _, i := range sum {
		m, _ := strconv.Atoi(i)
		total += int(m)
	}
	return total
}

func a2() int {
	txta := splitFile("tag1.txt")
	reg := regexp.MustCompile(`[\D]`)
	t := 0
	var b1 string
	for _, s := range txta {
		for k, v := range zahlenM {
			s = strings.ReplaceAll(s, k, v)
		}

		s2 := reg.ReplaceAllString(s, ``)
		if len(s) > 1 {
			b1 = string(s2[0]) + string(s2[len(s2)-1])
		} else {
			b1 = string(s2[0]) + string(s2[0])
		}

		intV1, _ := strconv.Atoi(b1)

		t += intV1
	}

	return t
}
