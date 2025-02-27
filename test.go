package main

import "fmt"

func main2() {
	options := []string{"apple", "banana", "cherry"}

	fmt.Println("Choose an option:")
	for i, opt := range options {
		fmt.Printf("%d: %s\n", i+1, opt)
	}

	var choice int
	fmt.Print("Enter the number of your choice: ")
	_, err := fmt.Scanln(&choice)
	if err != nil || choice < 1 || choice > len(options) {
		fmt.Println("Invalid choice. Please run the program again and select a valid option.")
		return
	}

	fmt.Println("You chose:", options[choice-1])
}
