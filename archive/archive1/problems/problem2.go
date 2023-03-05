//Even Fibonacci numbers

package problems

import "fmt"

func EvenFibonacciNumbers(number int) {
	fmt.Printf("Probelm 2: Even Fibonacci numbers\n\n")

	var fibNumbers []int
	// var evenFibNumbers []int

	a := 0
	b := 1
	c := b

	for true {
		slice := 0

		b = a + b
		if b >= number {
			fmt.Println()
			break
		}
		a = c
		fmt.Printf(" %d", b)
		fibNumbers = append(fibNumbers, b)
		slice++
		fmt.Print(fibNumbers)
	}
}
