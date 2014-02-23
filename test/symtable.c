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

#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include "symtable.h"

int main(void)
{
    SymTable *t;
    int a, b;

    t = symtable_init();
    a = 40;
    b = 41;

    symtable_insert(t, "a", SYMTABLE_ENTRY_NUMBER, &a);
    symtable_enter_scope(t);
    symtable_insert(t, "b", SYMTABLE_ENTRY_NUMBER, &b);
    symtable_enter_scope(t);

    if (symtable_lookup(t, "a")) printf("a=%d\n", *(int *)symtable_lookup(t, "a")); else printf("a not found\n");
    if (symtable_lookup(t, "b")) printf("b=%d\n", *(int *)symtable_lookup(t, "b")); else printf("b not found\n");
    symtable_leave_scope(t);
    if (symtable_lookup(t, "a")) printf("a=%d\n", *(int *)symtable_lookup(t, "a")); else printf("a not found\n");
    if (symtable_lookup(t, "b")) printf("b=%d\n", *(int *)symtable_lookup(t, "b")); else printf("b not found\n");
    symtable_leave_scope(t);
    if (symtable_lookup(t, "a")) printf("a=%d\n", *(int *)symtable_lookup(t, "a")); else printf("a not found\n");
    if (symtable_lookup(t, "b")) printf("b=%d\n", *(int *)symtable_lookup(t, "b")); else printf("b not found\n");

    return EXIT_SUCCESS;
}
