package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	// infinite loop for REPL (read-eval-print loop)
	for {
		fmt.Print("Pokedex > ")
		
		if !scanner.Scan() {	// read til the next /n and return true or fakse
			break
		}

		userInput := scanner.Text()
		words := cleanInput(userInput)
		if len(words) == 0 {	// if it is empty go to the next interation
			continue
		}

		command := words[0]	// the command is the first word of the input
		fmt.Printf("Your command was: %s\n", command)

	}

}
