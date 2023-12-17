package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isNum(d byte) bool {
	return d >= 48 && d <= 57
}

// Durchläuft einen Text und liest die
// Zahlen in ein Array
// param: text string
// return stringArray mit Textzahlen
func digitArr(txt string) []string {
	d := []string{}
	for _, s := range txt {
		if s >= 48 && s <= 57 {
			d = append(d, string(s))
		}
	}
	return d
}

// Splittet eine Textdatei
// file string: Dateipfad
// retrun []string mit den Zeilen
func splitFile(file string) []string {
	readFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	txt := []string{}
	for fileScanner.Scan() {
		txt = append(txt, fileScanner.Text())
	}
	return txt
}

func checkDigitOrPoint(b byte) bool {
	return (b >= 48 && b <= 57) || b == 46
}

func sumString(numbers []string) int {
	s := 0
	for _, n := range numbers {
		i, _ := strconv.Atoi(n)
		s += i
	}
	return s
}
