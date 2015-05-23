package parser

import (
	"errors"
	"github.com/tboronczyk/kiwi/ast"
	"github.com/tboronczyk/kiwi/scanner"
	"github.com/tboronczyk/kiwi/token"
)

type Parser struct {
	token   token.Token
	value   string
	scanner scanner.Scanner
}

func NewParser() *Parser {
	return &Parser{}
}

func (p Parser) match(tkn token.Token) bool {
	return p.token == tkn
}

func (p *Parser) advance() {
	for {
		p.token, p.value = p.scanner.Scan()
		if p.token != token.COMMENT {
			break
		}
	}
}

func (p *Parser) consume(t token.Token) {
	if p.token != t {
		panic(p.expected(t))
	}
	p.advance()
}

func (p Parser) expected(value interface{}) string {
	var str string
	switch value.(type) {
	case token.Token:
		str = value.(token.Token).String()
	default:
		str = value.(string)
	}
	return "Expected " + str + " but saw " + p.token.String()
}

func (p *Parser) InitScanner(scnr scanner.Scanner) {
	p.scanner = scnr
	p.advance()
}

func (p *Parser) Parse() (node ast.Node, err error) {
	if p.token == token.EOF {
		return nil, nil
	}
	defer func() {
		if e := recover(); e != nil {
			node = nil
			err = errors.New(e.(string))
		}
	}()
	return p.stmt(), nil
}

func (p *Parser) expr() ast.Node {
	n := p.relation()
	if !p.token.IsLogOp() {
		return n
	}

	node := ast.Operator{Op: p.token, Left: n}
	p.advance()
	node.Right = p.expr()

	return node
}

func (p *Parser) relation() ast.Node {
	n := p.simpleExpr()
	if !p.token.IsCmpOp() {
		return n
	}

	node := ast.Operator{Op: p.token, Left: n}
	p.advance()
	node.Right = p.relation()

	return node
}

func (p *Parser) simpleExpr() ast.Node {
	n := p.term()
	if !p.token.IsAddOp() {
		return n
	}

	node := ast.Operator{Op: p.token, Left: n}
	p.advance()
	node.Right = p.simpleExpr()

	return node
}

func (p *Parser) term() ast.Node {
	n := p.factor()
	if !p.token.IsMulOp() {
		return n
	}

	node := ast.Operator{Op: p.token, Left: n}
	p.advance()
	node.Right = p.term()

	return node
}

func (p *Parser) factor() ast.Node {
	if p.match(token.LPAREN) {
		defer p.consume(token.RPAREN)
		p.advance()
		return p.expr()
	}
	if p.token.IsExprOp() {
		node := ast.Operator{Op: p.token}
		p.advance()
		node.Left = p.factor()
		return node
	}
	return p.terminal()
}

func (p *Parser) terminal() ast.Node {
	if p.token == token.TRUE || p.token == token.FALSE ||
		p.token == token.NUMBER || p.token == token.STRING {
		defer p.advance()
		return ast.Literal{Type: p.token, Value: p.value}
	}

	node := p.identifier()
	if p.token != token.LPAREN {
		return node
	}
	return ast.FuncCall{Name: node, Args: p.parenExprList()}
}

func (p *Parser) parenExprList() ast.Node {
	p.consume(token.LPAREN)

	if p.token == token.RPAREN {
		p.consume(token.RPAREN)
		return nil
	}

	node := p.exprList()

	p.consume(token.RPAREN)
	return node
}

func (p *Parser) exprList() ast.Node {
	n := p.expr()
	if p.token != token.COMMA {
		return n
	}

	node := ast.List{Node: n}
	for p.token == token.COMMA {
		p.advance()
		node = ast.List{Node: p.expr(), Prev: node}
	}
	return node
}

func (p *Parser) identList() ast.Node {
	n := p.identifier()
	if p.token != token.COMMA {
		return n
	}

	node := ast.List{Node: n}
	for p.token == token.COMMA {
		p.advance()
		node = ast.List{Node: p.identifier(), Prev: node}
	}
	return node
}

func (p *Parser) stmt() ast.Node {
	switch p.token {
	case token.FUNC:
		return p.funcDefStmt()
	case token.IF:
		return p.ifStmt()
	case token.RETURN:
		return p.returnStmt()
	case token.WHILE:
		return p.whileStmt()
	case token.IDENTIFIER:
		return p.assignOrCallStmt()
	}
	panic(p.expected("statement"))
}

func (p *Parser) funcDefStmt() ast.FuncDef {
	p.consume(token.FUNC)
	node := ast.FuncDef{Name: p.identifier()}
	if p.token != token.LBRACE {
		node.Params = p.identList()
	}
	node.Body = p.braceStmtList()
	return node
}

func (p *Parser) ifStmt() ast.If {
	p.consume(token.IF)
	return ast.If{
		Condition: p.expr(),
		Body:      p.braceStmtList()}
}

func (p *Parser) returnStmt() ast.Return {
	defer p.consume(token.SEMICOLON)
	p.consume(token.RETURN)
	node := ast.Return{}
	if p.token != token.SEMICOLON {
		node.Expr = p.expr()
	}
	return node
}

func (p *Parser) braceStmtList() ast.List {
	defer p.consume(token.RBRACE)
	p.consume(token.LBRACE)
	return p.stmtList()
}

func (p *Parser) stmtList() ast.List {
	node := ast.List{}
	for p.token.IsStmtKeyword() || p.token == token.IDENTIFIER {
		node.Node = p.stmt()
		if p.token.IsStmtKeyword() || p.token == token.IDENTIFIER {
			node = ast.List{Prev: node}
		}
	}
	return node
}

func (p *Parser) whileStmt() ast.While {
	p.consume(token.WHILE)
	return ast.While{
		Condition: p.expr(),
		Body:      p.braceStmtList()}
}

func (p *Parser) assignOrCallStmt() ast.Node {
	defer p.consume(token.SEMICOLON)
	n := p.identifier()
	if p.token == token.ASSIGN {
		node := ast.Operator{Op: p.token, Left: n}
		p.advance()
		node.Right = p.expr()
		return node
	}
	if p.token == token.LPAREN {
		return ast.FuncCall{Name: n, Args: p.parenExprList()}
	}
	panic(p.expected(
		token.ASSIGN.String() + " or " + token.LPAREN.String()))
}

func (p *Parser) identifier() ast.Literal {
	defer p.consume(token.IDENTIFIER)
	return ast.Literal{Type: p.token, Value: p.value}
}
