package main

import "fmt"

func main() {
	var first, second int

	_, err := fmt.Scanln(&first)
	if err != nil{
		fmt.Println("Invalid first operand")
		return
	}

	_, err = fmt.Scanln(&second)
	if err != nil{
		fmt.Println("Invalid second operand")
		return
	}

	fmt.Println(first, second)
}
