package utils

import (
	"fmt"
	"jsonparser/pkg/parser"
)

// PrintJson prints the parsed JSON node.
func PrintJson(node parser.Node) {
	switch v := node.(type) {
	case *parser.Object:
		fmt.Println("{")
		printObject(v, 1)
		fmt.Println("}")
	case *parser.Array:
		fmt.Println("[")
		printArray(v, 1)
		fmt.Println("]")
	default:
		fmt.Println("Parsed JSON as an unknown type")
	}
}

func printObject(obj *parser.Object, indent int) {
	for key, value := range obj.Pairs {
		printIndent(indent)
		fmt.Printf("%q: ", key)
		printValue(value, indent)
	}
}

func printArray(arr *parser.Array, indent int) {
	for i, value := range arr.Elements {
		printIndent(indent)
		fmt.Printf("[%d]: ", i)
		printValue(value, indent)
	}
}

func printValue(value parser.Value, indent int) {
	switch v := value.(type) {
	case *parser.StringLiteral:
		fmt.Printf("%q\n", v.Value)
	case *parser.IntegerLiteral:
		fmt.Printf("%d\n", v.Value)
	case *parser.BooleanLiteral:
		fmt.Printf("%t\n", v.Value)
	case *parser.NullLiteral:
		fmt.Println("null")
	case *parser.Object:
		fmt.Println("{")
		printObject(v, indent+1)
		printIndent(indent)
		fmt.Println("}")
	case *parser.Array:
		fmt.Println("[")
		printArray(v, indent+1)
		printIndent(indent)
		fmt.Println("]")
	default:
		fmt.Println("unknown type")
	}
}

func printIndent(indent int) {
	for i := 0; i < indent; i++ {
		fmt.Print("  ")
	}
}
