package main

import "fmt"

func main() {
	var n int
	fmt.Println("Insira um número: ")
	fmt.Scanln(&n)
	if n%2 == 0 {
		fmt.Println(n, " é par")
	} else {
		fmt.Println(n, " é impar")
	}
}
