package parser

import (
	"fmt"
	"jsonparser/pkg/lexer"
	"jsonparser/pkg/token"
	"strconv"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []string
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	p.nextToken() // Initialize curToken
	p.nextToken() // Initialize peekToken
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) addError(msg string) {
	p.errors = append(p.errors, msg)
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.addError(msg)
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) currTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) parseValue() Value {
	fmt.Printf("Parsing value: %s\n", p.curToken.Literal) // Debug statement
	switch p.curToken.Type {
	case token.STRING:
		return &StringLiteral{Token: p.curToken.Literal, Value: p.curToken.Literal}
	case token.INT:
		value, _ := strconv.ParseInt(p.curToken.Literal, 0, 64)
		return &IntegerLiteral{Token: p.curToken.Literal, Value: value}
	case token.TRUE:
		return &BooleanLiteral{Token: p.curToken.Literal, Value: true}
	case token.FALSE:
		return &BooleanLiteral{Token: p.curToken.Literal, Value: false}
	case token.NULL:
		return &NullLiteral{Token: p.curToken.Literal}
	case token.LBRACE:
		return p.parseObject()
	case token.LBRACKET:
		return p.parseArray()
	default:
		p.addError(fmt.Sprintf("unexpected token %s", p.curToken.Type))
		return nil
	}
}

func (p *Parser) parseObject() *Object {
	obj := &Object{Pairs: make(map[string]Value)}
	fmt.Println("Parsing object") // Debug statement

	if !p.currTokenIs(token.LBRACE) {
		p.addError(fmt.Sprintf("Expected {, got %s", p.curToken.Type))
		return nil
	}

	p.nextToken() // Move past '{'

	for !p.currTokenIs(token.RBRACE) {
		if p.curToken.Type != token.STRING {
			p.addError(fmt.Sprintf("Expected string key, got %s", p.curToken.Type))
			return nil
		}

		key := p.curToken.Literal

		if !p.expectPeek(token.COLON) {
			return nil
		}

		p.nextToken() // Move past ':'
		value := p.parseValue()
		if value == nil {
			p.addError(fmt.Sprintf("Unexpected nil value for key %s", key))
			return nil
		}
		obj.Pairs[key] = value

		if p.peekTokenIs(token.RBRACE) {
			p.nextToken()
			break
		} else if p.peekTokenIs(token.COMMA) {
			p.nextToken()
		} else {
			p.addError(fmt.Sprintf("Expected , or }, got %s", p.peekToken.Type))
			return nil
		}

		p.nextToken() // Move to the next key or }
	}

	return obj
}

func (p *Parser) parseArray() *Array {
	arr := &Array{}
	fmt.Println("Parsing array") // Debug statement

	if !p.currTokenIs(token.LBRACKET) {
		p.addError(fmt.Sprintf("Expected [, got %s", p.curToken.Type))
		return nil
	}

	p.nextToken() // Move past '['

	for !p.currTokenIs(token.RBRACKET) {
		value := p.parseValue()
		if value == nil {
			p.addError("Unexpected nil value in array")
			return nil
		}
		arr.Elements = append(arr.Elements, value)

		if p.peekTokenIs(token.RBRACKET) {
			p.nextToken()
			break
		} else if p.peekTokenIs(token.COMMA) {
			p.nextToken()
		} else {
			p.addError(fmt.Sprintf("Expected , or ], got %s", p.peekToken.Type))
			return nil
		}

		p.nextToken() // Move to the next value or ]
	}

	return arr
}

func (p *Parser) ParseJson() Node {
	fmt.Println("Starting parse") // Debug statement
	switch p.curToken.Type {
	case token.LBRACE:
		return p.parseObject()
	case token.LBRACKET:
		return p.parseArray()
	default:
		p.addError(fmt.Sprintf("Expected { or [, got %s", p.curToken.Type))
		return nil
	}
}
