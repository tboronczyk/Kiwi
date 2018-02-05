package main

import "strconv"

// Token represents the token types used to represent Kiwi lexemes.
type Token uint

const (
	T_UNKNOWN Token = iota
	T_EOF

	addop_start
	// addition-level operators
	T_ADD
	T_SUBTRACT
	addop_end

	mulop_start
	// multiplication-level operators
	T_MULTIPLY
	T_DIVIDE
	T_MODULO
	mulop_end

	cmpop_start
	// comparision operators
	T_EQUAL
	T_NOT_EQUAL
	T_GREATER
	T_GREATER_EQ
	T_LESS
	T_LESS_EQ
	cmpop_end

	logop_start
	// logic operators
	T_AND
	T_OR
	T_NOT
	logop_end

	stmtkwd_start
	// statement keywords
	T_IF
	T_FUNC
	T_RETURN
	T_WHILE
	stmtkwd_end

	lit_start
	// literal values
	T_BOOL
	T_IDENTIFIER
	T_NUMBER
	T_STRING
	lit_end

	T_ASSIGN
	T_LBRACE
	T_RBRACE
	T_COLON
	T_COMMA
	T_COMMENT
	T_ELSE
	T_LPAREN
	T_RPAREN
)

var tokens = []string{
	T_UNKNOWN:    "UNKNOWN",
	T_EOF:        "EOF",
	T_ADD:        "+",
	T_SUBTRACT:   "-",
	T_MULTIPLY:   "*",
	T_DIVIDE:     "/",
	T_MODULO:     "%",
	T_EQUAL:      "=",
	T_NOT_EQUAL:  "~=",
	T_GREATER:    ">",
	T_GREATER_EQ: ">=",
	T_LESS:       "<",
	T_LESS_EQ:    "<=",
	T_AND:        "&&",
	T_OR:         "||",
	T_NOT:        "~",
	T_IF:         "if",
	T_FUNC:       "func",
	T_RETURN:     "return",
	T_WHILE:      "while",
	T_BOOL:       "BOOL",
	T_IDENTIFIER: "IDENTIFIER",
	T_NUMBER:     "NUMBER",
	T_STRING:     "STRING",
	T_ASSIGN:     ":=",
	T_LBRACE:     "{",
	T_RBRACE:     "}",
	T_COLON:      ":",
	T_COMMA:      ",",
	T_COMMENT:    "COMMENT",
	T_ELSE:       "else",
	T_LPAREN:     "(",
	T_RPAREN:     ")",
}

// String returns the string representation of a token.
func (t Token) String() string {
	str := ""
	if t >= 0 && t < Token(len(tokens)) {
		str = tokens[t]
	}
	if str == "" {
		str = "Token(" + strconv.Itoa(int(t)) + ")"
	}
	return str
}

// Precedence returns the relative precedence of an operator. A higher value
// is a higher precedence.
func Precedence(t Token) int {
	if t.IsLogOp() {
		return 1
	}
	if t.IsCmpOp() {
		return 2
	}
	if t.IsAddOp() {
		return 3
	}
	if t.IsMulOp() {
		return 4
	}
	return 0
}

// IsAddOp returns bool to indicate whether the token represents an
// addition-level operator.
func (t Token) IsAddOp() bool {
	return t > addop_start && t < addop_end
}

// IsMulOp returns bool to indicate whether the token represents a
// multiplication-level operator.
func (t Token) IsMulOp() bool {
	return t > mulop_start && t < mulop_end
}

// IsCmpOp returns bool to indicate whether the token represents a comparision
// operator.
func (t Token) IsCmpOp() bool {
	return t > cmpop_start && t < cmpop_end
}

// IsLogOp returns bool to indicate whether the token represents a logic
// operator.
func (t Token) IsLogOp() bool {
	return t > logop_start && t < logop_end
}

// IsBinOp returns bool to indicate whether the token represents a left-binding
// binary operator.
func (t Token) IsBinOp() bool {
	return (t.IsAddOp() || t.IsMulOp() || t.IsCmpOp() || t.IsLogOp()) &&
		t != T_NOT
}

// IsUnaryOp returns bool to indicate whether the token represents a
// right-binding operator.
func (t Token) IsUnaryOp() bool {
	return t.IsAddOp() || t == T_NOT
}

// IsStmtKeyword returns bool to indicate whether the token represents a keyword
// that may begin a statement.
func (t Token) IsStmtKeyword() bool {
	return t > stmtkwd_start && t < stmtkwd_end
}

// IsLiteral returns bool to indicate whether the token represents a literal
// value.
func (t Token) IsLiteral() bool {
	return t > lit_start && t < lit_end
}
