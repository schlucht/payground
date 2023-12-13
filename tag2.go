package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Game struct {
	id     int
	runden []Runde	
}
type Runde struct {
	crl Color
}
type Color struct {
	blue  int
	green int
	red   int
}

var test2 = []string{
	"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
	"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
	"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
	"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
	"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
}
var games []Game

func tag2() int {
	createGames()
	fmt.Println("------- Meine Daten ----")
	summe := 0

	for _, game := range games {		
		summe += game.id
		for _, runde := range game.runden {
			if runde.crl.blue > 14 ||
				runde.crl.red > 12 ||
				runde.crl.green > 13 {
					summe -= game.id
					fmt.Printf("ID: %d, Blau: %d, Rot: %d, Grün: %d\n",game.id, runde.crl.blue, runde.crl.red, runde.crl.green)
					break				
				}
		}
	}
	fmt.Println("------- Ende ----")
	return summe
}
func tag2_2() int {
	createGames()
	fmt.Println("------- Meine Daten ----")
	summe := 0
	// fmt.Printf("ID: %d, Blau: %d, Rot: %d, Grün: %d\n",game.id, runde.crl.blue, runde.crl.red, runde.crl.green)
	for _, game := range games {		
		
			var r,g,b int
		for _, runde := range game.runden {
			if r < runde.crl.red{
				r = runde.crl.red
			}
			if g < runde.crl.green{
				g = runde.crl.green
			}
			if b < runde.crl.blue{
				b = runde.crl.blue		
			}	
		}
		potenz := r *g *b
		summe += potenz
		fmt.Printf("R:%d G:%d B:%d P: %d S: %d \n", r, g, b, potenz, summe)
	}

	fmt.Println("------- Ende ----")
	return summe
}

func createGames() {
	for i, txt := range splitFile("tag2.txt") {
		runden := []Runde{}
		game := Game{
			id: i + 1,
		}
		txtRunde := strings.Split(txt, ":")[1]
		rund := Runde{}
		for _, rnd := range strings.Split(txtRunde, ";") {
			clClass := Color{
				blue:  0,
				red:   0,
				green: 0,
			}
			for _, crl := range strings.Split(rnd, ",") {
				crlNr := strings.Split(strings.Trim(crl, " "), " ")
				// fmt.Println(crlNr)
				nr := strings.Trim(crlNr[0], " ")
				c := strings.Trim(crlNr[1], " ")
				switch c {
				case "blue":
					clClass.blue = toInt(nr)
				case "red":
					clClass.red = toInt(nr)
				case "green":
					clClass.green = toInt(nr)
				}
			}
			rund.crl = clClass
			// fmt.Println(rund)
			runden = append(runden, rund)

		}
		game.runden = runden		
		games = append(games, game)		
	}
}

func toInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
