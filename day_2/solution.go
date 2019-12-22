package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func readInput(path string, separator string) []string {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	var input = strings.Split(string(content), separator)

	return input
}

func parseStringMapToIntMap(input []string) []int {
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

	return arrayOfNums
}

const (
	AddOp      = 1
	MultiplyOp = 2
	HaltOp     = 99
)

func main() {
	var input = parseStringMapToIntMap(
		readInput("input.txt", ","),
	)

	input[1] = 12
	input[2] = 2

	var length = len(input)

	var idx = 0
	for {
		if input[idx] == HaltOp {
			fmt.Printf("Halting program on %d (idx)\n", idx)
			break
		}

		if idx > length {
			fmt.Printf("Finishing program because %d (idx) is out of array length\n", idx)
			break
		}

		var optEl = input[idx]

		var firstOperandEl = input[input[idx+1]]
		var secondOperandEl = input[input[idx+2]]
		var replacePosition = input[idx+3]

		var resultOfOperation = 0

		switch optEl {
		case AddOp:
			fmt.Printf("Adding operation for %d, %d, to %d\n", firstOperandEl, secondOperandEl, replacePosition)
			resultOfOperation = firstOperandEl + secondOperandEl
			break
		case MultiplyOp:
			fmt.Printf("Multiply operation for %d, %d, to %d\n", firstOperandEl, secondOperandEl, replacePosition)
			resultOfOperation = firstOperandEl * secondOperandEl
			break
		default:
			log.Fatal("Default case")
		}

		input[replacePosition] = resultOfOperation

		idx += 4
	}

	fmt.Printf("Result for given input is %d\n", input[0])
}
