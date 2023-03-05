// Multiples of 3 or 5

package problems

import "fmt"

func IsItMultiple() {
	fmt.Print("Problem 1: Multiples of 3 or 5\n")

	var threeMultipleSum int
	var fiveMultipleSum int
	var fiftheenMultipleSum int

	currentNumber := 0

	for currentNumber != 1000 {

		if currentNumber%15 == 0 {
			fiftheenMultipleSum += currentNumber
		} else {
			//if it isnt 15 what is it
			if currentNumber%3 == 0 {
				threeMultipleSum += currentNumber
			}

			if currentNumber%5 == 0 {
				fiveMultipleSum += currentNumber
			}
		}

		//next number
		currentNumber += 1

	}

	var sumOfMultiples = threeMultipleSum + fiveMultipleSum + fiftheenMultipleSum

	fmt.Printf("Sum of multiples of 3: %d\n", threeMultipleSum)
	fmt.Printf("Sum of multiples of 5: %d\n", fiveMultipleSum)

	fmt.Printf("Sum of multiples for 3 and 5 under %d: %d\n", currentNumber, sumOfMultiples)

}
