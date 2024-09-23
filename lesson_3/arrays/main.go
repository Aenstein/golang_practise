package main

func main() {
	var a [3]int

	a[0] = 5
	a[1] = 4
	a[2] = 5

	for index, element := range a {
		println(index, element)
	}
}
