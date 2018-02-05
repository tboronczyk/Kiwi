package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenToString(t *testing.T) {
	tokens := []struct{ actual, expected string }{
		{UNKNOWN.String(), "UNKNOWN"},
		{EOF.String(), "EOF"},
		{ADD.String(), "+"},
		{SUBTRACT.String(), "-"},
		{MULTIPLY.String(), "*"},
		{DIVIDE.String(), "/"},
		{MODULO.String(), "%"},
		{EQUAL.String(), "="},
		{NOT_EQUAL.String(), "~="},
		{LESS.String(), "<"},
		{LESS_EQ.String(), "<="},
		{GREATER.String(), ">"},
		{GREATER_EQ.String(), ">="},
		{AND.String(), "&&"},
		{OR.String(), "||"},
		{NOT.String(), "~"},
		{IF.String(), "if"},
		{FUNC.String(), "func"},
		{RETURN.String(), "return"},
		{WHILE.String(), "while"},
		{BOOL.String(), "BOOL"},
		{IDENTIFIER.String(), "IDENTIFIER"},
		{NUMBER.String(), "NUMBER"},
		{STRING.String(), "STRING"},
		{ASSIGN.String(), ":="},
		{LBRACE.String(), "{"},
		{RBRACE.String(), "}"},
		{COLON.String(), ":"},
		{COMMA.String(), ","},
		{COMMENT.String(), "COMMENT"},
		{ELSE.String(), "else"},
		{LPAREN.String(), "("},
		{RPAREN.String(), ")"},
		{Token(254).String(), "Token(254)"},
	}

	for _, tkn := range tokens {
		assert.Equal(t, tkn.actual, tkn.expected)
	}
}

func TestIsAddOp(t *testing.T) {
	for i := 0; i < len(tokens); i++ {
		tkn := Token(i)
		if tkn == ADD || tkn == SUBTRACT {
			assert.True(t, tkn.IsAddOp(), tkn.String())
		} else {
			assert.False(t, tkn.IsAddOp(), tkn.String())
		}
	}
}

func TestTokenIsMulOp(t *testing.T) {
	for i := 0; i < len(tokens); i++ {
		tkn := Token(i)
		if tkn == MULTIPLY || tkn == DIVIDE || tkn == MODULO {
			assert.True(t, tkn.IsMulOp(), tkn.String())
		} else {
			assert.False(t, tkn.IsMulOp(), tkn.String())
		}
	}
}

func TestTokenIsCmpOp(t *testing.T) {
	for i := 0; i < len(tokens); i++ {
		tkn := Token(i)
		if tkn == EQUAL || tkn == NOT_EQUAL || tkn == LESS ||
			tkn == LESS_EQ || tkn == GREATER || tkn == GREATER_EQ {
			assert.True(t, tkn.IsCmpOp(), tkn.String())
		} else {
			assert.False(t, tkn.IsCmpOp(), tkn.String())
		}
	}
}

func TestTokenIsLogOp(t *testing.T) {
	for i := 0; i < len(tokens); i++ {
		tkn := Token(i)
		if tkn == AND || tkn == OR || tkn == NOT {
			assert.True(t, tkn.IsLogOp(), tkn.String())
		} else {
			assert.False(t, tkn.IsLogOp(), tkn.String())
		}
	}
}

func TestTokenIsBinOp(t *testing.T) {
	for i := 0; i < len(tokens); i++ {
		tkn := Token(i)
		if tkn == ADD || tkn == SUBTRACT ||
			tkn == MULTIPLY || tkn == DIVIDE || tkn == MODULO ||
			tkn == EQUAL || tkn == NOT_EQUAL || tkn == LESS ||
			tkn == LESS_EQ || tkn == GREATER || tkn == GREATER_EQ ||
			tkn == AND || tkn == OR {
			assert.True(t, tkn.IsBinOp(), tkn.String())
		} else {
			assert.False(t, tkn.IsBinOp(), tkn.String())
		}
	}
}

func TestTokenIsUnaryOp(t *testing.T) {
	for i := 0; i < len(tokens); i++ {
		tkn := Token(i)
		if tkn == ADD || tkn == SUBTRACT || tkn == NOT {
			assert.True(t, tkn.IsUnaryOp(), tkn.String())
		} else {
			assert.False(t, tkn.IsUnaryOp(), tkn.String())
		}
	}
}

func TestTokenIsStmtKeyword(t *testing.T) {
	for i := 0; i < len(tokens); i++ {
		tkn := Token(i)
		if tkn == IF || tkn == WHILE || tkn == FUNC || tkn == RETURN {
			assert.True(t, tkn.IsStmtKeyword(), tkn.String())
		} else {
			assert.False(t, tkn.IsStmtKeyword(), tkn.String())
		}
	}
}

func TestTokenIsLiteral(t *testing.T) {
	for i := 0; i < len(tokens); i++ {
		tkn := Token(i)
		if tkn == IDENTIFIER || tkn == BOOL || tkn == NUMBER ||
			tkn == STRING {
			assert.True(t, tkn.IsLiteral(), tkn.String())
		} else {
			assert.False(t, tkn.IsLiteral(), tkn.String())
		}
	}
}

func TestTokenPrecedence(t *testing.T) {
	tokens := []struct{ t1, t2 Token }{
		{MULTIPLY, ADD},
		{ADD, LESS},
		{LESS, AND},
	}
	for _, tkns := range tokens {
		p := Precedence(tkns.t1) > Precedence(tkns.t2)
		assert.True(t, p, tkns.t1.String()+" and "+tkns.t2.String())
	}
}

func TestTokenPrecedenceBadToken(t *testing.T) {
	i := Precedence(IF)
	assert.Equal(t, 0, i)
}
