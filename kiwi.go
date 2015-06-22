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

package main

import (
	"bufio"
	"fmt"
	"github.com/tboronczyk/kiwi/ast"
	"github.com/tboronczyk/kiwi/parser"
	"github.com/tboronczyk/kiwi/scanner"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	s := scanner.NewScanner(r)
	p := parser.NewParser()
	p.InitScanner(s)

	for {
		n, err := p.Parse()
		if n == nil {
			if err != nil {
				fmt.Println(err)
			}
			return
		}
		ast.Print(n)
	}
}
