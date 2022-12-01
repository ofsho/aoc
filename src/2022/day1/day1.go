package day1

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func Day() []int {
	out := parse("./day1/input.txt")
	largest := []int{0, 0}
	slargest := []int{0, 0}
	tlargest := []int{0, 0}

	for index, element := range out {
		if element > largest[0] {
			slargest[0] = largest[0]
			slargest[1] = largest[1]

			largest[0] = element
			largest[1] = index
		} else if element > slargest[0] {
			tlargest[0] = slargest[0]
			tlargest[1] = slargest[1]

			slargest[0] = element
			slargest[1] = index
		} else if element > tlargest[0] {
			tlargest[0] = element
			tlargest[1] = index
		}
	}

	return []int{largest[0] + slargest[0] + tlargest[0], largest[1]}
}

func parse(fname string) []int {
	outputs := []int{0}

	/** read the file */
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()

		if input != "\n" && input != "" {
			i, err := strconv.Atoi(input)
			outputs[len(outputs)-1] = outputs[len(outputs)-1] + i

			if err != nil {
				panic(err)
			}
		} else {
			outputs = append(outputs, 0)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return outputs
}
