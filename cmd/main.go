package main

import (
	"fmt"
	"jsonparser/pkg/lexer"
	"jsonparser/pkg/parser"
	"jsonparser/utils"
	"os"
)

func main() {
	inputFile := "internal/sample.json"

	input, err := os.ReadFile(inputFile)

	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		os.Exit(1)
	}

	l := lexer.NewLexer(string(input))
	p := parser.NewParser(l)

	parsedJson := p.ParseJson()

	if parsedJson == nil {
		fmt.Println("Parsing failed with errors:")
		for _, err := range p.Errors() {
			fmt.Println(err)
		}
		return
	}

	// Print the parsed JSON
	utils.PrintJson(parsedJson)

}
