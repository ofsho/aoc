package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	points := 0
	tscore := 0
	for scanner.Scan() {
		match := strings.Split(scanner.Text(), " ")
		points += (ar(ag(match[1])-ag(match[0])) + ag(match[1]))

		tscore += (rval(match[1]) + part2(match[0], match[1]))
		fmt.Println(part2("A", "X"))
	}

	fmt.Println(points)
	fmt.Println(tscore)
}

func ar(i int) int {
	if i == 1 || i == -2 {
		return 6
	} else if i == 0 {
		return 3
	} else {
		return 0
	}
}

func ag(str string) int {
	if str == "A" || str == "X" {
		return 1
	} else if str == "B" || str == "Y" {
		return 2
	} else if str == "C" || str == "Z" {
		return 3
	}
	return 0
}

func rval(i string) int {
	if i == "X" {
		return 0
	} else if i == "Y" {
		return 3
	} else {
		return 6
	}
}

func part2(i, b string) int {
	switch b {
	case "X":
		n := 0
		for range []int{1, 2, 3} {
			n += 1
			if ar(n-ag(i)) == 0 {
				return n
			}
		}
	case "Y":
		n := 0
		for range []int{1, 2, 3} {
			n += 1
			if ar(n-ag(i)) == 3 {
				return n
			}
		}
	case "Z":
		n := 0
		for range []int{1, 2, 3} {
			n += 1
			if ar(n-ag(i)) == 6 {
				return n
			}
		}
	}
	return 0
}
