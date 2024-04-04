package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to Pokedex!")
	for {
		fmt.Print("> ")
		var input string
		in := bufio.NewReader(os.Stdin)
		input, err := in.ReadString('\n')
		if err != nil {
			fmt.Println("Bad input... try again!")
			continue
		}
		inputItems := parseInputItems(input)
		err = executeCommand(inputItems)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
	}
}
