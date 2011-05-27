/*
 * Copyright (c) 2011, Timothy Boronczyk
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

#ifndef TOKEN_H
#define TOKEN_H

// operators
#define T_ASSIGN                 1
#define T_ADD                    2
#define T_ADD_ASSIGN             3
#define T_SUBTRACT               4
#define T_SUBTRACT_ASSIGN        5
#define T_MULTIPLY               6
#define T_MULTIPLY_ASSIGN        7
#define T_DIVIDE                 8
#define T_DIVIDE_ASSIGN          9
#define T_EQUAL                  10
#define T_NOT_EQUAL              11
#define T_LESS                   12
#define T_LESS_EQUAL             13
#define T_GREATER                14
#define T_GREATER_EQUAL          15
#define T_LOG_AND                16
#define T_LOG_OR                 17
#define T_LOG_XOR                18
#define T_LOG_NOT                19

// literals
#define T_WILDCARD               20
#define T_NUMBER                 21
#define T_NUMBER_INT_2           22
#define T_NUMBER_INT_8           23
#define T_NUMBER_INT_16          24

// punctuators
#define T_BRACE_LEFT             25
#define T_BRACE_RIGHT            26
#define T_PAREN_LEFT             27
#define T_PAREN_RIGHT            28

// comments
#define T_COMMENT                29
#define T_COMMENT_MULTI          30

// identifier
#define T_IDENTIFIER             31 

// keywords
#define T_IF                     32
#define T_ELSE                   33
#define T_IS                     34

typedef struct {
    int name;
    char *lexeme;
} Token;

void token_free(Token *t);

#endif

