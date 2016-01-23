package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	parser := &Parser{}

	// Infinite loop for the user input
	for {
		input := bufio.NewReader(os.Stdin)
		fmt.Printf("> ")
		inputString, err := input.ReadString('\n')
		if err != nil {
			fmt.Printf("Something went wrong: %v\n", err)
			break
		}

		parser.ParseCurrentInput(inputString)
		if parser.IsAction {
			if parser.IsExit {
				fmt.Println("See ya!")
				break
			}

		} else {
			fmt.Printf("\n Well hello %s\n", parser.GetCurrentInput())
		}
	}
}
