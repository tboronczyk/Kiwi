#ifndef SCANNER_H
#define SCANNER_H

/*
 * Copyright (c) 2012, Timothy Boronczyk
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

#include <stdio.h>
#include "unicode/ustdio.h"

typedef enum
{
    SCANERR_OK,
    SCANERR_ALLOC_SCANNER,
    SCANERR_ALLOC_FILENAME,
    SCANERR_FILEOPEN,
    SCANERR_ALLOC_BUFFER,
    SCANERR_REALLOC_BUFFER,
    SCANERR_UNEXPECTED_LEX
}
Scanner_ErrCode;

typedef struct s_Scanner Scanner;

Scanner_ErrCode scanner_init(Scanner **);
void scanner_free(Scanner *);

Scanner_ErrCode scanner_token(Scanner *);
int scanner_error(Scanner *, const char *);

struct s_Scanner
{
    int linenum,      /* current line number of file (used for error reporting) */
        linepos,      /* current position in file line */
        name;         /* name of token being scanned */
    unsigned int ti;  /* current position of pointer in *tbuf */
    size_t tlen;      /* size of *tbuf */
    UChar c,          /* current character read */
        *tbuf;        /* buffer in which token values are accumulated */
    char *fname;      /* name of file being scanned (used for error reporting) */
    UFILE *fp;        /* open file descriptor of file at *fname */
};

#endif
