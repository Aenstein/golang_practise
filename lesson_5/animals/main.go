package main

import (
	"animals/pets"
	"fmt"
	"os"
)

func main() {
	myCat := pets.Cat{
		Animal: pets.Animal{
			Name:    "mr. button",
			Age:     5,
			Weight:  7,
			IsSleep: true,
		},
	}
	myDog := pets.Dog{
		Animal: pets.Animal{
			Name: "spot",
			Age: 8,
			Weight: 30,
			IsSleep: false,
		},
	}

	var feedToCat uint8 = 3
	var feedToDog uint8 = 10

	catFed, err := feed(&myCat, feedToCat)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error feeding cat: %v\n", err)
	} else {
		fmt.Println("Cat ate:", catFed)
	}

	fmt.Print("\n\n\t =====\n\n\n")

	dogFed, err := feed(&myDog, feedToDog)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error feeding dog: %v\n", err)
	} else {
		fmt.Println("Dog ate:", dogFed)
	}

	fmt.Print("\n\n\t =====\n\n\n")
}

func feed(animal pets.EaterWalker, amount uint8) (uint8, error) {
	fmt.Println(animal.Walk())

	fmt.Printf("Now let's feed our %s\n", animal.GetName())
	return animal.Eat(amount)
}
