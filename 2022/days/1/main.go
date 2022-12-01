package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func Day() {
	file, err := os.Open("./day/1/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	out := []int{0}

	for scanner.Scan() {
		if scanner.Text() != "" {
			i, err := strconv.Atoi(scanner.Text())
			out[len(out)-1] += i
			if err != nil {
				log.Fatal(err)
			}
		} else {
			out = append(out, 0)
		}
	}

	sort.Slice(out, func(i, j int) bool {
		return out[i] > out[j]
	})

	fmt.Println(out[0])
	fmt.Println(out[0] + out[1] + out[2])
}
