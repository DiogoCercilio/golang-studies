package main

import "fmt"

func op() {

	var n1 int
	var n2 int
	var operation string
	var res int

	fmt.Println("Enter a number:")
	fmt.Scan(&n1)
	fmt.Println("Enter other number:")
	fmt.Scan(&n2)
	fmt.Println("Enter an operation (+, -, * or /):")
	fmt.Scan(&operation)

	switch operation {
	case "+":
		res = n1 + n2
	case "-":
		res = n1 - n2
	case "*":
		res = n1 * n2
	case "/":
		res = n1 / n2
	}

	fmt.Println("Result:", res)
}

func main() {
	var names []string

	names = append(names, "Camila")
	names = append(names, "Diogo")
	names = append(names, "Heloisa")

	fmt.Println(names[len(names)-1])

	for {
		op()
	}
}
