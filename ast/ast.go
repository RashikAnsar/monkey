package ast

import "monkey/token"

// Node represents each node in AST
type Node interface {
	TokenLiteral() string
}

// Statement Node in the AST
type Statement interface {
	Node
	statementNode()
}

// Expression Node in the AST
type Expression interface {
	Node
	expressionNode()
}

// Program node is going to be the Root node of every AST this parser produces
type Program struct {
	Statements []Statement
}

// ReturnStatement ast structure of return statement
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

// TokenLiteral return the literal value of the token its associated with
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// LetStatement Parsing
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
