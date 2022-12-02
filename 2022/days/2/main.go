package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	points := 0
	tpoints := 0
	for scanner.Scan() {
		match := scanner.Text()
		points += int(1 + (rune(match[2]) - 'X') + (4+
			(rune(match[2])-'X')-(rune(match[0])-'A'))%3*3)

		tpoints += int(1 + (rune(match[2])-'X')*3 + (2+
			(rune(match[2])-'X')+(rune(match[0])-'A'))%3)
	}

	fmt.Println(points)
	fmt.Println(tpoints)
}
