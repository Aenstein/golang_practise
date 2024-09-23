package main

import (
	"fmt"
	"math/big"
	"strings"
)

func processHex(hexStr string) string {
	return strings.TrimPrefix(hexStr, "0x")
}

func main() {
	var input string
	i := new(big.Int)

	fmt.Println("Введите шеснадцатеричное число:")
	fmt.Scanln(&input)

	input = strings.ToLower(input)

	if _, ok := i.SetString(processHex(input), 16); !ok {
		fmt.Println("Невалидное число")
	} else {
		fmt.Println(i)
	}
}
