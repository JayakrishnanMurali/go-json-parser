package parser

type Node interface {
	TokenLiteral() string
}

type Value interface {
	Node
	valueNode()
}

type StringLiteral struct {
	Value string
	Token string
}

func (sl *StringLiteral) TokenLiteral() string {
	return sl.Token
}

func (sl *StringLiteral) valueNode() {}

type IntegerLiteral struct {
	Value int64
	Token string
}

func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token
}

func (il *IntegerLiteral) valueNode() {}

type BooleanLiteral struct {
	Value bool
	Token string
}

func (bl *BooleanLiteral) TokenLiteral() string {
	return bl.Token
}

func (bl *BooleanLiteral) valueNode() {}

type NullLiteral struct {
	Token string
}

func (nl *NullLiteral) TokenLiteral() string {
	return nl.Token
}

func (nl *NullLiteral) valueNode() {}

type Object struct {
	Pairs map[string]Value
}

func (o *Object) TokenLiteral() string {
	return "{"
}

func (o *Object) valueNode() {}

type Array struct {
	Elements []Value
}

func (a *Array) TokenLiteral() string {
	return "["
}

func (a *Array) valueNode() {}
