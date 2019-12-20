package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"strconv"
)

func readInput(path string) []string {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	var input = strings.Split(string(content), "\n")

	return input
}

func main() {
	input := readInput("input.txt")

	var length = len(input)

	var arrayOfNums = make([]int, length)

	for i := 0; i < length; i++ {
		var el = input[i]

		num, err := strconv.Atoi(el)

		if err != nil {
			log.Fatal(err)
		}

		arrayOfNums[i] = num
	}

	var mass = 0

	for i := 0; i < length; i++ {
		var el = arrayOfNums[i]
		mass += (el / 3) - 2
	}

	fmt.Printf("Result for given input is %d\n", mass)
}
