package main

import (
	"fmt"
	"os"
)

func main() {
	var n int

	fmt.Println("Enter natural number:")
	fmt.Scanln(&n)

	if n < 1 {
		fmt.Println("Number not natural")
		os.Exit(1)
	}
	
	numbers := []int{}
	for i := 1; i <= n; i++ {
		numbers = append(numbers, i)
	}

	fmt.Print("a\\b\t")

	for _, a := range numbers {
		fmt.Printf("%8d", a)
	}

	fmt.Printf("\n\n")

	for _, a := range numbers {
		fmt.Printf("%d\t", a)

		for _, b := range numbers {
			fmt.Printf("%8d", a*b)
		}
		fmt.Printf("\n")
	}
}
