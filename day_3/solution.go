package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
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

const (
	Right = 'R'
	Left  = 'L'
	Up    = 'U'
	Down  = 'D'
)

func parseDirectionStringToWiresMap(directionString string) [][][]int {
	var rawWireMap = strings.Split(string(directionString), ",")

	var wirePartsLength = len(rawWireMap)
	var wiresMap = make([][][]int, wirePartsLength)

	previousPoint := []int{0, 0}

	for idx, value := range rawWireMap {
		var dirEl = value[0]
		var posEl = value[1:]

		num, err := strconv.Atoi(posEl)

		if err != nil {
			log.Fatal(err)
		}

		currentPoint := []int{0, 0}
		copy(currentPoint, previousPoint)

		switch dirEl {
		case Right:
			currentPoint[0] += num
		case Left:
			currentPoint[0] -= num
		case Up:
			currentPoint[1] += num
		case Down:
			currentPoint[1] -= num
		default:
			log.Fatal("Default case")
		}

		pPoint := []int{0, 0}
		cPoint := []int{0, 0}

		copy(pPoint, previousPoint)
		copy(cPoint, currentPoint)
		copy(previousPoint, currentPoint)

		line := [][]int{pPoint, cPoint}

		wiresMap[idx] = line
	}

	return wiresMap
}

func solution(firstInput [][][]int, secondInput [][][]int) (int, int) {
	var fLength = len(firstInput)
	var sLength = len(secondInput)

	var intersectionsPoints [][]int
	var vectorMovements [][]int

	// Note: Get intersection points
	for i := 0; i < fLength; i++ {
		var firstEl = firstInput[i]
		var firstHorizontal = checkIfHorizontal(firstEl)

		fmt.Printf("firstEl %v\n", firstEl)

		fXes := []int{firstEl[0][0], firstEl[1][0]}
		fYes := []int{firstEl[0][1], firstEl[1][1]}

		for j := 0; j < sLength; j++ {
			var secondEl = secondInput[j]
			var secondHorizontal = checkIfHorizontal(secondEl)

			fmt.Printf("secondEl %v\n", secondEl)

			if firstHorizontal == secondHorizontal {
				continue
			}

			sXes := []int{secondEl[0][0], secondEl[1][0]}
			sYes := []int{secondEl[0][1], secondEl[1][1]}

			var x = secondEl[0][0]
			var y = secondEl[0][1]

			var intersectionPoint []int

			if firstHorizontal == true && pointInRange(fXes, x) && pointInRange(sYes, fYes[0]) {
				intersectionPoint = []int{x, fYes[0]}
			}

			if firstHorizontal == false && pointInRange(fYes, y) && pointInRange(sXes, fXes[0]) {
				intersectionPoint = []int{fXes[0], y}
			}

			// Note: Lack of intersection point
			if len(intersectionPoint) == 0 {
				fmt.Printf("No intersection point...\n")
				continue
			}

			// Note: Ignore 0,0 point (central)
			if intersectionPoint[0] == 0 && intersectionPoint[1] == 0 {
				fmt.Printf("Central point... ignoring\n")
				continue
			}

			vectorMovement := []int{i, j}
			fmt.Printf("vectorMovement = %v\n", vectorMovement)
			fmt.Printf("intersectionPoint = %v\n", intersectionPoint)

			intersectionsPoints = append(intersectionsPoints, intersectionPoint)
			vectorMovements = append(vectorMovements, vectorMovement)
		}
	}

	fmt.Printf("intersectionsPoints %v\n", intersectionsPoints)

	var iLength = len(intersectionsPoints)
	var distances []int

	// Note: Calculate Manhattan distance for each intersection point
	for i := 0; i < iLength; i++ {
		var el = intersectionsPoints[i]
		var distance = manhattanDist(el)
		distances = append(distances, distance)
	}

	var smallestDistance = min(distances)

	// Note: Calculate vector transformation steps sum for each intersection point
	fmt.Printf("vectorMovements %v\n", vectorMovements)
	var vectorSteps []int

	for i := 0; i < iLength; i++ {
		var intersectionEl = intersectionsPoints[i]
		var fVectorMovements = vectorMovements[i][0]
		var sVectorMovemenets = vectorMovements[i][1]

		var transformation = 0

		transformation += countVectorSteps(firstInput, intersectionEl, fVectorMovements) +
			countVectorSteps(secondInput, intersectionEl, sVectorMovemenets)

		vectorSteps = append(vectorSteps, transformation)
	}

	var smallestVectorStep = min(vectorSteps)

	return smallestDistance, smallestVectorStep
}

func checkIfHorizontal(el [][]int) bool {
	var isHorizontal = true

	if el[0][0] == el[1][0] {
		isHorizontal = false
	}

	return isHorizontal
}

func pointInRange(el []int, point int) bool {
	var x0 = el[0]
	var x1 = el[1]

	var tempX0 = x0

	if x0 > x1 {
		x0 = x1
		x1 = tempX0
	}

	return x0 <= point && point <= x1
}

func min(array []int) int {
	var min = array[0]

	for _, value := range array {
		if value < min {
			min = value
		}
	}

	return min
}

func countVectorSteps(vectors [][][]int, intersectionPoint []int, movements int) int {
	var transformationSum = 0

	for i := 0; i < movements; i++ {
		var el = vectors[i]

		if i+1 == movements {
			var xVector = (int)(math.Abs(float64(intersectionPoint[0] - el[0][0])))
			var yVector = (int)(math.Abs(float64(intersectionPoint[1] - el[0][1])))

			transformationSum += xVector + yVector
			break
		}

		var xVector = (int)(math.Abs(float64(el[1][0] - el[0][0])))
		var yVector = (int)(math.Abs(float64(el[1][1] - el[0][1])))

		transformationSum += xVector + yVector
	}

	return transformationSum
}

func manhattanDist(point []int) int {
	return (int)(math.Abs(float64(point[0])) + math.Abs(float64(point[1])))
}

func main() {
	input := readInput("input.txt", "\n")

	wire1 := parseDirectionStringToWiresMap(input[0])
	wire2 := parseDirectionStringToWiresMap(input[1])

	smallestDistance, smallestVectorStep := solution(wire1, wire2)

	fmt.Printf("Result for first part for given input is %d\n", smallestDistance)
	fmt.Printf("Result for second part for given input is %d\n", smallestVectorStep)
}
