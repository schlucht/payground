package main

import (
	"fmt"
	"regexp"
	"strings"
)

var test = []string{
	"123.............458...689.556..3............197...../582/.......720.........................515..352..286.........670.741.....895.626....321",
	"...910.743..........................13..........................*.............775...956........@.........*................971.-.............",
	"....*......406.507.97..846..............968+.........253........730...574............#....308......*.....798..............*.......894.......",
	"....555...............*..**..%................980+43..=..239..........*......495................638.111.........*490...124...*........576...",
}
var testKlein = []string{
	"467..114..",
	"...*......",
	"..35..633.",
	"......#...",
	"617*......",
	".....+.58.",
	"..592.....",
	"......755.",
	"...$.*....",
	".664.598..",
}

// Z1: 515 720 286 895
// Z2: 910 956 971
// Z3: 846 968 253 730 574 798 971
// Z4: 555 980 43 495 638 111 490
// Z5:
func splitLine(lines []string) [][]string {
	// reg := regexp.MustCompile(`.{0,1}\d+.{0,1}`)
	reg := regexp.MustCompile(`\d+`)

	linedigits := [][]string{}
	for _, line := range lines {
		digits := reg.FindAllString(line, -1)
		pos := reg.FindAllStringIndex(line, -1)
		var zahl []string
		for i, b := range digits {
			a := ""
			c := ""

			if pos[i][0] > 0 {
				a = string(line[pos[i][0]-1])
			} else {
				a = ""
			}
			if pos[i][1] < len(line) {
				c = string(line[pos[i][1]])
			} else {
				c = ""
			}
			//fmt.Println(pos[i], a, b, c)
			s := fmt.Sprintf("%s%s%s",
				a,
				b,
				c,
			)
			//fmt.Println(s)
			zahl = append(zahl, s)

		}
		linedigits = append(linedigits, zahl)
	}
	return linedigits
}

func tag3() int {
	summe := 0

	trim := regexp.MustCompile(`\D`)

	txt := splitFile("tag3.txt")
	linedigits := splitLine(txt)

	var io []string
	for i, line := range linedigits {
		if len(line) > 0 {
			for _, digit := range line {
				pos := strings.Index(txt[i], digit)
				fmt.Println(digit)
				if i >= 0 {
					// Links
					if !checkDigitOrPoint(txt[i][pos]) {
						io = append(io, trim.ReplaceAllString(digit, ""))
						continue
					}
					// rechts
					if !checkDigitOrPoint(txt[i][pos+len(digit)-1]) {
						io = append(io, trim.ReplaceAllString(digit, ""))
						continue
					}
				}
				if i < len(linedigits)-1 {
					for d := range digit {
						if !checkDigitOrPoint(txt[i+1][pos+d]) {
							io = append(io, trim.ReplaceAllString(digit, ""))
							continue
						}
					}
				}
				if i > 0 {
					for d := range digit {
						if !checkDigitOrPoint(txt[i-1][pos+d]) {
							io = append(io, trim.ReplaceAllString(digit, ""))
							continue
						}
					}
				}
			}
		}
	}

	fmt.Println(io)
	summe = sumString(io)
	return summe
}
