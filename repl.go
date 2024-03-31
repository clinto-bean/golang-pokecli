package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)



var scanner = bufio.NewScanner(os.Stdin)

func startRepl(cfg *config) {
	
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := parseInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		
		command, exists := getCommands()[words[0]]
		if exists {
			err := command.callback(cfg, strings.Join(words[1:], " "))
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("unknown command")
			continue
		}
	}
}



