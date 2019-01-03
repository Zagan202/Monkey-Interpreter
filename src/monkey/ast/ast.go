package ast

import (
	"bytes"
	"monkey/token"
)

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

type Node interface {
	/*
		A TokenLiteral() method that returns
		the literal value of the token it’s associated with. TokenLiteral() will be used only for debugging
		and testing.
	*/
	TokenLiteral() string
	/*
		String() method to AST nodes.
		This will allow us
		to print AST nodes for debugging and to compare them with other AST nodes.
		This is going to be really handy in tests!
	*/
	String() string
}

/* Statement ...
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

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

/*
The statementNode and TokenLiteral methods are there to fulfill the Node and
Statement interfaces
*/
func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

func (ls *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

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

func (i *Identifier) String() string { return i.Value }

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

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

type ExpressionStatement struct {
	Token      token.Token //first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
