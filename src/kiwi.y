%{
#include <stdio.h>
#include "unicode/uchar.h"
#include "unicode/ustdio.h"
#include "unicode/ustring.h"
#include "y.tab.h"

extern int yylex(void);
extern void yyerror(const char *s);
%}

%token T_EOF T_IF T_AND T_OR T_NOT T_EQUAL T_ADD T_SUBTRACT T_CONCAT T_MULTIPLY
%token T_DIVIDE T_NOT_EQUAL T_LESS T_LESS_EQUAL T_GREATER T_GREATER_EQUAL
%token T_LBRACE T_RBRACE T_ELSE T_COMMA T_DOT T_ASSIGN T_RETURN T_VAR T_TRUE
%token T_FALSE T_LPAREN T_RPAREN T_WHILE T_FUNC T_IS T_MODULO T_ADD_ASSIGN
%token T_SUBTRACT_ASSIGN T_MULTIPLY_ASSIGN T_DIVIDE_ASSIGN T_MODULO_ASSIGN 

%token T_COMMENT

%union
{
	int number;
	UChar *string;
}

%token <number> T_NUMBER
%token <string> T_IDENTIFIER T_STRING

%start program

%%
program
	:
	| statementlist
	;

statementlist
	: statement
	| statementlist statement
	;

statement
	: complexstmt
	| simplestmt
	;

complexstmt
	: compoundstmt
	| functiondef
	;

compoundstmt
	: ifstmt
	| whilestmt
	;

ifstmt
	: T_IF expression compoundbody
	| T_IF expression compoundbody elsestmt
	;

expression
	: notexpr
	| expression T_AND notexpr
	| expression T_OR notexpr
	;

notexpr
	: compareexpr
	| T_NOT compareexpr
	;

compareexpr
	: minorexpr 
	| compareexpr compop minorexpr
	;

compop
	: T_EQUAL
	| T_NOT_EQUAL
	| T_LESS
	| T_LESS_EQUAL
	| T_GREATER
	| T_GREATER_EQUAL
        | T_IS
	;

minorexpr
	: term
	| minorexpr T_ADD term
	| minorexpr T_SUBTRACT term
	| minorexpr T_CONCAT term
	;

term
	: factor
	| term T_MULTIPLY factor
	| term T_DIVIDE factor
        | term T_MODULO factor
	;

factor
	: atom
	| T_ADD factor
	| T_SUBTRACT factor
	;

atom
	: T_IDENTIFIER
	| funccall
	| T_NUMBER 
	| T_TRUE
	| T_FALSE
	| T_STRING
	| T_LPAREN expression T_RPAREN 
	;

funccall
	: T_IDENTIFIER T_LPAREN T_RPAREN
	| T_IDENTIFIER T_LPAREN expressionlist T_RPAREN
	;

expressionlist
	: expression
	| expressionlist T_COMMA expression
	;

elsestmt
	: T_ELSE compoundbody
	| T_ELSE ifstmt
	;

compoundbody
	: T_LBRACE compoundbodylist T_RBRACE
	;

compoundbodylist
	: simplestmt
	| compoundstmt
	| compoundbodylist simplestmt
	| compoundbodylist compoundstmt
	;

whilestmt
	: T_WHILE expression compoundbody
	;

functiondef
	: T_FUNC T_IDENTIFIER compoundbody
	| T_FUNC T_IDENTIFIER funcparamlist compoundbody
	;

funcparamlist
	: T_IDENTIFIER
	| funcparamlist T_COMMA T_IDENTIFIER
	;

simplestmt
	: assignstmt T_DOT
	| returnstmt T_DOT
	| varstmt T_DOT
	| expression T_DOT
	;
	
assignstmt
	: T_IDENTIFIER assignop expression
	;

assignop
	: T_ASSIGN
        | T_ADD_ASSIGN
        | T_SUBTRACT_ASSIGN
        | T_MULTIPLY_ASSIGN
        | T_DIVIDE_ASSIGN
        | T_MODULO_ASSIGN
	;


returnstmt
	: T_RETURN expression
	;

varstmt
	: T_VAR varstmtlist
	;

varstmtlist
	: T_IDENTIFIER
	| assignstmt
	| varstmtlist T_COMMA T_IDENTIFIER
	| varstmtlist T_COMMA assignstmt
	;
%%

