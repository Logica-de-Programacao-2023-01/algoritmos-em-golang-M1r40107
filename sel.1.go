package main

import "fmt"

func main() {
	var num1 int
	var num2 int
	fmt.Println("Insira 1 número: ")
	fmt.Scanln(&num1)
	fmt.Println("Insira outro número: ")
	fmt.Scanln(&num2)
	if num1 < num2 {
		fmt.Println(num2, " é maior")
	} else if num2 < num1 {
		fmt.Println(num1, "é maior")
	} else if num1 == num2 {
		fmt.Println("O primeiro número é igual ao segundo número")
	}
}
