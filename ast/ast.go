/*
 * Copyright (c) 2012, 2015 Timothy Boronczyk
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are met:
 *
 *  1. Redistributions of source code must retain the above copyright notice,
 *     this list of conditions and the following disclaimer.
 *
 *  2. Redistributions in binary form must reproduce the above copyright
 *     notice, this list of conditions and the following disclaimer in the
 *     documentation and/or other materials provided with the distribution.
 *
 *  3. The names of the authors may not be used to endorse or promote products
 *     derived from this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED "AS IS" AND WITHOUT ANY EXPRESS OR IMPLIED
 * WARRANTIES, INCLUDING, WITHOUT LIMITATION, THE IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE.
 */

package ast

import (
	"fmt"
	"github.com/tboronczyk/kiwi/token"
)

type Node interface {
	print(string)
}

type (
	Literal struct {
		Type  token.Token
		Value string
	}

	Operator struct {
		Op    token.Token
		Left  Node
		Right Node
	}

	List struct {
		Node
		Next Node
	}

	FuncCall struct {
		Name string
		Body Node
	}

	If struct {
		Condition Node
		Body Node
	}

	While struct {
		Condition Node
		Body Node
	}
)

func NewLiteral(t token.Token, v string) Literal {
	return Literal{Type: t, Value: v}
}

func NewOperator(t token.Token) Operator {
	return Operator{Op: t}
}

func NewFuncCall(s string) FuncCall {
	return FuncCall{Name: s}
}

func NewList() List {
	return List{}
}

func NewIf() If {
	return If{}
}

func NewWhile() While {
	return While{}
}

func Print(n Node) {
	n.print("")
}

func (n Literal) print(s string) {
	fmt.Printf("%s%s (%s)\n", s, n.Value, n.Type.String())
}

func (n Operator) print(s string) {
	if n.Left != nil {
		n.Left.print(s + "OP.L ")
	}
	fmt.Printf("%sOP %s\n", s, n.Op.String())
	if n.Right != nil {
		n.Right.print(s + "OP.R ")
	}
}

func (n FuncCall) print(s string) {
	fmt.Printf("%sF.N %s\n", s, n.Name)
	if n.Body != nil {
		n.Body.print(s + "F.B ")
	}
}
	
func (n List) print(s string) {
	if n.Node != nil {
		n.Node.print(s + "L.N ")
	}
	if n.Next != nil {
		n.Next.print(s + "L.n ")
	}
}

func (n If) print(s string) {
	if n.Condition != nil {
		n.Condition.print(s + "IF.C ")
	}
	if n.Body != nil {
		n.Body.print(s + "IF.B ")
	}
}

func (n While) print(s string) {
	if n.Condition != nil {
		n.Condition.print(s + "WL.C ")
	}
	if n.Body != nil {
		n.Body.print(s + "WL.B ")
	}
}
