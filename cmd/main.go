package main

import (
	"bufio"
	"fmt"
	"jsonparser/pkg/lexer"
	"jsonparser/pkg/parser"
	"jsonparser/utils"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Choose input method: ")
	fmt.Println("1. Type JSON")
	fmt.Println("2. Read from sample.json")
	fmt.Print("Enter choice (1 or 2): ")

	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	var input string

	switch choice {
	case "1":
		fmt.Println("Enter your JSON:")
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)
	case "2":
		inputFile := "internal/sample.json"
		fileBytes, err := os.ReadFile(inputFile)
		if err != nil {
			fmt.Printf("Error reading file: %s\n", err)
			os.Exit(1)
		}
		input = string(fileBytes)
	default:
		fmt.Println("Invalid choice. Exiting.")
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
