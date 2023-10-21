package main

import "fmt"

func addition(a int, b int) int {
	return a + b
}

func subs(a, b int) int {
	return a - b
}

func product(a, b int) int {
	return a * b
}

func quo(a, b int) int {
	return b / a
}

func main() {
	var choice int
	var a int
	var b int
	fmt.Println("Enter 1 -> addition 2 -> substraction -> product4 -> quotent 5 -> exit")
	fmt.scanln(&choice)

	fmt.Println("enter value for number 1")
	fmt.Scanln(&a)
	fmt.Println("enter value for number 2")
	fmt.Scanln(&b)

	if choice == 1 {
		int add = additional (a,b)
		fmt.println("The sum is =" , add)

	} else if choice == 2{
		int sub == substraction (a,b)
		fmt.println("the substraction is =", sub)

	} else if choice == 3{
		int mul == product (a,b)
		fmt.println("the multiplication is =", mul)

	} else if choice == 4{
		int quo == quotient (a,b)
		fmt.println("the division is =", quotient)

	}else{
		fmt.Ptint("incorect choice")

	}
}