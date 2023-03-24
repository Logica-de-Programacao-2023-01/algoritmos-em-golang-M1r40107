package main

import "fmt"

func main() {
	arr := [6]float64{5.2, 4.7, 6.9, 7.5, 4.3, 1.8}
	var soma float64

	for _, x := range arr {
		soma += x
	}
	media := soma / float64(len(arr))
	fmt.Println("A média é", media)
}
