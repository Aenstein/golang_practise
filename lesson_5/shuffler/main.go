package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

type Question struct {
	Country string `json:"country"`
	Capital string `json:"capital"`
}

func shuffle(slice interface{}) {
	rv := reflect.ValueOf(slice)

	if rv.Kind() != reflect.Slice {
		fmt.Println("Provided interface is not a slice type")
		return
	}

	length := rv.Len()
	swap := reflect.Swapper(slice)
	rand.Shuffle(length, func(i, j int) {
		swap(i, j)
	})
}

func main() {
	questions := []Question{
		{Country: "France", Capital: "Paris"},
		{Country: "Germany", Capital: "Berlin"},
		{Country: "Italy", Capital: "Rome"},
	}

	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println("Before shuffle:", questions)
	shuffle(questions)
	fmt.Println("After shuffle:", questions)

	fmt.Println("Before shuffle:", numbers)
	shuffle(numbers)
	fmt.Println("After shuffle:", numbers)

	shuffle(42)
}
