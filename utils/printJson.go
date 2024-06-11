package utils

import (
	"fmt"
	"jsonparser/pkg/parser"
)

const (
	Reset       = "\033[0m"
	Red         = "\033[31m"
	Green       = "\033[32m"
	Yellow      = "\033[33m"
	Blue        = "\033[34m"
	Purple      = "\033[35m"
	Cyan        = "\033[36m"
	White       = "\033[37m"
	BrightBlack = "\033[90m"
	BrightBlue  = "\033[94m"
	BrightCyan  = "\033[96m"
	BrightWhite = "\033[97m"
)

// PrintJson prints the parsed JSON node as a hierarchical tree.
func PrintJson(node parser.Node) {
	switch v := node.(type) {
	case *parser.Object:
		printObject(v, 0, "")
	case *parser.Array:
		printArray(v, 0, "")
	default:
		fmt.Println(Red + "Parsed JSON as an unknown type" + Reset)
	}
}

func printObject(obj *parser.Object, indent int, prefix string) {
	for key, value := range obj.Pairs {
		printIndent(indent, prefix, true)
		fmt.Printf("%s%s%s:\n", BrightBlue, key, Reset)
		printValue(value, indent+1, prefix+"|  ")
	}
}

func printArray(arr *parser.Array, indent int, prefix string) {
	for i, value := range arr.Elements {
		printIndent(indent, prefix, true)
		fmt.Printf("%s[%d]%s:\n", BrightCyan, i, Reset)
		printValue(value, indent+1, prefix+"|  ")
	}
}

func printValue(value parser.Value, indent int, prefix string) {
	switch v := value.(type) {
	case *parser.StringLiteral:
		printIndent(indent, prefix, false)
		fmt.Printf("%sString: %s%s%s\n", Green, Yellow, v.Value, Reset)
	case *parser.IntegerLiteral:
		printIndent(indent, prefix, false)
		fmt.Printf("%sInteger: %s%d%s\n", Green, Yellow, v.Value, Reset)
	case *parser.BooleanLiteral:
		printIndent(indent, prefix, false)
		fmt.Printf("%sBoolean: %s%t%s\n", Green, Yellow, v.Value, Reset)
	case *parser.NullLiteral:
		printIndent(indent, prefix, false)
		fmt.Printf("%sNull%s\n", Green, Reset)
	case *parser.Object:
		printIndent(indent, prefix, false)
		fmt.Printf("%sObject:%s\n", Purple, Reset)
		printObject(v, indent+1, prefix+"|  ")
	case *parser.Array:
		printIndent(indent, prefix, false)
		fmt.Printf("%sArray:%s\n", Purple, Reset)
		printArray(v, indent+1, prefix+"|  ")
	default:
		printIndent(indent, prefix, false)
		fmt.Println(Red + "Unknown type" + Reset)
	}
}

func printIndent(_ int, prefix string, isKey bool) {
	fmt.Print(prefix)
	if isKey {
		fmt.Print("+--")
	} else {
		fmt.Print("|  ")
	}
}
