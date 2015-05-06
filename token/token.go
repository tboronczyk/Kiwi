/*
 * Copyright (c) 2012, 2015 Timothy Boronczyk
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are met:
 *
 * 1. Redistributions of source code must retain the above copyright notice,
 *    this list of conditions and the following disclaimer.
 *
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 *
 * 3. The names of the authors may not be used to endorse or promote products
 *    derived from this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED "AS IS" AND WITHOUT ANY EXPRESS OR IMPLIED
 * WARRANTIES, INCLUDING, WITHOUT LIMITATION, THE IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE.
 */

package token

import (
	"strconv"
)

type Token uint8

const (
	UNKNOWN Token = iota
	MALFORMED
	EOF
	COMMENT

	addop_start
	ADD
	SUBTRACT
	addop_end

	mulop_start
	MULTIPLY
	DIVIDE
	MODULO
	mulop_end

	cmpop_start
	EQUAL
	NOT_EQUAL
	GREATER
	GREATER_EQ
	LESS
	LESS_EQ
	cmpop_end

	logop_start
	AND
	OR
	NOT
	logop_end

	LPAREN
	RPAREN

	lit_start
	TRUE
	FALSE
	NUMBER
	STRING
	IDENTIFIER
	lit_end
)

var tokens = [...]string{
	UNKNOWN:    "UNKNOWN",
	MALFORMED:  "MALFORMED",
	EOF:        "EOF",
	COMMENT:    "COMMENT",
	ADD:        "+",
	SUBTRACT:   "-",
	MULTIPLY:   "*",
	DIVIDE:     "/",
	MODULO:     "%",
	EQUAL:      "=",
	NOT_EQUAL:  "~=",
	GREATER:    ">",
	GREATER_EQ: ">=",
	LESS:       "<",
	LESS_EQ:    "<=",
	AND:        "&&",
	OR:         "||",
	NOT:        "~",
	LPAREN:     "(",
	RPAREN:     ")",
	TRUE:       "true",
	FALSE:      "false",
	NUMBER:     "NUMBER",
	STRING:     "STRING",
	IDENTIFIER: "IDENTIFIER",
}

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

func (t Token) IsAddOp() bool {
	return t > addop_start && t < addop_end
}

func (t Token) IsMulOp() bool {
	return t > mulop_start && t < mulop_end
}

func (t Token) IsCmpOp() bool {
	return t > cmpop_start && t < cmpop_end
}

func (t Token) IsLogOp() bool {
	return t > logop_start && t < logop_end
}

func (t Token) IsLiteral() bool {
	return t > lit_start && t < lit_end
}
