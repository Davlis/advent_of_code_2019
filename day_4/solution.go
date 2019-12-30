package main

import (
	"fmt"
	"log"
	"math"
)

// Note: Puzzle input
var startRange = 254032
var finishRange = 789860

func solution(startPoint int, endPoint int) (firstSum int, secondSum int) {

	// Constraint: The value is within the range given in puzzle (254032, 789860)
	for num := startPoint; num <= endPoint; num++ {

		fSkip := false

		// Constraint: It is a six-digit number
		if countDigits(num) != 6 {
			fmt.Printf("Number %d does not have six digits\n", num)
			continue
		}

		// Constraint: Going from left to right, the digits never decrease
		var previousDigit = -1
		for i := 1; i <= 6; i++ {
			currentDigit := getDigit(num, i)

			if previousDigit > currentDigit {
				fSkip = true
				break
			}

			previousDigit = currentDigit
		}

		if fSkip {
			fmt.Printf("Digits in %d number decreases\n", num)
			continue
		}

		m := make(map[int]int)

		for i := 1; i <= 6; i++ {
			digit := getDigit(num, i)

			m[digit]++
		}

		fSkip = true
		sSkip := true

		// Constraint: Two adjacent digits are the same
		// Constraint: The two adjacent matching digits are not part of a larger group of matching digits
		for _, v := range m {
			if v >= 2 {
				fSkip = false
			}

			if v == 2 {
				sSkip = false
			}
		}

		if fSkip {
			continue
		}
		firstSum++

		if sSkip {
			continue
		}
		secondSum++
	}

	return firstSum, secondSum
}

func getDigit(number int, n int) int {
	if n < 1 {
		log.Fatal("N argument has to be >= 1")
	}

	digitCount := countDigits(number)

	if n > digitCount {
		log.Fatal("N argument has to be lower")
	}

	return number / (int)(math.Pow10(digitCount-n)) % 10
}

func countDigits(i int) (count int) {
	for i != 0 {

		i /= 10
		count = count + 1
	}
	return count
}

func main() {
	firstCount, secondCound := solution(startRange, finishRange)

	fmt.Printf("Result for first part for given input is %d\n", firstCount)
	fmt.Printf("Result for second part for given input is %d\n", secondCound)
}
