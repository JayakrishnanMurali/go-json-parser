package utils

import (
	"fmt"
	"jsonparser/pkg/parser"
)

// PrintJson prints the parsed JSON node.
func PrintJson(node parser.Node) {
	switch v := node.(type) {
	case *parser.Object:
		printObject(v, 1)
	case *parser.Array:
		printArray(v, 1)
	default:
		fmt.Println("Parsed JSON as an unknown type")
	}
}

func printObject(obj *parser.Object, indent int) {
	for key, value := range obj.Pairs {
		printIndent(indent)
		fmt.Printf("%s: ", key)
		printValue(value, indent+1)
	}
}

func printArray(arr *parser.Array, indent int) {
	for i, value := range arr.Elements {
		printIndent(indent)
		fmt.Printf("[%d]: ", i)
		printValue(value, indent+1)
	}
}

func printValue(value parser.Value, indent int) {
	switch v := value.(type) {
	case *parser.StringLiteral:
		fmt.Printf("String: %s\n", v.Value)
	case *parser.IntegerLiteral:
		fmt.Printf("Integer: %d\n", v.Value)
	case *parser.BooleanLiteral:
		fmt.Printf("Boolean: %t\n", v.Value)
	case *parser.NullLiteral:
		fmt.Println("Null")
	case *parser.Object:
		fmt.Println("Object:")
		printObject(v, indent+1)
	case *parser.Array:
		fmt.Println("Array:")
		printArray(v, indent+1)
	default:
		fmt.Printf("Unknown type\n")
	}
}

func printIndent(indent int) {
	for i := 0; i < indent; i++ {
		fmt.Print("  ")
	}
}
