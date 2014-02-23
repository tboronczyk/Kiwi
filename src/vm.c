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
#include <stdlib.h>
#include <stdarg.h>
#include "vm.h"

#define UNUSED(x) (void)(x)

#define VM_PROGBUF_SIZE_INIT 5
#define VM_PROGBUF_SIZE_INCR 5

VM_Mach *vm_mach_init(void)
{
    VM_Mach *vm;
    int i;

    /* allocate memory for machine */
    if ((vm = (VM_Mach *)calloc(1, sizeof(VM_Mach))) == NULL) {
        perror("Allocate machine failured");
        exit(EXIT_FAILURE);
    }

    /* initialize machine */
    vm->sp = -1;
    vm->ip = 0;
    for (i = 0; i < VM_MACH_NUM_REGS; i++) {
        vm->regs[i] = (int *)calloc(1, sizeof(int));
    }

    return vm;
}

void vm_mach_free(VM_Mach *vm)
{
    int i;
    for (i = 0; i < VM_MACH_NUM_REGS; i++) {
        free(vm->regs[i]);
    }
    free(vm);
}

static VM_Instr *vm_instr_init(VM_Opcode op, ...)
{
    VM_Instr *instr;
    va_list ap;

    /* allocate memory for instruction */
    if ((instr = (VM_Instr *)calloc(1, sizeof(VM_Instr))) == NULL) {
        perror("Allocate instruction failed");
        exit(EXIT_FAILURE);
    }

    /* initalize instruction */
    instr->op = op;
    va_start(ap, op);
    switch (op) {
    case OP_NOOP: break;
    case OP_MOVE: instr->dest = va_arg(ap, int); instr->src = va_arg(ap, int); break;
    case OP_XCHG: instr->dest = va_arg(ap, int); instr->src = va_arg(ap, int); break;
/*
    case OP_VAR:
    case OP_LOAD:
    case OP_STOR:
*/
    case OP_PUSH: instr->dest = va_arg(ap, int); break;
    case OP_POP:  instr->dest = va_arg(ap, int); break;
    case OP_ADD:  instr->dest = va_arg(ap, int); instr->src = va_arg(ap, int); break;
    case OP_SUB:  instr->dest = va_arg(ap, int); instr->src = va_arg(ap, int); break;
    case OP_MUL:  instr->dest = va_arg(ap, int); instr->src = va_arg(ap, int); break;
    case OP_DIV:  instr->dest = va_arg(ap, int); instr->src = va_arg(ap, int); break;
    case OP_NEG:  instr->dest = va_arg(ap, int); break;
/*
    case OP_CCAT:
*/
    case OP_AND:  instr->dest = va_arg(ap, int); instr->src = va_arg(ap, int); break;
    case OP_OR:   instr->dest = va_arg(ap, int); instr->src = va_arg(ap, int); break;
    case OP_NOT:  instr->dest = va_arg(ap, int); break;
/*
    case OP_CMP:
    case OP_JMP:
*/
    }
    va_end(ap);

    return instr;
}

static void vm_instr_free(VM_Instr *instr)
{
    free(instr);
}

void vm_mach_exec(VM_Mach *vm, VM_ProgBuf *b)
{
    int dest, src, tmp1, tmp2;
    VM_Instr *instr;

    for (vm->ip = 0 ; vm->ip < b->tail; vm->ip++) {
        instr = b->instr[vm->ip];

        switch (instr->op) {
        case OP_NOOP:
            break;

        case OP_MOVE:
            dest = instr->dest;
            src = instr->src;
            *vm->regs[dest] = src;
            break;

        case OP_XCHG:
            dest = instr->dest;
            src = instr->src;
            tmp1 = *vm->regs[dest];
            tmp2 = *vm->regs[src];
            *vm->regs[dest] = tmp2;
            *vm->regs[src] = tmp1;
            break;
/*
        case OP_VAR:
            break;

        case OP_LOAD:
            break;

        case OP_STOR:
            break;
*/
        case OP_PUSH:
            dest = instr->dest;

            if (vm->sp != VM_MACH_SIZE_STACK - 1) {
                vm->sp++;
                vm->stack[vm->sp] = *vm->regs[dest];
            }
            else {
                perror("Stack push failed");
                exit(EXIT_FAILURE);
            }
            break;

        case OP_POP:
            dest = instr->dest;

            if (vm->sp != -1) {
                *vm->regs[dest] = vm->stack[vm->sp];
                vm->sp--;
            }
            else {
                perror("Stack pop failed");
                exit(EXIT_FAILURE);
            }
            break;

        case OP_ADD:
            dest = instr->dest;
            src = instr->src;
            *vm->regs[dest] += *vm->regs[src];
            break;

        case OP_SUB:
            dest = instr->dest;
            src = instr->src;
            *vm->regs[dest] -= *vm->regs[src];
            break;

        case OP_MUL:
            dest = instr->dest;
            src = instr->src;
            *vm->regs[dest] *= *vm->regs[src];
            break;

        case OP_DIV:
            dest = instr->dest;
            src = instr->src;
            *vm->regs[dest] /= *vm->regs[src];
            break;

        case OP_NEG:
            dest = instr->dest;
            *vm->regs[dest] = -*vm->regs[dest];
            break;
/*
        case OP_CCAT:
            break;
*/
        case OP_AND:
            dest = instr->dest;
            src = instr->src;
            *vm->regs[dest] = (int)(*vm->regs[dest] != 0 && *vm->regs[src] != 0);
            break;

        case OP_OR:
            dest = instr->dest;
            src = instr->src;
            *vm->regs[dest] = (int)(*vm->regs[dest] !=0 || *vm->regs[src] != 0);
            break;

        case OP_NOT:
            dest = instr->dest;
            *vm->regs[dest] = (int)(!(*vm->regs[dest] != 0));
            break;
/*
        case OP_CMP:
            break;

        case OP_JMP:
            break;
*/
        }
#ifdef DEBUG
        printf(":%d %d %d\n", *vm->regs[0], *vm->regs[1], *vm->regs[2]);
#endif
    }
}

static VM_ProgBuf *vm_progbuf_init(void)
{
    VM_ProgBuf *b;

    /* allocate memory for program buffer */
    if ((b = (VM_ProgBuf *)calloc(1, sizeof(VM_ProgBuf))) == NULL) {
        perror("Allocate program buffer failed");
        exit(EXIT_FAILURE);
    }

    /* initialize buffer */
    b->tail = 0;
    b->len = VM_PROGBUF_SIZE_INIT;
    if ((b->instr = (VM_Instr **)calloc(b->len, sizeof(VM_Instr *))) == NULL) {
        perror("Allocate program buffer instruction storage failed");
        exit(EXIT_FAILURE);
    }

    return b;
}

static void vm_progbuf_free(VM_ProgBuf *b)
{
    int i;
    for (i = 0; i < b->tail; i++) {
        vm_instr_free(b->instr[i]);
    }
    free(b->instr);
    free(b);
}

static void vm_progbuf_grow(VM_ProgBuf *b)
{
    /* increase storage capacity of buffer */
    b->len += VM_PROGBUF_SIZE_INCR;
    if ((b->instr = (VM_Instr **)realloc(b->instr, sizeof(VM_Instr *) * b->len)) == NULL) {
        perror("Reallocate program buffer instruction storage failed");
        exit(EXIT_FAILURE);
    }
}

static void vm_progbuf_push(VM_ProgBuf *b, VM_Instr *i)
{
    b->instr[b->tail] = i;
    b->tail++;
    /* increase buffer size if necessary */
    if (b->tail == b->len) {
        vm_progbuf_grow(b);
    }
}

int main()
{
    VM_ProgBuf *b = vm_progbuf_init();
    VM_Mach *vm = vm_mach_init();

    /* load program */
    vm_progbuf_push(b, vm_instr_init(OP_NOOP));
    vm_progbuf_push(b, vm_instr_init(OP_MOVE, 0, 10));
    vm_progbuf_push(b, vm_instr_init(OP_MOVE, 1, 1));
    vm_progbuf_push(b, vm_instr_init(OP_SUB, 0, 1));

    /* execute program */
    vm_mach_exec(vm, b);

    vm_progbuf_free(b);
    vm_mach_free(vm);

    return EXIT_SUCCESS;
}

