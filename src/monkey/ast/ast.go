package ast

import "monkey/token"

type Node interface {
	TokenLiteral() string
	/*
		A TokenLiteral() method that returns
		the literal value of the token it’s associated with. TokenLiteral() will be used only for debugging
		and testing.
	*/
}

/* Statement
let <identifier> = <expression>;
Statements do not produce values*/
type Statement interface {
	Node
	statementNode()
}

/*
Expression produce values, statements don’t. let x = 5 doesn’t produce a value,
whereas 5 does (the value it produces is 5).
*/
type Expression interface {
	Node
	expressionNode()
}

//Program node is going to be the root node of every AST our parser produces
type Program struct {
	Statements []Statement
}

type LetStatement struct {
	Token token.Token //the token.LET token
	Name  *Identifier
	Value Expression
}

/*
The statementNode and TokenLiteral methods are there to fulfill the Node and
Statement interfaces
*/
func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token
	Value string
}

/*
The expressionNode and TokenLiteral methods are there to fulfill the Node and
Statement interfaces
*/
func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

//ReturnStatement ...
//return <expression>;
type ReturnStatement struct {
	Token       token.Token //the 'return' token
	ReturnValue Expression
}

/*
The statementNode and TokenLiteral methods are there to fulfill the Node and
Statement interfaces
*/
func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
