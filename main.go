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

		commandName := words[0]
		cmd, ok := commands[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		if err := cmd.callback(); err != nil {
			fmt.Println(err)
		}
	}

}
