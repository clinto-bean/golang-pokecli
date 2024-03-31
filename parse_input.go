package main

import "strings"

func parseInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}